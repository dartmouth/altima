/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"altima/pkg/repo"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type updateOptions struct{}

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		o := &updateOptions{}
		o.run()
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (o *updateOptions) run() {
	viper.SetConfigName(settings.ConfigFilename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"name", "url", "status"})

	for k, v := range viper.GetStringMap("repositories") {
		// Download and save repo index
		err := repo.DownloadIndexFile(string(k), v.(map[string]interface{})["url"].(string), settings.RepositoryCacheDir)
		if err == nil {
			table.Append([]string{string(k), v.(map[string]interface{})["url"].(string), "Updated"})
		} else {
			table.Append([]string{string(k), v.(map[string]interface{})["url"].(string), "Error"})
		}
	}
	table.Render() // Send output
}
