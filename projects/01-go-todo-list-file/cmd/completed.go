package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completedCmd = &cobra.Command{
	Use:   "completed 'index task'",
	Short: "Completed Task change status of task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error read index : ", err.Error)
		}
		err = todos.completed(index)
		if err != nil {
			fmt.Println("Error completed todo", err.Error())
		} else {
			fmt.Println("Completed todo")
		}
	},
}
