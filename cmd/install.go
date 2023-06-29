/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"altima/pkg/repo"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs modules from the module index",
	Long:  `Installs one or more modules from the module index. You can optionally supply a version number`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		modules := getModules(args)

		for _, module := range modules {
			if module.Version != "" {
				fmt.Printf("Searching URL for module %q in version %q...\n", module.Name, module.Version)
			} else {
				fmt.Printf("Searching URL for module %q...\n", module.Name)
			}

			url, err := repo.Search(module.Name, module.Version, settings.RepositoryCacheDir)

			if err != nil {
				msg := fmt.Sprintf("Failed to find module %q", module.Name)
				if module.Version != "" {
					msg += fmt.Sprintf(" in version %q", module.Version)
				}
				msg += " in index!"
				fmt.Println(fmt.Errorf(msg))
				continue
			}

			fmt.Printf("Found URL %q...\n", url)
			fmt.Printf("Installing module %q...\n", module.Name)
			err = repo.InstallModule(module.Name, url, settings.ModulesDir)

			if err != nil {
				msg := fmt.Sprintf("Failed to install module %q", module.Name)
				if module.Version != "" {
					msg += fmt.Sprintf(" in version %q", module.Version)
				}
				msg += "!"
				fmt.Println(fmt.Errorf(msg))
				continue
			}
			fmt.Printf("Module %q installed successfully.\n", module.Name)
		}
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

type Module struct {
	Name    string
	Version string
}

func getModules(args []string) []Module {
	// Module names and versions are supplied on the command line like this:
	// moduleA moduleB==v0.0.2 moduleC
	// This function turns the slice of arguments into a slice of Module objects.

	modules := make([]Module, 0)
	for _, arg := range args {
		if !strings.Contains(arg, "==") {
			modules = append(modules, Module{arg, ""})
		} else {
			parts := strings.Split(arg, "==")
			name := parts[0]
			version := parts[1]
			modules = append(modules, Module{name, version})
		}
	}

	return modules
}
