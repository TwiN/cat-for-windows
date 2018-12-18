package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"io/ioutil"
	"strings"
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
		var files []string
		for _, arg := range args {
			if strings.HasSuffix(arg, "-") {
			} else {
				files = append(files, arg)
			}
		}
		content := getFilesContent(files)
		if showLineNumber {
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

var showLineNumber = false

func init() {
	rootCmd.Flags().BoolVarP(&showLineNumber, "number", "n", false, "Print the number of each output line")
}

func Execute()  {
	rootCmd.Execute()
}

