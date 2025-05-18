package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fleet",
	Short: "A modern package manager",
	Long:  `Fleet is a package manager built with Go, designed to install, manage, and configure different tools in a simple and efficient way.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	rootCmd.AddCommand(NewSimulateCmd())
}
