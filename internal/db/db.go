package db

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/SadenYernar/gRPC-HTTP-Cobra-PostgreSQL-Docker-goose.git/internal/logger"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	Client *pgxpool.Pool
}

func Get(ctx context.Context, connStr string) (*DB, error) {
	dbConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		logger.Errorf("unable to parse db config: %s")
	}

	var dbPool *pgxpool.Pool

	for i := 1; i < 8; i++ {
		logger.Infof("trying to connect to the db server (attempt %d)...\n", i)

		dbPool, err = pgxpool.ConnectConfig(ctx, dbConfig)
		if err == nil {
			break
		}
		logger.Errorf("got error: %v\n", err)

		// Sleep a bit before trying again
		time.Sleep(time.Duration(i*i) * time.Second)
	}

	//Stop execution if the database was not initialized
	if dbPool == nil {
		logger.Errorf("could not connect to the database")
	}

	// Get a connection from the pool and check if the database connection is active and working
	db, err := dbPool.Acquire(ctx)
	if err != nil {
		logger.Errorf("failed to get connection on startup: %v\n", err)
	}

	if db.Conn().Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// Add the connection back to the pool
	db.Release()

	return &DB{
		dbPool,
	}, nil
}

func (d *DB) Close() {
	d.Client.Close()
}

func IsUniqueViolationErr(err error) bool {
	pgErr := &pgconn.PgError{Code: "23505"}
	return errors.As(err, &pgErr)
}
