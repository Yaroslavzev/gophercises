/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"fmt"
	"os"
	"task/cmd"
	"task/db"
)

func main() {
	must(db.Init())
	cmd.Execute()

}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
