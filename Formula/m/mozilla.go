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

var firefox = &Formula{
	Name:        "Mozilla Firefox",
	Description: "Mozilla Firefox Web Browser",
	Homepage:    "https://www.mozilla.org/firefox/",
	URL:         "https://ftp.mozilla.org/pub/firefox/releases/117.0.1/linux-x86_64/en-US/firefox-117.0.1.tar.bz2",
	Sha256:      "d6b9b0cb6e63a9e0c61ea18d8e1ed013b6eeb0f30b5a7dc5f7459fc8a94900f7",
	License:     "MPL-2.0",
	Install: func() error {
		fmt.Println("Downloading Mozilla Firefox...")
		cmd := exec.Command("curl", "-LO", "https://ftp.mozilla.org/pub/firefox/releases/117.0.1/linux-x86_64/en-US/firefox-117.0.1.tar.bz2")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Mozilla Firefox...")
		cmd = exec.Command("tar", "-xjf", "firefox-117.0.1.tar.bz2")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Mozilla Firefox...")
		cmd = exec.Command("mv", "firefox", "/opt/firefox")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("ln", "-s", "/opt/firefox/firefox", "/usr/local/bin/firefox")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Mozilla Firefox...")
		cmd := exec.Command("firefox", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := firefox.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := firefox.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Mozilla Firefox installed and tested successfully!")
}
