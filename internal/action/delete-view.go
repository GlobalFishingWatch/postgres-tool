package action

import (
	"context"
	"fmt"
	"github.com/GlobalFishingWatch/postgres-tool/internal/common"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/jackc/pgx/v4"
	"log"
)


func DeleteView(params types.DeleteViewParams, postgresConfig types.PostgresConfigParams) {
	ctx := context.Background()
	psClient := common.CreatePostgresClient(ctx, postgresConfig)
	defer psClient.Close(ctx)
	log.Println("Deleting a view")
	deleteView(ctx, psClient, params.ViewName)
}

func deleteView(ctx context.Context, psClient *pgx.Conn, viewName string) {
	deleteViewCommand := fmt.Sprintf(`DROP VIEW IF EXISTS %s;`, viewName)
	log.Printf("→ PG →→ Deleting view with name %s and command %s", viewName, deleteViewCommand)
	_, err := psClient.Exec(ctx, deleteViewCommand)
	if err != nil {
		log.Fatalf("→ PG →→ Error deleting view: %v", err)
	}

	log.Printf("→ PG →→ Successfully deleting view with name %v", viewName)
}