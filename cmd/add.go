/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/go-cli/bi/bill"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new bill",
	Long:  `Add will create a new bil item on to the list`,
	Run:   addRun,
}

var value int
var date string

func addRun(cmd *cobra.Command, args []string) {
	items, err := bill.ReadItems(dataFile)
	if err != nil {
		fmt.Printf("Error reading file: %s\n error: %s\n", dataFile, err)
	}
	for _, x := range args {
		item := bill.Item{Text: x, Value: value, Date: date}
		items = append(items, item)
	}

	if err := bill.SaveItems(dataFile, items); err != nil {
		fmt.Println("Error: " + err.Error())
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&value, "value", "v", 0, "Set the value to the bill")
	addCmd.Flags().StringVarP(&date, "date", "d", "", "Set the date to add to the bill")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
