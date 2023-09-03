package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// disableCmd represents the disable command
var disableCmd = &cobra.Command{
	Use:   "disable <module_name1> <module_name2> ...",
	Short: "Disables a module",
	Long:  "This command disables a module, so that it will not be initialized the next time `altima init` is called.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName(settings.ConfigFilename)
		viper.SetConfigType("toml")
		viper.AddConfigPath(settings.ConfigDir)
		viper.ReadInConfig()

		modules := viper.GetStringMap("modules")

		for _, module := range args {

			_, exists := modules[module]

			if !exists {
				fmt.Println(fmt.Errorf("Could not find module %q!", module))
			} else {
				viper.Set("modules."+module+".enabled", false)
			}
		}

		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(disableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// disableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// disableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
