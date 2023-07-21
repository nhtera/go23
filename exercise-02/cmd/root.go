/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	mysort "github.com/nhtera/exercise-02/utils"
	"github.com/spf13/cobra"
)

var (
	inputType string
	input     []string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sorter",
	Short: "A CLI tool for sorting arrays",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		sorted := mysort.Sort(inputType, input)
		fmt.Println(sorted)
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.exercise-02.yaml)")
	rootCmd.PersistentFlags().StringVarP(&inputType, "type", "t", "", "input type (int, float, string, mix)")
	rootCmd.PersistentFlags().StringSliceVarP(&input, "input", "i", []string{}, "input array")
	// rootCmd.PersistentFlags().StringArrayVarP(&input, "input", "i", []string{}, "input array")
	rootCmd.MarkPersistentFlagRequired("type")
	rootCmd.MarkPersistentFlagRequired("input")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
