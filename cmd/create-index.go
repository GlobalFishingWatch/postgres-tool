package cmd

import (
	"github.com/GlobalFishingWatch/postgres-tool/internal/action"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	createIndexCmd.Flags().StringP("postgres-address", "", "", "The address of the database")
	createIndexCmd.MarkFlagRequired("postgres-address")
	createIndexCmd.Flags().StringP("postgres-user", "", "", "The destination credentials user")
	createIndexCmd.MarkFlagRequired("postgres-user")
	createIndexCmd.Flags().StringP("postgres-password", "", "", "The destination credentials password")
	createIndexCmd.MarkFlagRequired("postgres-password")
	createIndexCmd.Flags().StringP("postgres-database", "", "", "The destination database name")
	createIndexCmd.MarkFlagRequired("postgres-database")

	createIndexCmd.Flags().StringP("table-name", "", "", "The name of the table to add a index")
	createIndexCmd.MarkFlagRequired("table-name")
	createIndexCmd.Flags().StringP("index-name", "", "", "The name of the new index")
	createIndexCmd.MarkFlagRequired("index-name")
	createIndexCmd.Flags().StringP("column", "", "", "The name of the column")
	createIndexCmd.MarkFlagRequired("column")

	viper.BindPFlag("create-index-postgres-address", createIndexCmd.Flags().Lookup("postgres-address"))
	viper.BindPFlag("create-index-postgres-user", createIndexCmd.Flags().Lookup("postgres-user"))
	viper.BindPFlag("create-index-postgres-password", createIndexCmd.Flags().Lookup("postgres-password"))
	viper.BindPFlag("create-index-postgres-database", createIndexCmd.Flags().Lookup("postgres-database"))
	viper.BindPFlag("create-index-table-name", createIndexCmd.Flags().Lookup("table-name"))
	viper.BindPFlag("create-index-index-name", createIndexCmd.Flags().Lookup("index-name"))
	viper.BindPFlag("create-index-column", createIndexCmd.Flags().Lookup("column"))

	rootCmd.AddCommand(createIndexCmd)
}

var createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Add a index",
	Long:  `Add a index
Format:
	postgres create-index --postgres-address= --postgres-user= --postgres-password= --postgres-database= --table-name= 
Example:
	postgres create-index \
	  --view-name="vessels"	\
	  --postgres-address="localhost:5432" \
	  --postgres-user="postgres" \
	  --postgres-password="XaD2sd$34Sdas1$ae" \
	  --postgres-database="postgres" 
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("→ Executing delete view command")

		params := types.CreateIndexParam{
			TableName:  viper.GetString("create-index-table-name"),
			IndexName:  viper.GetString("create-index-index-name"),
			Column:  viper.GetString("create-index-column"),
		}
		log.Println(params)
		postgresConfig := types.PostgresConfigParams{
			Addr:     viper.GetString("create-index-postgres-address"),
			User:     viper.GetString("create-index-postgres-user"),
			Password: viper.GetString("create-index-postgres-password"),
			Database: viper.GetString("create-index-postgres-database"),
		}

		action.CreateIndex(params, postgresConfig)
		log.Println("→ Executing delete view command finished")
	},
}

