package database

import (
	"context"
	"fmt"
	"time"

	logPkg "github.com/2SSK/EchoNext/internal/logger"
	pgxzero "github.com/jackc/pgx-zerolog"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"
	"github.com/rs/zerolog"
)

type Database struct {
	Pool *pgxpool.Pool
	log  *zerolog.Logger
}

const DatabasePingTimeout = 10

func New(dsn string, appLogger *zerolog.Logger) (*Database, error) {
	pgxPoolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx pool config: %w", err)
	}

	// Set up pgx tracing
	pgxLogger := logPkg.NewPgxLogger(appLogger.GetLevel())
	pgxPoolConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   pgxzero.NewLogger(pgxLogger),
		LogLevel: tracelog.LogLevel(logPkg.GetPgxTraceLogLevel(appLogger.GetLevel())),
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxPoolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	database := &Database{
		Pool: pool,
		log:  appLogger,
	}

	ctx, cancel := context.WithTimeout(context.Background(), DatabasePingTimeout*time.Second)
	defer cancel()
	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	appLogger.Info().Msg("connected to the database")

	return database, nil
}

func (db *Database) Close() error {
	db.log.Info().Msg("closing database connection pool")
	db.Pool.Close()
	return nil
}
