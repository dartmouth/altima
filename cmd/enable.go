/*
Copyright © 2023 Simon Stone <simon.stone@dartmouth.edu>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// enableCmd represents the enable command
var enableCmd = &cobra.Command{
	Use:   "enable <module_name1> <module_name2> ...",
	Short: "Enables a module",
	Long:  "This command enables a module, so that it will be initialized the next time `altima init` is called.",
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
				viper.Set("modules."+module+".enabled", true)
			}
		}

		viper.WriteConfig()
	},
}

func init() {
	rootCmd.AddCommand(enableCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
