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

var diffutils = &Formula{
	Name:        "diffutils",
	Description: "File comparison tools",
	Homepage:    "https://www.gnu.org/software/diffutils/",
	URL:         "https://ftp.gnu.org/gnu/diffutils/diffutils-3.8.tar.xz",
	Sha256:      "2749b2f2b8ac5b2d073e67adf1719e9e20b73b1d2f65f365b8f96f0bdf80ea21",
	License:     "GPL-3.0-or-later",
	Install: func() error {
		fmt.Println("Downloading diffutils...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/diffutils/diffutils-3.8.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting diffutils...")
		cmd = exec.Command("tar", "-xf", "diffutils-3.8.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring diffutils...")
		cmd = exec.Command("./configure", "--prefix=/usr/local")
		cmd.Dir = "diffutils-3.8"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing diffutils...")
		cmd = exec.Command("make")
		cmd.Dir = "diffutils-3.8"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "diffutils-3.8"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing diffutils...")
		cmd := exec.Command("diff", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := diffutils.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := diffutils.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("diffutils installed and tested successfully!")
}
