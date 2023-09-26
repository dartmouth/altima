/*
Copyright © 2023 Simon Stone <simon.stone@dartmouth.edu>
*/
package cmd

import (
	"altima/pkg/util"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Changes a value in the configuration",
	Long: `This command changes a value in the configuration.

	If the specified key does not exist in the file, an error is raised.
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		viper.SetConfigName(settings.ConfigFilename)
		viper.SetConfigType("toml")
		viper.AddConfigPath(settings.ConfigDir)
		viper.ReadInConfig()

		key := args[0]

		if !viper.InConfig(key) {
			fmt.Println(fmt.Errorf("Could not find the key %q in the config file %q!", key, settings.ConfigFilename))
			return
		}

		val := util.DeduceType(args[1])

		viper.Set(key, val)
		viper.WriteConfig()
	},
}

func init() {

	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// enableCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// enableCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
