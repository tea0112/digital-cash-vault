package cmd

import (
	"digital-cash-vault/config"
	"digital-cash-vault/internal/interfaces/migrations"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

func rootRunFunc(cmd *cobra.Command, args []string) {
	log.Println("Starting Digital Cash Vault Application")
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Start server
	log.Printf("Server running on port %s", appConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+appConfig.ServerPort, appConfig.Router)) // router satisfies http.Handler
}

func migrateRunFunc(cmd *cobra.Command, args []string) {
	log.Println("Migrate Digital Cash Vault Database")
	appConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	migrations.AutoMigrate(appConfig.DB)
}
