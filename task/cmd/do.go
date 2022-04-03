/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark the task complete",
	Long: `Mark the provided task complete from the stored tasks
		Example: task do <task>`,
	Run: func(cmd *cobra.Command, args []string) {
		ids := make([]int, 0, len(args))
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				panic(err)
			}
			ids = append(ids, id)
		}
		fmt.Println(ids)
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
