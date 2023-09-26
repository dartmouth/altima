package cmd

import (
	"altima/pkg/util"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type repoRemoveOptions struct {
	name string
}

// repoRemoveCmd represents the repoRemove command
var repoRemoveCmd = &cobra.Command{
	Use:   "remove NAME",
	Short: "Remove a module repositories",
	Long:  `Remove a module repositories`,
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
	// Verify input
	matched, err := regexp.MatchString("^[^-][a-zA-Z-]+$", o.name)
	util.CheckError(err)
	if !matched {
		fmt.Println("ERROR: Repo NAME must only contain letters, numbers, and dashes (can't start with dash)")
		os.Exit(1)
	}

	// Load current repositories
	viper.SetConfigName(settings.ConfigFilename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()

	// Delete repository index cache if it exists
	os.Remove(filepath.Join(settings.RepositoryCacheDir, o.name+".yaml"))

	// Remove repository
	repositories := viper.GetStringMap("repositories")
	delete(repositories, o.name)
	viper.Set("repositories", repositories)

	// Write updated config
	viper.WriteConfigAs(filepath.Join(settings.ConfigDir, settings.ConfigFilename))
	fmt.Println("Removed repository `" + o.name + "`")
}
