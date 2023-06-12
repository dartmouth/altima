/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

type initOptions struct {
}

var moduleReservedWords = []string{"enabled", "name", "version", "repo_name"}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [-]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 && args[0] == "-" {
			o := &initOptions{}
			o.run()
		} else {
			fmt.Printf(`
# For ZSH, appending the following to ~/.zshrc

  export ALTIMA_EXECUTABLE_DIR="%s"
  command -v altima >/dev/null || export PATH="$ALTIMA_EXECUTABLE_DIR:$PATH"
  eval "$(altima completion zsh)"
  eval "$(altima init -)"

# For BASH, appending the following to ~/.bash_profile

  export ALTIMA_EXECUTABLE_DIR="%s"
  command -v altima >/dev/null || export PATH="$ALTIMA_EXECUTABLE_DIR:$PATH"
  eval "$(altima completion bash)"
  eval "$(altima init -)"

`, settings.ExecutableDir, settings.ExecutableDir)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// # Restart your shell for the changes to take effect.

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (o *initOptions) run() {
	viper.SetConfigName(settings.ConfigFilename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()

	for m, _ := range viper.GetStringMap("modules") {
		fmt.Printf("\n# Loading module: %s\n", m)
		dat, err := os.ReadFile(filepath.Join(settings.ModulesDir, m, "init.sh"))
		check(err)
		rendered := string(dat)
		// Replace all standard variables
		rendered = strings.Replace(rendered, "${module_dir}", filepath.Join(settings.ModulesDir, m), 2)

		// Replace all configured variables
		for k, v := range viper.GetStringMap("modules." + m) {
			if !slices.Contains(moduleReservedWords, k) {
				rendered = strings.Replace(rendered, "${"+k+"}", fmt.Sprint(v), 2)
			}
		}

		// Display rendered containt
		fmt.Print(rendered)
	}
}
