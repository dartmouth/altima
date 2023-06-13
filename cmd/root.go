/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"altima/pkg/cli"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var settings *cli.EnvSettings

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "altima",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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
	settings = cli.New()

	err := os.MkdirAll(settings.ConfigDir, os.ModePerm)
	check(err)
	err = os.MkdirAll(settings.CacheDir, os.ModePerm)
	check(err)
	err = os.MkdirAll(settings.ModulesDir, os.ModePerm)
	check(err)
	err = os.MkdirAll(settings.RepositoryCacheDir, os.ModePerm)
	check(err)

	if _, err := os.Stat(filepath.Join(settings.ConfigDir, settings.ConfigFilename)); os.IsNotExist(err) {
		_, err := os.Create(filepath.Join(settings.ConfigDir, settings.ConfigFilename))
		check(err)
	}

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.altima.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
