/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// repoListCmd represents the repoList command
var repoListCmd = &cobra.Command{
	Use:   "list",
	Short: "list module repositories",
	Long:  `list module repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		repoList()
	},
}

func init() {
	repoCmd.AddCommand(repoListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func repoList() {
	viper.SetConfigName(settings.RepositoryConfigFile)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"name", "url"})

	for k, _ := range viper.AllSettings() {
		table.Append([]string{k, viper.GetString(k + ".url")})
	}
	table.Render() // Send output
}
