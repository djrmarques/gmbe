/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"fmt"
)

var version string

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
		flags := cmd.Flags()
		v, _ := flags.GetBool("version")

		if v {
			versionString := fmt.Sprintf("gmbe version %s", version)
			fmt.Println(versionString)
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
	rootCmd.Flags().Bool("version", false, "Display the version number")

}


