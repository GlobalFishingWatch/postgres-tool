package pkg

import (
	"github.com/GlobalFishingWatch/postgres-tool/internal/action"
	"github.com/GlobalFishingWatch/postgres-tool/types"
)

func ExecuteRawSql(params types.ExecuteRawSqlParams, postgresConfig types.PostgresConfigParams) {
	action.ExecuteRawSql(params, postgresConfig)
}

func DeleteView(params types.DeleteViewParams, postgresConfig types.PostgresConfigParams) {
	action.DeleteView(params, postgresConfig)
}

func DeleteTable(params types.DeleteTableParams, postgresConfig types.PostgresConfigParams) {
	action.DeleteTable(params, postgresConfig)
}

func CreateView(params types.CreateViewParams, postgresConfig types.PostgresConfigParams) {
	action.CreateView(params, postgresConfig)
}

func CreateIndex(params types.CreateIndexParam, postgresConfig types.PostgresConfigParams) {
	action.CreateIndex(params, postgresConfig)
}
