package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dcv",
	Short: "Starting Digital Cash Vault Application",
	Run:   rootRunFunc,
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database to latest version",
	Run:   migrateRunFunc,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
