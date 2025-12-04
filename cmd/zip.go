/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var filePaths []string

// zipCmd represents the zip command
var zipCmd = &cobra.Command{
	Use:   "zip",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(filePaths) == 0 {
			fmt.Println("No files specified. Use --files to provide a list of files.")
			return
		}
		resultFileName := "result.zip"
		if len(args) > 0 {
			resultFileName = args[0]
		}
		fmt.Println(resultFileName)
		fmt.Println("creating zip archive...")
		archive, err := os.Create(resultFileName)
		if err != nil {
			panic(err)
		}
		defer archive.Close()
		zipWriter := zip.NewWriter(archive)

		fmt.Printf("Processing files: %s\n", strings.Join(filePaths, ", "))
		// Add your file processing logic here
		for _, filePath := range filePaths {
			fmt.Printf("  processing: %s\n", filePath)
			fileName := filepath.Base(filePath)
			f1, err := os.Open(filePath)
			if err != nil {
				panic(err)
			}
			defer f1.Close()

			fmt.Printf("  writing file %s to archive...\n", filePath)
			w1, err := zipWriter.Create("result/" + fileName)
			if err != nil {
				panic(err)
			}
			if _, err := io.Copy(w1, f1); err != nil {
				panic(err)
			}
		}

		fmt.Println("closing zip archive...")
		zipWriter.Close()
	},
}

func init() {
	rootCmd.AddCommand(zipCmd)
	// Define the 'files' flag as a StringSlice
	zipCmd.Flags().StringSliceVarP(&filePaths, "files", "f", []string{}, "Comma-separated list of files to process")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// zipCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// zipCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
