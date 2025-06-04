package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/graph"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/pkg"
	pb "github.com/yaninyzwitty/pgxpool-twitter-roach/shared/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var cfg pkg.Config
	if err := cfg.LoadConfig("config.yaml"); err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	// Create gRPC connection
	grpcClientAddr := fmt.Sprintf(":%d", cfg.GrpcServer.Port)
	grpcConn, err := grpc.NewClient(grpcClientAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("failed to connect to grpc server", "error", err)
		os.Exit(1)
	}
	defer grpcConn.Close()

	// gRPC user service client
	socialServiceClient := pb.NewUserServiceClient(grpcConn)

	// GraphQL server setup
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			SocialServiceClient: socialServiceClient,
		},
	}))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.SetQueryCache(lru.New[*ast.QueryDocument](100))
	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// Setup chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	r.Handle("/query", srv)

	// Start HTTP server with graceful shutdown
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	srvHTTP := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	// Graceful shutdown setup
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		slog.Info("API Gateway running", "port", cfg.Server.Port)
		if err := srvHTTP.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	<-quit
	slog.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srvHTTP.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exited gracefully")
}
