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

// Definición del paquete a2ps como una fórmula
var a2ps = &Formula{
	Name:        "a2ps",
	Description: "Any-to-PostScript filter",
	Homepage:    "https://www.gnu.org/software/a2ps/",
	URL:         "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz",
	Sha256:      "87ff9d801cb11969181d5b8cf8b65e65e5b24bb0c76a1b825e8098f2906fbdf4",
	License:     "GPL-3.0-or-later",
	Install: func() error {
		fmt.Println("Downloading a2ps...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/a2ps/a2ps-4.15.6.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting a2ps...")
		cmd = exec.Command("tar", "-xzf", "a2ps-4.15.6.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring a2ps...")
		cmd = exec.Command("./configure", "--prefix=/usr/local")
		cmd.Dir = "a2ps-4.15.6"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing a2ps...")
		cmd = exec.Command("make")
		cmd.Dir = "a2ps-4.15.6"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "a2ps-4.15.6"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing a2ps...")
		cmd := exec.Command("a2ps", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := a2ps.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := a2ps.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("a2ps installed and tested successfully!")
}
