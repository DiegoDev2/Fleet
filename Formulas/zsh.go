// Copyright (C) 2024 Fleet Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// InstallZsh installs zsh on Linux and macOS systems

package formulas

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// InstallZsh installs zsh on Linux and macOS systems
func InstallZsh() error {
	var installCmd *exec.Cmd

	// Determine the operating system
	switch runtime.GOOS {
	case "linux":
		// Identify the package manager in Linux
		if isCommandAvailable("apt") {
			installCmd = exec.Command("sudo", "apt", "install", "-y", "zsh")
		} else if isCommandAvailable("yum") {
			installCmd = exec.Command("sudo", "yum", "install", "-y", "zsh")
		} else if isCommandAvailable("pacman") {
			installCmd = exec.Command("sudo", "pacman", "-S", "--noconfirm", "zsh")
		} else {
			return fmt.Errorf("no compatible package manager found on Linux")
		}
	case "darwin":
		// Use Homebrew on macOS
		if isCommandAvailable("brew") {
			installCmd = exec.Command("brew", "install", "zsh")
		} else {
			return fmt.Errorf("Homebrew is not installed on macOS")
		}
	default:
		return fmt.Errorf("this script only supports Linux and macOS")
	}

	// Execute the installation command
	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr

	if err := installCmd.Run(); err != nil {
		return fmt.Errorf("error installing zsh: %v", err)
	}

	fmt.Println("Zsh installed successfully")
	return nil
}

// UninstallZsh removes zsh from Linux and macOS systems
func UninstallZsh() error {
	var uninstallCmd *exec.Cmd

	// Determine the operating system
	switch runtime.GOOS {
	case "linux":
		// Identify the package manager in Linux
		if isCommandAvailable("apt") {
			uninstallCmd = exec.Command("sudo", "apt", "remove", "-y", "zsh")
		} else if isCommandAvailable("yum") {
			uninstallCmd = exec.Command("sudo", "yum", "remove", "-y", "zsh")
		} else if isCommandAvailable("pacman") {
			uninstallCmd = exec.Command("sudo", "pacman", "-R", "--noconfirm", "zsh")
		} else {
			return fmt.Errorf("no compatible package manager found on Linux")
		}
	case "darwin":
		// Use Homebrew on macOS
		if isCommandAvailable("brew") {
			uninstallCmd = exec.Command("brew", "uninstall", "zsh")
		} else {
			return fmt.Errorf("Homebrew is not installed on macOS")
		}
	default:
		return fmt.Errorf("this script only supports Linux and macOS")
	}

	// Execute the uninstallation command
	uninstallCmd.Stdout = os.Stdout
	uninstallCmd.Stderr = os.Stderr

	if err := uninstallCmd.Run(); err != nil {
		return fmt.Errorf("error uninstalling zsh: %v", err)
	}

	fmt.Println("Zsh uninstalled successfully")
	return nil
}

// UpdateZsh updates zsh to the latest version on Linux and macOS systems
func UpdateZsh() error {
	var updateCmd *exec.Cmd

	// Determine the operating system
	switch runtime.GOOS {
	case "linux":
		// Identify the package manager in Linux
		if isCommandAvailable("apt") {
			updateCmd = exec.Command("sudo", "apt", "update", "-y")
			updateCmd.Run() // Update package list
			updateCmd = exec.Command("sudo", "apt", "upgrade", "-y", "zsh")
		} else if isCommandAvailable("yum") {
			updateCmd = exec.Command("sudo", "yum", "update", "-y", "zsh")
		} else if isCommandAvailable("pacman") {
			updateCmd = exec.Command("sudo", "pacman", "-Syu", "--noconfirm", "zsh")
		} else {
			return fmt.Errorf("no compatible package manager found on Linux")
		}
	case "darwin":
		// Use Homebrew on macOS
		if isCommandAvailable("brew") {
			updateCmd = exec.Command("brew", "update")
			updateCmd.Run() // Update Homebrew package list
			updateCmd = exec.Command("brew", "upgrade", "zsh")
		} else {
			return fmt.Errorf("Homebrew is not installed on macOS")
		}
	default:
		return fmt.Errorf("this script only supports Linux and macOS")
	}

	// Execute the update command
	updateCmd.Stdout = os.Stdout
	updateCmd.Stderr = os.Stderr

	if err := updateCmd.Run(); err != nil {
		return fmt.Errorf("error updating zsh: %v", err)
	}

	fmt.Println("Zsh updated successfully")
	return nil
}

// isCommandAvailable checks if a command is available on the system
func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}
