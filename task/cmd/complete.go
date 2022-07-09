/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ledongthuc/goterators"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// comleteCmd represents the comlete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()

		filteredTasks := goterators.Filter(tasks, func(item db.Task) bool {
			return !item.Done
		})

		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(filteredTasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			task := filteredTasks[id-1]
			id, err := db.DoneTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to done \"%d\". Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as done.\n", id)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
