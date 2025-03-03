package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete 'index' ",
	Short: "Delete task ",
	Long:  "Delete task of the table",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid index: " + args[0])
			return
		}
		err = todos.delete(index)
		if err != nil {
			fmt.Println("Error delete: ", err.Error())
		} else {
			fmt.Println("Delete todo at index :", index)
		}
	},
}
