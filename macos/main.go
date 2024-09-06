package main

import (
	"fmt"
	"os/exec"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var messages = map[string]string{
	"installing":         "Installing",
	"uninstalling":       "Uninstalling",
	"upgrading":          "Upgrading",
	"patience":           "This operation is taking longer than expected. Please be patient...",
	"error":              "Error",
	"success":            "Installation successful! You can now use 'latte' from anywhere.",
	"failed":             "Installation failed.",
	"enterPackage":       "Enter the package name:",
	"checkUpdate":        "Checking for updates...",
	"updateAvailable":    "Update available! Downloading the latest version...",
	"updateNotAvailable": "No updates available. You're using the latest version.",
}

func main() {
	Command()
}

func Command() {
	var rootCmd = &cobra.Command{
		Use:   "Latte",
		Short: "A CLI tool for managing packages",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install [package]",
			Short: "Install a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
		&cobra.Command{
			Use:   "uninstall [package]",
			Short: "Uninstall a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
		&cobra.Command{
			Use:   "upgrade [package]",
			Short: "Upgrade a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
		&cobra.Command{
			Use:   "version",
			Short: "Show the version of LattePkg",
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
		&cobra.Command{
			Use:   "upload [package]",
			Short: "upload your package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {

			},
		},
	)

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Usage: latte [command] [flags]")
		color.Yellow.Println("Commands:")
		color.Green.Println("  install   Install a package")
		color.Green.Println("  uninstall Uninstall a package")
		color.Green.Println("  upgrade   Upgrade a package")
		color.Green.Println("  version   Show the version of LattePkg")
		color.Red.Println("Flags:")
		color.Magenta.Println("  --help    Show this help message")
	})

	rootCmd.Execute()
}

func executeCommand(cmd *exec.Cmd) (string, error) {
	cmd.Stdout = nil
	cmd.Stderr = nil
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing command: %s", err)
	}
	return string(output), nil
}
