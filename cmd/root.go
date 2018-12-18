package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"io/ioutil"
)


var rootCmd = &cobra.Command{
	Use:   "cat",
	Short: "Concatenate file(s) and print the result to standard output",
	Long:  "A Go port of the popular 'cat' command used to concatenate file(s) to standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("USAGE: cat <file>...")
			os.Exit(1)
		}
		printFiles(args)
	},
}


func printFiles(files []string)  {
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(content))
	}
}


func Execute()  {
	rootCmd.Execute()
}

