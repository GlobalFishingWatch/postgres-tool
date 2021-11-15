package action

import (
	"context"
	"fmt"
	"github.com/GlobalFishingWatch/postgres-tool/internal/common"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
)


func CreateIndex(params types.CreateIndexParam, postgresConfig types.PostgresConfigParams) {
	ctx := context.Background()
	psClient := common.CreatePostgresClient(ctx, postgresConfig)
	defer psClient.Close(ctx)
	log.Println("Creating a new index")
	createIndex(ctx, psClient, params.TableName, params.IndexName, params.Column)
}

func createIndex(ctx context.Context, psClient *pgx.Conn, tableName string, indexName string, column string) {
	createViewCommand := fmt.Sprintf(`
		CREATE INDEX %s 
    		ON %s(%s)
		`, indexName, tableName, column)
	log.Printf("→ PG →→ Creating index with command %s", createViewCommand)
	_, err := psClient.Exec(ctx, createViewCommand)
	if err != nil {
		log.Fatalf("→ PG →→ Error creating index: %v", err)
	}
	log.Printf("→ PG →→ Successfully created view with name %v", indexName)
}
