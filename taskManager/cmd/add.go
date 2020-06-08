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
		fmt.Println("ADD TASKS")
	},
}
