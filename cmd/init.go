package cmd

import (
	"altima/pkg/util"
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
	Short: "Initialize altima into the environment",
	Long: `This command generates the commands to activate altima in your shell
as well as generating the module shell components when '-' is provided.
`,
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
		// Check if module enabled
		if viper.GetBool("modules." + m + ".enabled") {
			fmt.Printf("\n# Loading module: %s\n", m)
			dat, err := os.ReadFile(filepath.Join(settings.ModulesDir, m, "init.sh"))
			util.CheckError(err)
			rendered := string(dat)
			// Replace all standard variables
			rendered = strings.ReplaceAll(rendered, "${module_dir}", filepath.Join(settings.ModulesDir, m))
			rendered = strings.ReplaceAll(rendered, "${module_name}", m)
			rendered = strings.ReplaceAll(rendered, "${altima_config_path}", filepath.Join(settings.ConfigDir, settings.ConfigFilename))

			// Replace all configured variables
			for k, v := range viper.GetStringMap("modules." + m) {
				if !slices.Contains(moduleReservedWords, k) {
					rendered = strings.ReplaceAll(rendered, "${"+k+"}", fmt.Sprint(v))
				}
			}
			// Display rendered containt
			fmt.Print(rendered)
		}
	}
}
