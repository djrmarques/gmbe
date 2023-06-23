/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"github.com/djrmarques/go-md-block-extrator/extract"
	"path/filepath"
	"io/ioutil"
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
		if f_exists, _ := ValidateFileExists(f); !f_exists {
			log.Fatalf("File %s does not exist", f)
		}

		// Get blocks from files
		blocks, _ := extract.ExtractBlocksFromFile(f)

		// Builds the path to the output file
		base_dir, file_name := GetFileBaseFolderAndName(f)

		if o != "" {
			base_dir = o
		}

		full_output_path := filepath.Join(base_dir, file_name)

		// Save each block in the respective file
		counter_languages_block := make(map[string]uint)
		for _, b := range blocks {

			if val, ok := counter_languages_block[b.T]; ok {
				counter_languages_block[b.T] = val + 1
			} else {
				counter_languages_block[b.T] = 1
			}
			
			block_type_n := counter_languages_block[b.T]
			
			block_extention := LanguageExtensions[b.T]
			new_file_name := b.T + "_" + strconv.FormatUint(uint64(block_type_n), 10) + block_extention
			full_path_to_file := filepath.Join(full_output_path, new_file_name)

			err := ioutil.WriteFile(full_path_to_file, []byte(b.Content), 0644)
			if err != nil {
				log.Fatalf("Error writing to file: %s: %s", full_path_to_file, err)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.extractor.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("file", "f", "", "The file to parse")
	rootCmd.Flags().StringP("output-path", "o", "", "Where to store the code blocks.")
}


