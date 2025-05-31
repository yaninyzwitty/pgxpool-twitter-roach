package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool *pgxpool.Pool
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

// NewDB initializes and connects to the database with retries.
func NewDB(maxRetries int, retryInterval time.Duration, cfg *DBConfig) (*DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode,
	)

	var pool *pgxpool.Pool

	for attempt := 1; attempt <= maxRetries; attempt++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		pgxCfg, err := pgxpool.ParseConfig(connStr)
		if err != nil {
			cancel()
			return nil, fmt.Errorf("failed to parse pgx config: %w", err)
		}
		pgxCfg.MaxConns = 500
		pgxCfg.MinConns = 1

		pool, err = pgxpool.NewWithConfig(ctx, pgxCfg)
		if err == nil && pool.Ping(ctx) == nil {
			cancel()
			slog.Info("Successfully connected to database", "attempt", attempt)
			return &DB{pool: pool}, nil
		}

		slog.Warn("Failed to connect to database", "attempt", attempt, "error", err)
		cancel()
		time.Sleep(retryInterval)
	}

	return nil, fmt.Errorf("could not connect to database after %d attempts", maxRetries)
}

// Close gracefully closes the database connection pool.
func (db *DB) Close() {
	if db.pool != nil {
		slog.Info("Closing database connection pool")
		db.pool.Close()
	}
}

// Pool returns the underlying pgxpool.Pool instance.
func (db *DB) Pool() *pgxpool.Pool {
	return db.pool
}
