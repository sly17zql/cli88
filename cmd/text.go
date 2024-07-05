/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cli88/convert"
	"fmt"
	"github.com/spf13/cobra"
)

var inputPdfPath string
var outputMDName string

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := convert.PdfToMarkdown(inputPdfPath, outputMDName); err != nil {
			fmt.Println("transfer error")
			fmt.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(textCmd)

	transferImageSimple.Flags().StringVarP(&inputPdfPath, "input", "i", "", "Path to the input pdf file")
	transferImageSimple.Flags().StringVarP(&outputMDName, "output", "o", "", "Path to the output md file")

}
