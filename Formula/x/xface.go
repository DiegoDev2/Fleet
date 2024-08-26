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

var xfce = &Formula{
	Name:        "Xfce",
	Description: "Xfce Desktop Environment",
	Homepage:    "https://www.xfce.org/",
	URL:         "https://archive.xfce.org/src/xfce/xfce4/4.18/xfce4-4.18.0.tar.bz2",
	Sha256:      "1e1b7dfc6c7d42cfb8ff78a3ae6d7b726c16d5275ea686c8dbf08b36f36b98a4",
	License:     "GPL-2.0",
	Install: func() error {
		fmt.Println("Downloading Xfce...")
		cmd := exec.Command("curl", "-LO", "https://archive.xfce.org/src/xfce/xfce4/4.18/xfce4-4.18.0.tar.bz2")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Xfce...")
		cmd = exec.Command("tar", "-xjf", "xfce4-4.18.0.tar.bz2")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Xfce...")
		cmd = exec.Command("cd", "xfce4-4.18.0", "&&", "./configure", "&&", "make", "&&", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Xfce...")
		cmd := exec.Command("xfce4-session", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := xfce.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := xfce.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Xfce installed and tested successfully!")
}
