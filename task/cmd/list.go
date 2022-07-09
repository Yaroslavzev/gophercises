/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task/db"

	"github.com/ledongthuc/goterators"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show list of tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		filteredTasks := goterators.Filter(tasks, func(item db.Task) bool {
			return !item.Done
		})
		if len(filteredTasks) == 0 {
			fmt.Println("You have no tasks to complete! Why not take a vacation? 🏖")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range filteredTasks {
			fmt.Printf("%d. %s, key: %d, done: %s\n", i+1, task.Value, task.Key, strconv.FormatBool(task.Done))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
