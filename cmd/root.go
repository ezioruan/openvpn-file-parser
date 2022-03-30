/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/ezioruan/openvpn-file-parser/parser"
	"github.com/spf13/cobra"
)

var Desciption = "An Openvpn file Parse to Split .ovpn file to .ca .key and .cert files"
var inputFile string
var outputDir string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openvpn-file-parser",
	Short: Desciption,
	Long:  Desciption,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ParseFromFile with %s, will Split files to %s \n", inputFile, outputDir)
		config, err := parser.ParseFromFile(inputFile)
		if err != nil {
			fmt.Printf("ParseFromFile error %v \n", err)
			return
		}
		err = config.SplitFiles(outputDir)
		if err != nil {
			fmt.Printf("ParseFromFile error %v \n ", err)
			return
		}
		fmt.Printf("Split Files success ")
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
	rootCmd.Flags().StringVarP(&inputFile, "input", "i", "", "The Openvpn file path")
	rootCmd.Flags().StringVarP(&outputDir, "output", "o", "", "The Dir for the Split files")
}
