/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/djrmarques/gmbe/extract"
	"path/filepath"
	"strconv"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gmbe -f path/to/file.md",
	Short: "Extract code blocks from Markdown files",
	Long: `
Go Markdown Block Extractor will extract all the codeblocks from markdown files and save them as their own separate files.
This allows to run checks like linting, formating or any other custom checks to make sure that the code blocks in the markdown files, which most likely are some sort of documentation, are actually valid.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Unpack flags
		flags := cmd.Flags()
		f, _ := flags.GetString("file")
		o, _ := flags.GetString("output-path")
		j, _ := flags.GetBool("join-blocks")

		// Check if a file exists
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.extractor.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "f", "", "The file to parse")
	rootCmd.Flags().BoolP("join-blocks", "j", false, "Join blocks of the same type on the same file")
	rootCmd.Flags().StringP("output-path", "o", "", "Where to store the code blocks.")
}


