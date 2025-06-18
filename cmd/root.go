package cmd

import (
	"digital-cash-vault/config"
	"digital-cash-vault/migrations"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "dcv",
	Short: "Starting Digital Cash Vault Application",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting Digital Cash Vault Application")
		_, err := config.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}
	},
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Migrate Digital Cash Vault Database")
		appConfig, err := config.LoadConfig()
		if err != nil {
			log.Fatal(err)
		}

		migrations.AutoMigrate(appConfig.DB)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
