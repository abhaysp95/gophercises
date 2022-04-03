package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
	Use: "add",
	Short: "add task",
	Long: `add task to this todo application`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added %q to your task list\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
