/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/HrithikSawant/go-json-parser/internal/utils"
	"github.com/HrithikSawant/go-json-parser/lexer"
	"github.com/HrithikSawant/go-json-parser/parser"
	"github.com/spf13/cobra"
)

var filePath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-json-parser",
	Short: "A CLI tool for parsing and working with JSON in Go",
	Long: `go-json-parser is a CLI tool that reads a JSON file or standard input and determines whether it is syntactically valid.

Examples:
  go-json-parser myfile.json
  echo "{}" | go-json-parser

Output:
  DEBUG: State = Start                | Token = {          | Literal = {
  DEBUG: State = ExpectKeyOrEnd       | Token = }          | Literal = }
  DEBUG: State = Done                 | Token = EOF        | Literal = 
  Valid JSON structure`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },

	Run: func(cmd *cobra.Command, args []string) {
		var reader io.Reader
		var err error

		if len(args) > 0 {
			filePath := args[0]

			// Use utility function
			reader, err = utils.OpenFile(filePath)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				os.Exit(1)
			}

			if filepath.Ext(filePath) != ".json" {
				fmt.Fprintln(os.Stderr, "Error: File must have a .json extension")
				os.Exit(1)
			}
		} else if utils.IsInputFromPipe() {
			reader = bufio.NewReader(os.Stdin)
		} else {
			cmd.Help()
			return
		}

		// Read all input
		var builder strings.Builder
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			builder.WriteString(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		// Run lexer and parser
		input := builder.String()
		lex := lexer.NewLexer(input)
		parser := parser.NewParser(lex)

		if parser.Parse() {
			fmt.Println("Valid JSON structure")
		} else {
			fmt.Fprintln(os.Stderr, "Invalid JSON structure")
			os.Exit(1)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-json-parser.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
