package action

import (
	"context"
	"fmt"
	"github.com/GlobalFishingWatch/postgres-tool/internal/common"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
)

var psClient *pgx.Conn

func CreateView(params types.CreateViewParams, postgresConfig types.PostgresConfigParams) {
	ctx := context.Background()

	psClient = common.CreatePostgresClient(ctx, postgresConfig)
	defer psClient.Close(ctx)
	log.Println("Creating a view")
	createView(ctx, params.ViewName, params.TableName)
}

func createView(ctx context.Context, viewName string, tableName string) {
	createViewCommand := fmt.Sprintf(`
		CREATE VIEW %s AS
    		SELECT *
    		FROM %s
		`, viewName, tableName)
	log.Printf("→ PG →→ Creating view with command %s", createViewCommand)
	_, err := psClient.Exec(ctx, createViewCommand)
	if err != nil {
		log.Fatalf("→ PG →→ Error creating view: %v", err)
	}
	log.Printf("→ PG →→ Successfully created view with name %v", viewName)
}
