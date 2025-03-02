package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add 'new task' ",
	Short: "Add new Task",
	Long:  "Add a new Task to the list",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := args[0]
		todos.add(title)
		fmt.Println("Added todo : ", title)
	},
}
