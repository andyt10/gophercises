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
		fmt.Println("LIST TASKS")

		tasks := src.GetAll()
		//loop and pretty print tasks
	},
}
