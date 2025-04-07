/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/go-cli/bi/bill"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Sholl all the bills",
	Long:  `Shows all the bills that have been added to the list`,
	Run:   listRun,
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

func listRun(cmd *cobra.Command, args []string) {
	items, err := bill.ReadItems(dataFile)
	if err != nil {
		fmt.Println(err)
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	fmt.Fprintln(w, "Conta\tValor\tData")
	for _, i := range items {
		_, _ = fmt.Fprintln(w, i.Text+"\t"+i.PrettyV()+"\t"+i.Date)
	}

	if err = w.Flush(); err != nil {
		fmt.Println(err)
	}
}
