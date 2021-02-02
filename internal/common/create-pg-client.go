package common

import (
	"context"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

func CreatePostgresClient(ctx context.Context, postgresConfig types.PostgresConfigParams) *pgx.Conn {
	uri := "postgresql://" + postgresConfig.Addr + "/" + postgresConfig.Database + "?user=" + postgresConfig.User + "&password=" + postgresConfig.Password
	conn, err := pgx.Connect(ctx, uri)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return conn
}