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

type repoRemoveOptions struct {
	name string
}

// repoRemoveCmd represents the repoRemove command
var repoRemoveCmd = &cobra.Command{
	Use:   "remove [NAME]",
	Short: "remove one or more module repositories",
	Long:  `remove one or more module repositories`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		o := &repoRemoveOptions{}
		o.name = args[0]
		o.run()
	},
}

func init() {
	repoCmd.AddCommand(repoRemoveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoRemoveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoRemoveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (o *repoRemoveOptions) run() {
	// TODO: Purge cache for the repo being removed

	// Load current repositories
	viper.SetConfigName(settings.ConfigFilename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()

	// Remove repository
	repositories := viper.GetStringMap("repositories")
	delete(repositories, o.name)
	viper.Set("repositories", repositories)

	// Write file
	viper.WriteConfigAs(filepath.Join(settings.ConfigDir, settings.ConfigFilename))
	fmt.Println("Removed repository `" + o.name + "`")
}
