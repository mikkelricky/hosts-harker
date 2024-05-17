/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var (
	diff          = false
	noInteraction = false
	hostsPath     = "/etc/hosts"
	domains       = []string{}

	mainScript = func() string {
		//      If the app is run with `go run`, the executable will (probably) be in the temporary folder.
		if executable, err := os.Executable(); err == nil && !strings.HasPrefix(executable, os.TempDir()) {
			// We think that we're not being run with `go run`.
			return os.Args[0]
		}

		return "go run github.com/mikkelricky/hosts-harker"
	}()

	rootCmd = &cobra.Command{
		Use:   "hosts-harker",
		Short: "Manage domains in your hosts file",
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&hostsPath, "hosts-path", "", hostsPath, "path to hosts file")
	rootCmd.PersistentFlags().BoolVarP(&noInteraction, "no-interaction", "", noInteraction, "No interaction")
	rootCmd.PersistentFlags().BoolVarP(&diff, "diff", "", diff, "Show changes")
}
