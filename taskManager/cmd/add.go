package cmd

import (
	"cor_gophercises/taskManager/src"
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

		task := src.ListItem{
			Item:   args[0],
			DoneAt: 0,
		}
		fmt.Printf("Trying to add task: %v \n", task)
		src.Add(task)
	},
}
