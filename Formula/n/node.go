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

var nodejs = &Formula{
	Name:        "Node.js",
	Description: "Node.js JavaScript Runtime",
	Homepage:    "https://nodejs.org/",
	URL:         "https://nodejs.org/dist/v18.17.1/node-v18.17.1-linux-x64.tar.xz",
	Sha256:      "f8d564d69e4cb1c9a7bbd44d85e7f2601cc6b601e0f7749d6e37994d6a674d2c",
	License:     "MIT",
	Install: func() error {
		fmt.Println("Downloading Node.js...")
		cmd := exec.Command("curl", "-LO", "https://nodejs.org/dist/v18.17.1/node-v18.17.1-linux-x64.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Node.js...")
		cmd = exec.Command("tar", "-xJf", "node-v18.17.1-linux-x64.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Node.js...")
		cmd = exec.Command("mv", "node-v18.17.1-linux-x64", "/usr/local/nodejs")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("ln", "-s", "/usr/local/nodejs/bin/node", "/usr/local/bin/node")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("ln", "-s", "/usr/local/nodejs/bin/npm", "/usr/local/bin/npm")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Node.js...")
		cmd := exec.Command("node", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := nodejs.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := nodejs.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Node.js installed and tested successfully!")
}
