/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/SeanBracksDev/go-grep/internal/grep"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-grep",
	Short: "Search for a string in a file(s)",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lineNumbers, err := cmd.Flags().GetBool("line-numbers")
		if err != nil {
			panic(err)
		}

		var opts []grep.Option
		if lineNumbers {
			opts = append(opts, grep.WithLineNumbers())
		}

		// check if there is somethinig to read from STDIN
		searchString := args[0]
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			var stdin []byte
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				currentLine := append(scanner.Bytes(), []byte("\n")...)
				stdin = append(stdin, currentLine...)
			}
			if err := scanner.Err(); err != nil {
				panic(err)
			}
			grep.Search(bytes.NewReader(stdin), searchString, opts...)
		} else {
			filesToSearch := args[1:]
			for _, file := range filesToSearch {
				isDir, err := grep.IsDir(file)
				if err != nil {
					if os.IsNotExist(err) {
						fmt.Printf("%s: No such file or directory\n", file)
					} else {
						fmt.Println("Error:", err)
					}
				}
				if isDir {
					fmt.Printf("%s: Is a directory\n", file)
				}

				fileReader, err := os.Open(file)
				if err != nil {
					fmt.Println("Error:", err)
					break
				}
				defer fileReader.Close()

				grep.Search(fileReader, searchString, append(opts, grep.WithFilePath(file))...)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-grep.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("line-numbers", "n", false, "Show line numbers")
}
