package connection

import (
	"context"
	"log"
	"task-manager/internal/config"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConnection(config *config.ServiceConfig) (*pgxpool.Pool, error) {

	dbPool, dbErr := pgxpool.New(context.Background(), config.DbConnectionUrl)
	if dbErr != nil {
		log.Fatal("database connection error", dbErr)
		return nil, dbErr
	}

	return dbPool, nil
}

func PingDB(db *pgxpool.Pool) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return db.Ping(ctx)
}
