/*
Copyright © 2023 Simon Stone <simon.stone@dartmouth.edu>
*/
package cmd

import (
	"altima/pkg/config"

	"github.com/spf13/cobra"
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

		key := args[0]
		val := config.DeduceType(args[1])

		err := config.UpdateConfig(key, val)
		if err != nil {
			panic(err)
		}
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
