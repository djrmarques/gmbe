/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/djrmarques/go-md-block-extrator/extract"
	"path/filepath"
	"strconv"
	"log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "extractor",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// Unpack flags
		flags := cmd.Flags()
		f, _ := flags.GetString("file")
		o, _ := flags.GetString("output-path")

		// Check if a file exists
		fExists, err := ValidateFileExists(f)
		if !fExists {
			log.Fatalf("File %s does not exist", f)
		}
		FailIfError(err)

		// Get blocks from files
		blocks, _ := extract.ExtractBlocksFromFile(f)

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
	rootCmd.Flags().StringP("output-path", "o", "", "Where to store the code blocks.")
}


