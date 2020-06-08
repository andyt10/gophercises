package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "How To",
	Short: "This is a short something",
	Long:  "A looooong something",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(doCmd)
	rootCmd.AddCommand(addCmd)
}
