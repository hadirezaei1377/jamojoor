package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new book",
	Long:  `Add a new book with a specified title and author.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("Please provide the book title and author.")
			return
		}
		title := args[0]
		author := args[1]
		books.AddBook(title, author)
		fmt.Println("Book added successfully.")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
