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

var flex = &Formula{
	Name:        "flex",
	Description: "Fast Lexical Analyzer",
	Homepage:    "https://github.com/westes/flex",
	URL:         "https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz",
	Sha256:      "e87cacd4da12ad5b7f2dbb4177bcb24ffdcdfb55587b3a0c93097a1f7d5b4dcb",
	License:     "BSD-2-Clause",
	Install: func() error {
		fmt.Println("Downloading flex...")
		cmd := exec.Command("curl", "-LO", "https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting flex...")
		cmd = exec.Command("tar", "-xzf", "flex-2.6.4.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring flex...")
		cmd = exec.Command("./configure", "--prefix=/usr/local")
		cmd.Dir = "flex-2.6.4"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing flex...")
		cmd = exec.Command("make")
		cmd.Dir = "flex-2.6.4"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "flex-2.6.4"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing flex...")
		cmd := exec.Command("flex", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := flex.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := flex.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("flex installed and tested successfully!")
}
