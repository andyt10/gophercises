package cmd

import (
	"cor_gophercises/taskManager/src"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  "List tasks TODO",
	Run: func(cmd *cobra.Command, args []string) {

		allTasks := src.GetAll()
		activeTasks := filterActive(allTasks)

		fmt.Printf("--- Found %v Active Tasks ---\n", len(activeTasks))

		for i, v := range activeTasks {
			fmt.Printf("TASK:%v --- %v --- Added: %v\n", i+1, v.Data.Item, v.Data.Added)
		}
	},
}
