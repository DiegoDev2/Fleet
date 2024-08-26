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

var vlc = &Formula{
	Name:        "VLC Media Player",
	Description: "VLC Media Player",
	Homepage:    "https://www.videolan.org/vlc/",
	URL:         "https://get.videolan.org/vlc/3.0.18/linux/vlc-3.0.18.tar.xz",
	Sha256:      "f098bdeddb79c5f31c6d62dfc26a7f00c43462f0c5c37986a18d7e244b6c134d",
	License:     "GPL-2.0",
	Install: func() error {
		fmt.Println("Downloading VLC...")
		cmd := exec.Command("curl", "-LO", "https://get.videolan.org/vlc/3.0.18/linux/vlc-3.0.18.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting VLC...")
		cmd = exec.Command("tar", "-xJf", "vlc-3.0.18.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing VLC...")
		cmd = exec.Command("cd", "vlc-3.0.18", "&&", "./configure", "&&", "make", "&&", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing VLC...")
		cmd := exec.Command("vlc", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := vlc.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := vlc.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("VLC Media Player installed and tested successfully!")
}
