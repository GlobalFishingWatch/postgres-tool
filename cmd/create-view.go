package cmd

import (
	"github.com/GlobalFishingWatch/postgres-tool/internal/action"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	createViewCmd.Flags().StringP("postgres-address", "", "", "The address of the database")
	createViewCmd.MarkFlagRequired("postgres-address")
	createViewCmd.Flags().StringP("postgres-user", "", "", "The destination credentials user")
	createViewCmd.MarkFlagRequired("postgres-user")
	createViewCmd.Flags().StringP("postgres-password", "", "", "The destination credentials password")
	createViewCmd.MarkFlagRequired("postgres-password")
	createViewCmd.Flags().StringP("postgres-database", "", "", "The destination database name")
	createViewCmd.MarkFlagRequired("postgres-database")

	createViewCmd.Flags().StringP("view-name", "", "", "The name of the view to associate the table")
	createViewCmd.MarkFlagRequired("view-name")
	createViewCmd.Flags().StringP("table-name", "", "", "The name of the table to associate the view")
	createViewCmd.MarkFlagRequired("table-name")

	viper.BindPFlag("create-view-postgres-address", createViewCmd.Flags().Lookup("postgres-address"))
	viper.BindPFlag("create-view-postgres-user", createViewCmd.Flags().Lookup("postgres-user"))
	viper.BindPFlag("create-view-postgres-password", createViewCmd.Flags().Lookup("postgres-password"))
	viper.BindPFlag("create-view-postgres-database", createViewCmd.Flags().Lookup("postgres-database"))
	viper.BindPFlag("create-view-view-name", createViewCmd.Flags().Lookup("view-name"))
	viper.BindPFlag("create-view-table-name", createViewCmd.Flags().Lookup("table-name"))

	rootCmd.AddCommand(createViewCmd)
}

var createViewCmd = &cobra.Command{
	Use:   "create-view",
	Short: "Create a new view",
	Long:  `Create new view
Format:
	postgres create-view --postgres-address= --postgres-user= --postgres-password= --postgres-database= --view-name= --table-name=
Example:
	postgres create-view \
	  --view-name="vessels"	\
	  --table-name="vessels_2021_02_01" \
	  --postgres-address="localhost:5432" \
	  --postgres-user="postgres" \
	  --postgres-password="XaD2sd$34Sdas1$ae" \
	  --postgres-database="postgres" 
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("→ Executing Create view command")

		params := types.CreateViewParams{
			TableName: viper.GetString("create-view-table-name"),
			ViewName:  viper.GetString("create-view-view-name"),
		}

		postgresConfig := types.PostgresConfigParams{
			Addr:     viper.GetString("create-view-postgres-address"),
			User:     viper.GetString("create-view-postgres-user"),
			Password: viper.GetString("create-view-postgres-password"),
			Database: viper.GetString("create-view-postgres-database"),
		}

		action.CreateView(params, postgresConfig)
		log.Println("→ Executing Create view command finished")
	},
}

