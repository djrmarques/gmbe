/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"github.com/spf13/cobra"
)


// extractCmd represents the extract command
var extractCmd = &cobra.Command{
	Use:   "extract",
	Short: "Extracts the code blocks from a certain file",
	Long: `Use this command to extract code blocks from various markdown blocks.
By default, this command will create a folder with the same name of the file where it will store the code blocks.
It is also possible to specify an output directory, and to group all code blocks of the same type on the same file.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Unpack flags
		flags := cmd.Flags()
		f, _ := flags.GetString("file")
		o, _ := flags.GetString("output-path")
		j, _ := flags.GetBool("join-blocks")

		err := extractToFile(f, o, j)

		if err != nil {
			log.Fatal(err)
		}
		
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
