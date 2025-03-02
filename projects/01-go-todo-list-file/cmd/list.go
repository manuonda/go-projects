package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Task list show a list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List of Tasks")
		todos.list()
	},
}
