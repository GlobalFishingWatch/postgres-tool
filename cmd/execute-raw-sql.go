package cmd

import (
	"github.com/GlobalFishingWatch/postgres-tool/internal/action"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	executeRawSqlCmd.Flags().StringP("postgres-address", "", "", "The address of the database")
	executeRawSqlCmd.MarkFlagRequired("postgres-address")
	executeRawSqlCmd.Flags().StringP("postgres-user", "", "", "The destination credentials user")
	executeRawSqlCmd.MarkFlagRequired("postgres-user")
	executeRawSqlCmd.Flags().StringP("postgres-password", "", "", "The destination credentials password")
	executeRawSqlCmd.MarkFlagRequired("postgres-password")
	executeRawSqlCmd.Flags().StringP("postgres-database", "", "", "The destination database name")
	executeRawSqlCmd.MarkFlagRequired("postgres-database")

	executeRawSqlCmd.Flags().StringP("sql", "", "", "The sql statements to execute")
	executeRawSqlCmd.MarkFlagRequired("sql")

	viper.BindPFlag("execute-raw-sql-postgres-address", executeRawSqlCmd.Flags().Lookup("postgres-address"))
	viper.BindPFlag("execute-raw-sql-postgres-user", executeRawSqlCmd.Flags().Lookup("postgres-user"))
	viper.BindPFlag("execute-raw-sql-postgres-password", executeRawSqlCmd.Flags().Lookup("postgres-password"))
	viper.BindPFlag("execute-raw-sql-postgres-database", executeRawSqlCmd.Flags().Lookup("postgres-database"))
	viper.BindPFlag("execute-raw-sql", executeRawSqlCmd.Flags().Lookup("sql"))

	rootCmd.AddCommand(executeRawSqlCmd)
}

var executeRawSqlCmd = &cobra.Command{
	Use:   "execute-raw-sql",
	Short: "Execute raw sql",
	Long:  `Execute raw sql
Format:
	postgres execute-raw-sql --postgres-address= --postgres-user= --postgres-password= --postgres-database= --table-name= 
Example:
	postgres execute-raw-sql \
	  --sql="CREATE INDEX events_table_event_id ON public.events_table (event_id);"	\
	  --postgres-address="localhost:5432" \
	  --postgres-user="postgres" \
	  --postgres-password="XaD2sd$34Sdas1$ae" \
	  --postgres-database="postgres" 
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("→ Executing raw sql command")

		params := types.ExecuteRawSqlParams{
			Sql:  viper.GetString("execute-raw-sql"),
		}
		log.Println(params)

		postgresConfig := types.PostgresConfigParams{
			Addr:     viper.GetString("execute-raw-sql-postgres-address"),
			User:     viper.GetString("execute-raw-sql-postgres-user"),
			Password: viper.GetString("execute-raw-sql-postgres-password"),
			Database: viper.GetString("execute-raw-sql-postgres-database"),
		}

		action.ExecuteRawSql(params, postgresConfig)
		log.Println("→ Executing raw sql finished")
	},
}

