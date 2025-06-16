package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "dcv",
	Short: "Starting Digital Cash Vault Application",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Starting Digital Cash Vault Application")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
