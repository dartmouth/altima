package cmd

import (
	"altima/pkg/repo"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// installCmd represents the install command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls currently installed modules",
	Long:  `Uninstalls one or more currently installed modules. Modules must be referred to by their alias, if one has been set at installation.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		modules := repo.GetModulesFromString(args)

		viper.SetConfigName(settings.ConfigFilename)
		viper.SetConfigType("toml")
		viper.AddConfigPath(settings.ConfigDir)
		viper.ReadInConfig()

		modulesConfig := viper.GetStringMap("modules")

		for _, module := range modules {
			name := module.Name
			if module.Alias != "" {
				name = module.Alias
			}
			fmt.Printf("Uninstalling %q...", name)
			_, is_installed := modulesConfig[name]
			if !is_installed {
				fmt.Print("failed. ")
				msg := fmt.Sprintf("Module %q is not currently installed.", name)
				fmt.Println(fmt.Errorf(msg))
				continue
			}
			err := repo.UninstallModule(module, settings.ModulesDir)
			if err != nil {
				fmt.Print("failed. ")
				msg := fmt.Sprintf("An error occurred uninstalling module %q.", name)
				fmt.Println(fmt.Errorf(msg))
				continue
			}
			delete(modulesConfig, name)
			fmt.Println("success.")
		}
		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(uninstallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// installCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// installCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
