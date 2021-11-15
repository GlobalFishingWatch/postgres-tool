package cmd

import (
	"github.com/GlobalFishingWatch/postgres-tool/internal/action"
	"github.com/GlobalFishingWatch/postgres-tool/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func init() {
	deleteViewCmd.Flags().StringP("postgres-address", "", "", "The address of the database")
	deleteViewCmd.MarkFlagRequired("postgres-address")
	deleteViewCmd.Flags().StringP("postgres-user", "", "", "The destination credentials user")
	deleteViewCmd.MarkFlagRequired("postgres-user")
	deleteViewCmd.Flags().StringP("postgres-password", "", "", "The destination credentials password")
	deleteViewCmd.MarkFlagRequired("postgres-password")
	deleteViewCmd.Flags().StringP("postgres-database", "", "", "The destination database name")
	deleteViewCmd.MarkFlagRequired("postgres-database")

	deleteViewCmd.Flags().StringP("view-name", "", "", "The name of the view to delete")
	deleteViewCmd.MarkFlagRequired("view-name")

	viper.BindPFlag("delete-view-postgres-address", deleteViewCmd.Flags().Lookup("postgres-address"))
	viper.BindPFlag("delete-view-postgres-user", deleteViewCmd.Flags().Lookup("postgres-user"))
	viper.BindPFlag("delete-view-postgres-password", deleteViewCmd.Flags().Lookup("postgres-password"))
	viper.BindPFlag("delete-view-postgres-database", deleteViewCmd.Flags().Lookup("postgres-database"))
	viper.BindPFlag("delete-view-view-name", deleteViewCmd.Flags().Lookup("view-name"))

	rootCmd.AddCommand(deleteViewCmd)
}

var deleteViewCmd = &cobra.Command{
	Use:   "delete-view",
	Short: "delete a view",
	Long:  `Delete new view
Format:
	postgres delete-view --postgres-address= --postgres-user= --postgres-password= --postgres-database= --view-name= 
Example:
	postgres delete-view \
	  --view-name="vessels"	\
	  --postgres-address="localhost:5432" \
	  --postgres-user="postgres" \
	  --postgres-password="XaD2sd$34Sdas1$ae" \
	  --postgres-database="postgres" 
`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("→ Executing delete view command")

		params := types.DeleteViewParams{
			ViewName:  viper.GetString("delete-view-view-name"),
		}

		postgresConfig := types.PostgresConfigParams{
			Addr:     viper.GetString("delete-view-postgres-address"),
			User:     viper.GetString("delete-view-postgres-user"),
			Password: viper.GetString("delete-view-postgres-password"),
			Database: viper.GetString("delete-view-postgres-database"),
		}

		action.DeleteView(params, postgresConfig)
		log.Println("→ Executing delete view command finished")
	},
}

