package cmd

import (
	"fmt"
	"gophercises/task/db"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
	Use: "add",
	Short: "add task",
	Long: `add task to this todo application`,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added %q in your task list\n", task)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
