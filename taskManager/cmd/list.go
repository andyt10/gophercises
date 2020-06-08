package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",
	Long:  "List tasks TODO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("LIST TASKS")
	},
}
