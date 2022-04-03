/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"gophercises/task/db"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark the task complete",
	Long: `Mark the provided task complete from the stored tasks
	Example: task do <task no.>

task no. is the serial number you see as prefix when you run "task list"`,
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, 0, len(args))
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}
		tasks, err := db.AllTask()
		if err != nil {
			log.Fatal(err)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("invalid id:", id)
				continue
			}
			task := tasks[id - 1]
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark '%d' task as complete. (%v)\n", id, err)
			} else {
				fmt.Printf("Marked '%d' task as complete.\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
