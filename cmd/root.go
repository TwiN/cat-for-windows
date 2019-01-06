package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"io/ioutil"
	"strings"
)

var printLineNumber = false
var rootCmd = &cobra.Command{
	Use:   "cat FILE...",
	Short: "Concatenate file(s) and print the result to standard output",
	Long:  "A Go port of the popular 'cat' command used to concatenate file(s) to standard output.",
	Example: "cat file1.txt file2.txt",
	Version: "0.0.2",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var files []string
		for _, arg := range args {
			if strings.HasSuffix(arg, "-") {
			} else {
				files = append(files, arg)
			}
		}
		content := getFilesContent(files)
		if printLineNumber {
			for i, line := range strings.Split(content, "\n") {
				fmt.Printf("%4d  %s\n", i, line)
			}
		} else {
			fmt.Print(content)
		}
	},
}


func getFilesContent(files []string) string {
	result := ""
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			result += err.Error() + "\n"
		} else {
			result += string(content)
		}
	}
	return result
}


func init() {
	rootCmd.Flags().BoolVarP(&printLineNumber, "number", "n", false, "Print the number of each output line")
}


func Execute()  {
	rootCmd.Execute()
}
