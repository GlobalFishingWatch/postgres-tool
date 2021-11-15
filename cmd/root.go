package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"os"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "postgres-tool",
	Short: "A CLI to manage actions using Postgres",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

