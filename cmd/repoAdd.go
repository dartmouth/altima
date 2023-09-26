package cmd

import (
	"altima/pkg/repo"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type repoAddOptions struct {
	name string
	url  string
}

// repoAddCmd represents the repoAdd command
var repoAddCmd = &cobra.Command{
	Use:   "add NAME URL",
	Short: "Add a module repository",
	Long:  `Add a module repository`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := &repoAddOptions{}
		o.name = strings.ToLower(args[0])
		o.url = args[1]
		o.run()
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

func (o *repoAddOptions) run() {
	// Assert that repo name is only alpha-numberic with dashes
	matched, err := regexp.MatchString("^[^-][a-zA-Z-]+$", o.name)
	check(err)
	if !matched {
		fmt.Println("ERROR: Repo NAME must only contain letters, numbers, and dashes (can't start with dash)")
		os.Exit(1)
	}

	// Download and save repo index
	err = repo.DownloadIndexFile(o.name, o.url, settings.RepositoryCacheDir)
	check(err)

	// Write repo to config
	viper.SetConfigName(settings.ConfigFilename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()
	viper.Set("repositories."+o.name, map[string]string{"url": o.url})
	viper.WriteConfigAs(filepath.Join(settings.ConfigDir, settings.ConfigFilename))
	fmt.Println("Added repository `" + o.name + "` from `" + o.url + "`")
}
