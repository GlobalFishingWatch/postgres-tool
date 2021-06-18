package action

import (
	"context"
	"github.com/GlobalFishingWatch/postgres-tool/internal/common"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)


func ExecuteRawSql(params types.ExecuteRawSqlParams, postgresConfig types.PostgresConfigParams) {
	ctx := context.Background()
	psClient := common.CreatePostgresClient(ctx, postgresConfig)
	defer psClient.Close(ctx)
	log.Println("Executing the command")
	retries := 0
	executeCommand(ctx, psClient, params.Sql, retries)
}

func executeCommand(ctx context.Context, psClient *pgx.Conn, sql string, retries int) {
	log.Printf("→ PG →→ Executing the next SQL:  %s. Retries: %v", sql, retries)
	_, err := psClient.Exec(ctx, sql)
	if err != nil {
		retries = retries + 1;
		log.Printf("→ PG →→ Error executing the SQL:  %s. Retries: %v", sql, retries)
		if retries < 3 {
			time.Sleep(60 * time.Second)
			executeCommand(ctx, psClient, sql, retries)
		} else {
			log.Fatalf("→ PG →→ Error executing the sql: %v. Process finished.", err)
		}
	}

	log.Print("→ PG →→ Successfully executing the command")
}