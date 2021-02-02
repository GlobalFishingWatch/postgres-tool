package cmd

import (
	"github.com/GlobalFishingWatch/postgres-tool/internal/action"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	deleteTableCmd.Flags().StringP("postgres-address", "", "", "The address of the database")
	deleteTableCmd.MarkFlagRequired("postgres-address")
	deleteTableCmd.Flags().StringP("postgres-user", "", "", "The destination credentials user")
	deleteTableCmd.MarkFlagRequired("postgres-user")
	deleteTableCmd.Flags().StringP("postgres-password", "", "", "The destination credentials password")
	deleteTableCmd.MarkFlagRequired("postgres-password")
	deleteTableCmd.Flags().StringP("postgres-database", "", "", "The destination database name")
	deleteTableCmd.MarkFlagRequired("postgres-database")

	deleteTableCmd.Flags().StringP("table-name", "", "", "The name of the table to delete")
	deleteTableCmd.MarkFlagRequired("table-name")

	viper.BindPFlag("delete-table-postgres-address", deleteTableCmd.Flags().Lookup("postgres-address"))
	viper.BindPFlag("delete-table-postgres-user", deleteTableCmd.Flags().Lookup("postgres-user"))
	viper.BindPFlag("delete-table-postgres-password", deleteTableCmd.Flags().Lookup("postgres-password"))
	viper.BindPFlag("delete-table-postgres-database", deleteTableCmd.Flags().Lookup("postgres-database"))
	viper.BindPFlag("delete-table-table-name", deleteTableCmd.Flags().Lookup("table-name"))

	rootCmd.AddCommand(deleteTableCmd)
}

var deleteTableCmd = &cobra.Command{
	Use:   "delete-table",
	Short: "delete a table",
	Long:  `Delete new table
Format:
	postgres delete-table --postgres-address= --postgres-user= --postgres-password= --postgres-database= --table-name= 
Example:
	postgres delete-table \
	  --table-name="vessels"	\
	  --postgres-address="localhost:5432" \
	  --postgres-user="postgres" \
	  --postgres-password="XaD2sd$34Sdas1$ae" \
	  --postgres-database="postgres" 
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("→ Executing delete table command")

		params := types.DeleteTableParams{
			TableName:  viper.GetString("delete-table-table-name"),
		}

		postgresConfig := types.PostgresConfigParams{
			Addr:     viper.GetString("delete-table-postgres-address"),
			User:     viper.GetString("delete-table-postgres-user"),
			Password: viper.GetString("delete-table-postgres-password"),
			Database: viper.GetString("delete-table-postgres-database"),
		}

		action.DeleteTable(params, postgresConfig)
		log.Println("→ Executing delete table command finished")
	},
}

