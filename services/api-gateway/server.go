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
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/graph"
	"github.com/yaninyzwitty/pgxpool-twitter-roach/pkg"
)

func main() {
	var cfg pkg.Config
	if err := cfg.LoadConfig("config.yaml"); err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	app := fiber.New(fiber.Config{
		Immutable: true,
	})
	app.Use(logger.New())

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	app.Post("/query", FiberHandler(srv))
	app.Get("/", FiberHandler(playground.Handler("GraphQL Playground", "/query")))

	slog.Info("API Gateway running", "port", cfg.Server.Port)
	addr := fmt.Sprintf(":%d", cfg.Server.Port)

	// Create a channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	// Register the channel to receive specific signals
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := app.Listen(addr); err != nil {
			slog.Error("failed to start server", "error", err)
			os.Exit(1)
		}
	}()
	// Block until we receive a signal
	<-quit
	slog.Info("Shutting down server...")

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	slog.Info("Server exited gracefully")

}

func FiberHandler(h http.Handler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(h)(c.Context())
		return nil
	}
}
