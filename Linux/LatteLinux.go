package main

import (
	"fmt"
	"os/exec"
	"strings"

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
		Use:   "latte",
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
		&cobra.Command{
			Use:   "version",
			Short: "Show the version of LattePkg",
			Run: func(cmd *cobra.Command, args []string) {
				version()
			},
		},
		&cobra.Command{
			Use:   "search [package]",
			Short: "Search for a package",
			Run: func(cmd *cobra.Command, args []string) {
				Search(args[0])
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
		color.Green.Println("  search    Search for packages")
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
	var cmd *exec.Cmd
	if isDebianBased() {
		cmd = exec.Command("sudo", "apt-get", "install", "-y", pkg)
	} else if isRedHatBased() {
		cmd = exec.Command("sudo", "dnf", "install", "-y", pkg)
	} else {
		color.Red.Println("Unsupported Linux distribution.")
		return
	}

	color.Green.Println(messages["installing"] + " " + pkg + "...")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
	} else {
		if output != "" {
			color.Green.Println(output)
		}
	}
}

func Uninstall(pkg string) {
	var cmd *exec.Cmd
	if isDebianBased() {
		cmd = exec.Command("sudo", "apt-get", "remove", "-y", pkg)
	} else if isRedHatBased() {
		cmd = exec.Command("sudo", "dnf", "remove", "-y", pkg)
	} else {
		color.Red.Println("Unsupported Linux distribution.")
		return
	}

	color.Green.Println(messages["uninstalling"] + " " + pkg + "...")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
	} else {
		if output != "" {
			color.Green.Println(output)
		}
	}
}

func Update(pkg string) {
	var cmd *exec.Cmd
	if isDebianBased() {
		cmd = exec.Command("sudo", "apt-get", "update", "-y", pkg)
	} else if isRedHatBased() {
		cmd = exec.Command("sudo", "dnf", "upgrade", "-y", pkg)
	} else {
		color.Red.Println("Unsupported Linux distribution.")
		return
	}

	color.Green.Println(messages["upgrading"] + " " + pkg + "...")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println(messages["error"]+":", err)
	} else {
		if output != "" {
			color.Green.Println(output)
		}
	}
}

func version() {
	com := "LattePkg Version"
	ver := com + " 1.0.2"
	color.Greenln(ver)
}

func Search(search string) {
	var cmd *exec.Cmd
	if isDebianBased() {
		cmd = exec.Command("apt-cache", "search", search)
	} else if isRedHatBased() {
		cmd = exec.Command("dnf", "search", search)
	} else {
		color.Red.Println("Unsupported Linux distribution.")
		return
	}

	color.Green.Println("Searching for", search, "....")
	output, err := cmd.CombinedOutput()
	if err != nil {
		color.Red.Println("Error:", err)
	} else {
		color.Green.Println(string(output))
	}
}

func isDebianBased() bool {
	cmd := exec.Command("lsb_release", "-is")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println("Error detecting distribution:", err)
		return false
	}
	return strings.TrimSpace(string(output)) == "Debian" || strings.TrimSpace(string(output)) == "Ubuntu"
}

func isRedHatBased() bool {
	cmd := exec.Command("lsb_release", "-is")
	output, err := executeCommand(cmd)
	if err != nil {
		color.Red.Println("Error detecting distribution:", err)
		return false
	}
	return strings.TrimSpace(string(output)) == "RedHatEnterpriseServer" || strings.TrimSpace(string(output)) == "CentOS"
}
