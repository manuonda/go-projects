package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "Tasks Todo List",
	Long:  "At Crud of tasks Todo List",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use --help to see available commands")
	},
}

func Execute() {
	todos.readFromFile()
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(completedCmd)
}
