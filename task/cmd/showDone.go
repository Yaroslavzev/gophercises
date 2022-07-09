/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/ledongthuc/goterators"
	"os"
	"strconv"
	"task/db"

	"github.com/spf13/cobra"
)

// showDoneCmd represents the showDone command
var showDoneCmd = &cobra.Command{
	Use:   "showDone",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		filteredTasks := goterators.Filter(tasks, func(item db.Task) bool {
			return item.Done
		})
		if len(filteredTasks) == 0 {
			fmt.Println("You have to do more today!")
			return
		}
		fmt.Println("You have done the following tasks:")
		for i, task := range filteredTasks {
			fmt.Printf("%d. %s, key: %d, done: %s\n", i+1, task.Value, task.Key, strconv.FormatBool(task.Done))
		}
	},
}

func init() {
	rootCmd.AddCommand(showDoneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showDoneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showDoneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
