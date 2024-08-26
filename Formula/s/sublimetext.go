package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Definición de la estructura Formula
type Formula struct {
	Name        string
	Description string
	Homepage    string
	URL         string
	Sha256      string
	License     string
	Install     func() error
	Test        func() error
}

func (f *Formula) InstallPackage() error {
	fmt.Printf("Installing %s...\n", f.Name)
	if err := f.Install(); err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}
	return nil
}

func (f *Formula) TestPackage() error {
	fmt.Printf("Testing %s...\n", f.Name)
	if err := f.Test(); err != nil {
		return fmt.Errorf("testing failed: %v", err)
	}
	return nil
}

var sublimeText = &Formula{
	Name:        "Sublime Text",
	Description: "Sublime Text Editor",
	Homepage:    "https://www.sublimetext.com/",
	URL:         "https://download.sublimetext.com/sublime_text_4_build_4143_x64.tar.gz",
	Sha256:      "b7b9d95c8677a82d7aee47d4d1469b5fc7c1d0b4b807b91ef8e2e3e2bbf90323",
	License:     "Proprietary",
	Install: func() error {
		fmt.Println("Downloading Sublime Text...")
		cmd := exec.Command("curl", "-LO", "https://download.sublimetext.com/sublime_text_4_build_4143_x64.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Sublime Text...")
		cmd = exec.Command("tar", "-xzf", "sublime_text_4_build_4143_x64.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Sublime Text...")
		cmd = exec.Command("mv", "sublime_text", "/opt/sublime_text")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("ln", "-s", "/opt/sublime_text/sublime_text", "/usr/local/bin/sublime_text")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Sublime Text...")
		cmd := exec.Command("sublime_text", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := sublimeText.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := sublimeText.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Sublime Text installed and tested successfully!")
}
