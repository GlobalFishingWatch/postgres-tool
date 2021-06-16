package action

import (
	"context"
	"github.com/GlobalFishingWatch/postgres-tool/internal/common"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
)


func ExecuteRawSql(params types.ExecuteRawSqlParams, postgresConfig types.PostgresConfigParams) {
	ctx := context.Background()
	psClient := common.CreatePostgresClient(ctx, postgresConfig)
	defer psClient.Close(ctx)
	log.Println("Executing the command")
	executeCommand(ctx, psClient, params.Sql)
}

func executeCommand(ctx context.Context, psClient *pgx.Conn, sql string) {
	log.Printf("→ PG →→ Executing the next SQL:  %s", sql)
	_, err := psClient.Exec(ctx, sql)
	if err != nil {
		log.Fatalf("→ PG →→ Error executing the sql: %v", err)
	}

	log.Print("→ PG →→ Successfully executing the command")
}