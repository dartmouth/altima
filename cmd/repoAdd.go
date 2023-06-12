/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type repoAddOptions struct {
	name string
	url  string
}

// repoAddCmd represents the repoAdd command
var repoAddCmd = &cobra.Command{
	Use:   "add [NAME] [URL]",
	Short: "add a module repository",
	Long:  `add a module repository`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		o := &repoAddOptions{}
		o.name = args[0]
		o.url = args[1]
		o.run()
	},
}

func init() {
	repoCmd.AddCommand(repoAddCmd)
	// settings = cli.EnvSettings

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// repoAddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// repoAddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func (o *repoAddOptions) run() {
	// TODO: Download the module index for the new repository URL to confirm that it is well formed
	viper.SetConfigName(settings.ConfigFilename)
	viper.SetConfigType("toml")
	viper.AddConfigPath(settings.ConfigDir)
	viper.ReadInConfig()
	viper.Set("repositories."+o.name, map[string]string{"url": o.url})
	viper.WriteConfigAs(filepath.Join(settings.ConfigDir, settings.ConfigFilename))
	fmt.Println("Added repository `" + o.name + "` from `" + o.url + "`")

	// // Block deprecated repos
	// if !o.allowDeprecatedRepos {
	// 	for oldURL, newURL := range deprecatedRepos {
	// 		if strings.Contains(o.url, oldURL) {
	// 			return fmt.Errorf("repo %q is no longer available; try %q instead", o.url, newURL)
	// 		}
	// 	}
	// }

	// // Ensure the file directory exists as it is required for file locking
	// err := os.MkdirAll(filepath.Dir(o.repoFile), os.ModePerm)
	// if err != nil && !os.IsExist(err) {
	// 	return err
	// }

	// // Acquire a file lock for process synchronization
	// repoFileExt := filepath.Ext(o.repoFile)
	// var lockPath string
	// if len(repoFileExt) > 0 && len(repoFileExt) < len(o.repoFile) {
	// 	lockPath = strings.TrimSuffix(o.repoFile, repoFileExt) + ".lock"
	// } else {
	// 	lockPath = o.repoFile + ".lock"
	// }
	// fileLock := flock.New(lockPath)
	// lockCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// defer cancel()
	// locked, err := fileLock.TryLockContext(lockCtx, time.Second)
	// if err == nil && locked {
	// 	defer fileLock.Unlock()
	// }
	// if err != nil {
	// 	return err
	// }

	// b, err := os.ReadFile(o.repoFile)
	// if err != nil && !os.IsNotExist(err) {
	// 	return err
	// }

	// var f repo.File
	// if err := yaml.Unmarshal(b, &f); err != nil {
	// 	return err
	// }

	// if o.username != "" && o.password == "" {
	// 	if o.passwordFromStdinOpt {
	// 		passwordFromStdin, err := io.ReadAll(os.Stdin)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		password := strings.TrimSuffix(string(passwordFromStdin), "\n")
	// 		password = strings.TrimSuffix(password, "\r")
	// 		o.password = password
	// 	} else {
	// 		fd := int(os.Stdin.Fd())
	// 		fmt.Fprint(out, "Password: ")
	// 		password, err := term.ReadPassword(fd)
	// 		fmt.Fprintln(out)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		o.password = string(password)
	// 	}
	// }

	// c := repo.Entry{
	// 	Name:                  o.name,
	// 	URL:                   o.url,
	// 	Username:              o.username,
	// 	Password:              o.password,
	// 	PassCredentialsAll:    o.passCredentialsAll,
	// 	CertFile:              o.certFile,
	// 	KeyFile:               o.keyFile,
	// 	CAFile:                o.caFile,
	// 	InsecureSkipTLSverify: o.insecureSkipTLSverify,
	// }

	// // Check if the repo name is legal
	// if strings.Contains(o.name, "/") {
	// 	return errors.Errorf("repository name (%s) contains '/', please specify a different name without '/'", o.name)
	// }

	// // If the repo exists do one of two things:
	// // 1. If the configuration for the name is the same continue without error
	// // 2. When the config is different require --force-update
	// if !o.forceUpdate && f.Has(o.name) {
	// 	existing := f.Get(o.name)
	// 	if c != *existing {

	// 		// The input coming in for the name is different from what is already
	// 		// configured. Return an error.
	// 		return errors.Errorf("repository name (%s) already exists, please specify a different name", o.name)
	// 	}

	// 	// The add is idempotent so do nothing
	// 	fmt.Fprintf(out, "%q already exists with the same configuration, skipping\n", o.name)
	// 	return nil
	// }

	// r, err := repo.NewChartRepository(&c, getter.All(settings))
	// if err != nil {
	// 	return err
	// }

	// if o.repoCache != "" {
	// 	r.CachePath = o.repoCache
	// }
	// if _, err := r.DownloadIndexFile(); err != nil {
	// 	return errors.Wrapf(err, "looks like %q is not a valid chart repository or cannot be reached", o.url)
	// }

	// f.Update(&c)

	// if err := f.WriteFile(o.repoFile, 0600); err != nil {
	// 	return err
	// }
	// fmt.Fprintf(out, "%q has been added to your repositories\n", o.name)
	// return nil
}
