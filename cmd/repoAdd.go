/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type repoAddOptions struct {
	name string
	url  string
}

// repoAddCmd represents the repoAdd command
var repoAddCmd = &cobra.Command{
	Use:   "add [NAME] [URL]",
	Short: "add a module repository",
	Long:  `add a module repository`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := &repoAddOptions{}
		o.name = args[0]
		o.url = args[1]
		o.repoAdd()
	},
}

func init() {
	repoCmd.AddCommand(repoAddCmd)
	// settings = cli.EnvSettings

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (o *repoAddOptions) repoAdd() {
	// TODO: Download the module index for the new repository URL to confirm that it is well formed
	viper.SetConfigName(settings.RepositoryConfigFile)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()
	viper.Set(o.name, map[string]string{"url": o.url})
	viper.WriteConfigAs(filepath.Join(settings.ConfigDir, settings.RepositoryConfigFile))
	fmt.Println("Added repository `" + o.name + "` from `" + o.url + "`")
}
