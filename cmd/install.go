/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"altima/pkg/repo"
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs modules from the module index",
	Long:  `Installs one or more modules from the module index. You can optionally supply a version number and an alias (to avoid name collision).`,
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

			installName := module.Name
			fmt.Printf("Installing module %q...\n", installName)

			if module.Alias != "" {
				installName = module.Alias
				fmt.Printf("Using install name %q...\n", installName)
			}

			err = repo.InstallModule(installName, url, settings.ModulesDir)

			if err != nil {
				msg := fmt.Sprintf("Failed to install module %q", installName)
				if module.Version != "" {
					msg += fmt.Sprintf(" in version %q", module.Version)
				}
				msg += "!"
				fmt.Println(fmt.Errorf(msg))
				continue
			}
			fmt.Printf("Module %q installed successfully.\n", installName)
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
	Alias   string
}

func getModules(args []string) []Module {
	// Module names, versions and aliases are supplied on the command line like this:
	// moduleA moduleB==v0.0.2 moduleC>myAlias moduleD==v0.0.3>myOtherAlias
	// This function turns the slice of arguments into a slice of Module objects.

	modules := make([]Module, 0)
	for _, arg := range args {
		modules = append(modules, Module{
			getName(arg),
			getVersion(arg),
			getAlias(arg),
		})
	}

	return modules
}

func getName(s string) string {
	// The name of the module comes before the version or the alias
	pattern := regexp.MustCompile("^(.*?)(?:==|>|$)")
	match := pattern.FindStringSubmatch(s)
	if match == nil {
		return ""
	}

	return match[1]
}

func getVersion(s string) string {
	// The version always follows `==` and may precede `>`
	pattern := regexp.MustCompile("==(.*?)(?:>|$)")
	match := pattern.FindStringSubmatch(s)
	if match == nil {
		return ""
	}

	return match[1]
}

func getAlias(s string) string {
	// The alias always follows `>`
	pattern := regexp.MustCompile(">(.*)")
	match := pattern.FindStringSubmatch(s)
	if match == nil {
		return ""
	}

	return match[1]
}
