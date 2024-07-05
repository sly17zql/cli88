package cmd

import (
	"cli88/convert"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var inputHeicPath string
var originalName string
var outputImgPath string
var format string

// imageCmd represents the convert command
var transferImageSimple = &cobra.Command{
	Use:   "transferImageSimple",
	Short: "单个转换图片",
	Long:  `单个转换图片`,
	Run: func(cmd *cobra.Command, args []string) {
		format := strings.ToLower(format)
		if err := convert.HeicTransferSimple(inputHeicPath, outputImgPath, format); err != nil {
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
		if err := convert.TransferMultiply(inputHeicPath, originalName, outputImgPath, format); err != nil {
			fmt.Println(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(transferImageSimple)
	rootCmd.AddCommand(transferImageMultiply)

	transferImageSimple.Flags().StringVarP(&inputHeicPath, "input", "i", "", "Path to the input HEIC file")
	transferImageSimple.Flags().StringVarP(&outputImgPath, "output", "o", "", "Path to the output convert file")
	transferImageSimple.Flags().StringVarP(&format, "format", "f", "png", "Output format: jpeg or png")

	transferImageMultiply.Flags().StringVarP(&inputHeicPath, "input", "i", "", "Path to the input HEIC file")
	transferImageMultiply.Flags().StringVarP(&originalName, "originalName", "g", "", "pattern file name start")
	transferImageMultiply.Flags().StringVarP(&outputImgPath, "output", "o", "", "Path to the output convert file")
	transferImageMultiply.Flags().StringVarP(&format, "format", "f", "png", "Output format: jpeg or png")
}
