package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var messages = map[string]string{
	"installing":         "Installing",
	"uninstalling":       "Uninstalling",
	"upgrading":          "Upgrading",
	"patience":           "This operation is taking longer than expected. Please be patient...",
	"error":              "Error",
	"success":            "Installation successful! You can now use 'turn' from anywhere.",
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
		Use:   "turn",
		Short: "A CLI tool for managing packages",
	}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "install [package]",
			Short: "Install a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Install(args[0])
			},
		},
		&cobra.Command{
			Use:   "uninstall [package]",
			Short: "Uninstall a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Uninstall(args[0])
			},
		},
		&cobra.Command{
			Use:   "upgrade [package]",
			Short: "Upgrade a package",
			Args:  cobra.ExactArgs(1),
			Run: func(cmd *cobra.Command, args []string) {
				Update(args[0])
			},
		},
	)

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("Usage: turn [command] [flags]")
		color.Yellow.Println("Commands:")
		color.Green.Println("  install   Install a package")
		color.Green.Println("  uninstall Uninstall a package")
		color.Green.Println("  upgrade   Upgrade a package")
		color.Green.Println("  self-update  Update the tool itself")
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

func Install(pkg string) {
	cmd := exec.Command("brew", "install", pkg)
	done := make(chan bool)

	go func() {
		color.Green.Println(messages["installing"] + " " + pkg + "...")
		output, err := executeCommand(cmd)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
		} else {
			if output != "" {
				color.Green.Println(output)
			}
		}
		done <- true
	}()

	select {
	case <-done:
		// Command finished within 30 seconds
	case <-time.After(30 * time.Second):
		color.Yellow.Println(messages["patience"])
	}
}

func Uninstall(pkg string) {
	cmd := exec.Command("brew", "uninstall", pkg)
	done := make(chan bool)

	go func() {
		color.Green.Println(messages["uninstalling"] + " " + pkg + "...")
		output, err := executeCommand(cmd)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
		} else {
			if output != "" {
				color.Green.Println(output)
			}
		}
		done <- true
	}()

	select {
	case <-done:
		// Command finished within 30 seconds
	case <-time.After(30 * time.Second):
		color.Yellow.Println(messages["patience"])
	}
}

func Update(pkg string) {
	cmd := exec.Command("brew", "upgrade", pkg)
	done := make(chan bool)

	go func() {
		color.Green.Println(messages["upgrading"] + " " + pkg + "...")
		output, err := executeCommand(cmd)
		if err != nil {
			color.Red.Println(messages["error"]+":", err)
		} else {
			if output != "" {
				color.Green.Println(output)
			}
		}
		done <- true
	}()

	select {
	case <-done:
		// Command finished within 30 seconds
	case <-time.After(30 * time.Second):
		color.Yellow.Println(messages["patience"])
	}
}
