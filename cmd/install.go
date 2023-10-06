package cmd

import (
	"altima/pkg/repo"
	"altima/pkg/util"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs modules from the module index",
	Long:  `Installs one or more modules from the module index. You can optionally supply a version number and an alias (to avoid name collision).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		modules := repo.GetModulesFromString(args)

		viper.SetConfigName(settings.ConfigFilename)
		viper.SetConfigType("toml")
		viper.AddConfigPath(settings.ConfigDir)
		viper.ReadInConfig()

		modulesConfig := viper.GetStringMap("modules")

		for _, module := range modules {
			if module.Version != "" {
				fmt.Printf("Searching URL for module %q in version %q...\n", module.Name, module.Version)
			} else {
				fmt.Printf("Searching URL for module %q...\n", module.Name)
			}

			module, err := repo.Search(module, settings.RepositoryCacheDir)

			util.CheckError(err)

			fmt.Printf("Found URL %q...\n", module.Url)

			installName := module.Name
			fmt.Printf("Installing module %q...\n", installName)

			if module.Alias != "" {
				installName = module.Alias
				fmt.Printf("Using install name %q...\n", installName)
			}

			err = repo.InstallModule(module, settings.ModulesDir)

			if err != nil {
				msg := fmt.Sprintf("Failed to install module %q", installName)
				if module.Version != "" {
					msg += fmt.Sprintf(" in version %q", module.Version)
				}
				msg += "!"
				fmt.Println(fmt.Errorf(msg))
				continue
			}

			fmt.Println("Updating altima config...")

			moduleConfigFile := filepath.Join(settings.ModulesDir, installName, "default_config.toml")

			buf, err := os.ReadFile(moduleConfigFile)
			if err != nil {
				fmt.Println("failed to open default config file!")
				continue
			}

			var newConfig map[string]any
			toml.Unmarshal(buf, &newConfig)

			newConfig["name"] = module.Name
			newConfig["version"] = module.Version
			newConfig["repo_name"] = module.Repo
			newConfig["enabled"] = true

			modulesConfig[installName] = newConfig

			fmt.Printf("Module %q installed successfully.\n", installName)
		}
		viper.Set("modules", modulesConfig)
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(installCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
