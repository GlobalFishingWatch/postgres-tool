package action

import (
	"context"
	"fmt"
	"github.com/GlobalFishingWatch/postgres-tool/internal/common"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
)

func DeleteTable(params types.DeleteTableParams, postgresConfig types.PostgresConfigParams) {
	ctx := context.Background()
	psClient := common.CreatePostgresClient(ctx, postgresConfig)
	defer psClient.Close(ctx)
	log.Println("Deleting a table")
	deleteTable(ctx, psClient, params.TableName)
}

func deleteTable(ctx context.Context, psClient *pgx.Conn, tableName string) {
	deleteTableCommand := fmt.Sprintf(`DROP TABLE IF EXISTS %s;`, tableName)
	log.Printf("→ PG →→ Deleting table with name %s and command %s", tableName, deleteTableCommand)
	_, err := psClient.Exec(ctx, deleteTableCommand)
	if err != nil {
		log.Fatalf("→ PG →→ Error deleting table: %v", err)
	}

	log.Printf("→ PG →→ Successfully deleting table with name %v", tableName)
}