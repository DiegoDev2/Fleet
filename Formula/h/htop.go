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

var htop = &Formula{
	Name:        "htop",
	Description: "Interactive process viewer",
	Homepage:    "https://htop.dev/",
	URL:         "https://github.com/htop-dev/htop/releases/download/3.2.2/htop-3.2.2.tar.xz",
	Sha256:      "cba91c4bfa0288b16ebf0264b3578a7f113a20ba4e283ec6f59836e8f61b00f1",
	License:     "GPL-2.0-or-later",
	Install: func() error {
		fmt.Println("Downloading htop...")
		cmd := exec.Command("curl", "-LO", "https://github.com/htop-dev/htop/releases/download/3.2.2/htop-3.2.2.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting htop...")
		cmd = exec.Command("tar", "-xf", "htop-3.2.2.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring htop...")
		cmd = exec.Command("./configure", "--prefix=/usr/local")
		cmd.Dir = "htop-3.2.2"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing htop...")
		cmd = exec.Command("make")
		cmd.Dir = "htop-3.2.2"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "htop-3.2.2"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing htop...")
		cmd := exec.Command("htop", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := htop.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := htop.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("htop installed and tested successfully!")
}
