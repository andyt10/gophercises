package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do tasks",
	Long:  "Mark Task as done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			errorExit("Only one argument should be passed to do CMD")
		}

		task := args[0]
		fmt.Printf("Add Task: %v\n", task)
	},
}
