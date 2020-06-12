package cmd

import (
	"cor_gophercises/taskManager/src"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "How To",
	Short: "This is a short something",
	Long:  "A looooong something",
}

func Execute() error {
	src.InitDb("")
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(addCmd)
}

func errorExit(msg string) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
