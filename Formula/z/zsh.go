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

var zsh = &Formula{
	Name:        "Zsh",
	Description: "Zsh Shell",
	Homepage:    "https://www.zsh.org/",
	URL:         "https://sourceforge.net/projects/zsh/files/zsh/5.9/zsh-5.9.tar.xz",
	Sha256:      "3e42ec69ef11b4bb64d1d5e41a6a91dd9d9bca9e6b2d36c8d70253f44d9bda07",
	License:     "BSD-3-Clause",
	Install: func() error {
		fmt.Println("Downloading Zsh...")
		cmd := exec.Command("curl", "-LO", "https://sourceforge.net/projects/zsh/files/zsh/5.9/zsh-5.9.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Zsh...")
		cmd = exec.Command("tar", "-xJf", "zsh-5.9.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Zsh...")
		cmd = exec.Command("cd", "zsh-5.9", "&&", "./configure", "&&", "make", "&&", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Zsh...")
		cmd := exec.Command("zsh", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := zsh.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := zsh.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Zsh installed and tested successfully!")
}
