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
		fmt.Println("DO TASKS")
	},
}
