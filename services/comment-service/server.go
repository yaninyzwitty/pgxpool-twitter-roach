package main

import (
	"log/slog"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/pkg"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/pkg/database"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/pkg/utils"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
)

func main() {
	var cfg pkg.Config
	if err := cfg.LoadConfig("config.yaml"); err != nil {
		slog.Error("Failed to load config.yaml", slog.Any("error", err))
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		slog.Error("Failed to load .env", slog.Any("error", err))
		os.Exit(1)
	}

	cockroachDBPassword := utils.GetEnvOrDefault("COCROACH_DB_PASSWORD", "")
	if cockroachDBPassword == "" {
		slog.Warn("Missing CockroachDB password from environment")
	}

	roachConfig := &database.DBConfig{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.Username,
		Password: cockroachDBPassword,
		Database: cfg.Database.DbName,
		SSLMode:  cfg.Database.SslMode,
	}

	db, err := database.NewDB(cfg.Database.MaxRetries, 1*time.Second, roachConfig)
	if err != nil {
		slog.Error("Failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}
	defer db.Close()
	pool := db.Pool()
	slog.Info("Connected to database successfully")

	serverAddress := utils.BuildServerAddr(cfg.GrpcServer.Port)

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		slog.Error("Failed to listen", slog.Any("error", err))
		os.Exit(1)
	}

	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(),
	)
	repository := NewPostgresCommentServiceRepository(pool)
	NewCommentServiceGrpcHandler(grpcServer, repository)

	// Setup graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		sig := <-quit
		slog.Info("Initiating graceful shutdown", slog.Any("signal", sig))
		grpcServer.GracefulStop()
	}()

	slog.Info("Starting gRPC server", slog.String("address", serverAddress))
	if err := grpcServer.Serve(lis); err != nil {
		slog.Error("Failed to serve gRPC server", slog.Any("error", err))
		os.Exit(1)
	}

	wg.Wait()
	slog.Info("gRPC server stopped gracefully")
}
