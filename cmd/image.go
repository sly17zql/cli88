package cmd

import (
	"cli88/image"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var inputPath string
var originalName string
var outputPath string
var format string

// imageCmd represents the image command
var transferImageSimple = &cobra.Command{
	Use:   "transferImageSimple",
	Short: "单个转换图片",
	Long:  `单个转换图片`,
	Run: func(cmd *cobra.Command, args []string) {
		format := strings.ToLower(format)
		if err := image.HeicTransferSimple(inputPath, outputPath, format); err != nil {
			fmt.Println(err)
		}

	},
}

var transferImageMultiply = &cobra.Command{
	Use:   "transferImageMultiply",
	Short: "批量转换图片",
	Long:  `批量转换图片`,
	Run: func(cmd *cobra.Command, args []string) {
		format := strings.ToLower(format)
		if err := image.TransferMultiply(inputPath, originalName, outputPath, format); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(transferImageSimple)
	rootCmd.AddCommand(transferImageMultiply)

	transferImageSimple.Flags().StringVarP(&inputPath, "input", "i", "", "Path to the input HEIC file")
	transferImageSimple.Flags().StringVarP(&outputPath, "output", "o", "", "Path to the output image file")
	transferImageSimple.Flags().StringVarP(&format, "format", "f", "png", "Output format: jpeg or png")

	transferImageMultiply.Flags().StringVarP(&inputPath, "input", "i", "", "Path to the input HEIC file")
	transferImageMultiply.Flags().StringVarP(&originalName, "originalName", "g", "", "pattern file name start")
	transferImageMultiply.Flags().StringVarP(&outputPath, "output", "o", "", "Path to the output image file")
	transferImageMultiply.Flags().StringVarP(&format, "format", "f", "png", "Output format: jpeg or png")
}
