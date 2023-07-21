/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/djrmarques/gmbe/extract"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
	"strconv"
)

// Extracts all the blocks from a given file into various files
func extractToFile(f, o string, j bool) {
	// Check if a file exists

	if f == "" {
		log.Fatal("Please specify a file using the -f flag. ")
	}

	fExists, err := ValidateFileExists(f)
	if !fExists {
		log.Fatalf("File %s does not exist", f)
	}
	FailIfError(err)

	// Get blocks from files
	blocks, _ := extract.ExtractBlocksFromFile(f, j)

	// Builds the path to the output file
	baseDir, file_name := GetFileBaseFolderAndName(f)

	if o != "" {
		baseDir = o
	}

	outputFolder := filepath.Join(baseDir, file_name)
	FailIfError(CreateFolderIfNotExist(outputFolder))

	// Save each block in the respective file
	// Keeps track of how many blocks of each type
	// are in a file
	counterLanguesBlock := make(map[string]uint)
	for _, b := range blocks {

		if val, ok := counterLanguesBlock[b.T]; ok {
			counterLanguesBlock[b.T] = val + 1
		} else {
			counterLanguesBlock[b.T] = 1
		}

		blockTypeN := counterLanguesBlock[b.T]

		blockExtention := LanguageExtensions[b.T]
		outputFileName := b.T + "_" + strconv.FormatUint(uint64(blockTypeN), 10) + blockExtention
		outputFilePath := filepath.Join(outputFolder, outputFileName)
		CreateEmptyFile(outputFilePath)
		FailIfError(WriteStringToFile(outputFilePath, b.Content))
	}

}

// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Unpack flags
		flags := cmd.Flags()
		f, _ := flags.GetString("file")
		o, _ := flags.GetString("output-path")
		j, _ := flags.GetBool("join-blocks")

		extractToFile(f, o, j)
	},
}

func init() {
	rootCmd.AddCommand(extractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// extractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	extractCmd.Flags().StringP("file", "f", "", "The file to parse")
	extractCmd.Flags().BoolP("join-blocks", "j", false, "Join blocks of the same type on the same file")
	extractCmd.Flags().StringP("output-path", "o", "", "Where to store the code blocks.")
}
