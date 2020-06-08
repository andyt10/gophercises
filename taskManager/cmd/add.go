package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add task",
	Long:  "Add a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errorExit("Only 1 argument should be passed to CMD")
		}

		task := args[0]
		fmt.Printf("Adding new task: %v\n", task)
	},
}
