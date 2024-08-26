package main

import (
	"fmt"
	"os"
	"os/exec"
)

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

var ubuntu = &Formula{
	Name:        "Ubuntu",
	Description: "Ubuntu Linux Distribution",
	Homepage:    "https://ubuntu.com/",
	URL:         "https://releases.ubuntu.com/22.04.2/ubuntu-22.04.2-live-server-amd64.iso",
	Sha256:      "a89c18b5b788b8b6229e142b8343dd54bd05b8f7e5cb0b8b6a8e3c15f31c1c07",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading Ubuntu ISO...")
		cmd := exec.Command("curl", "-LO", "https://releases.ubuntu.com/22.04.2/ubuntu-22.04.2-live-server-amd64.iso")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Ubuntu ISO downloaded.")
		return nil
	},
	Test: func() error {
		fmt.Println("Testing Ubuntu ISO...")
		cmd := exec.Command("file", "ubuntu-22.04.2-live-server-amd64.iso")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := ubuntu.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := ubuntu.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Ubuntu ISO downloaded and tested successfully!")
}
