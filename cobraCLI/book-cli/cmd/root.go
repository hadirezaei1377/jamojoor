package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "book-cli",
	Short: "Book CLI is a simple tool to manage your books",
	Long:  `A simple CLI tool for adding, removing, listing, and searching books.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
