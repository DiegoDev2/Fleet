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

var emacs = &Formula{
	Name:        "emacs",
	Description: "GNU Emacs text editor",
	Homepage:    "https://www.gnu.org/software/emacs/",
	URL:         "https://ftp.gnu.org/gnu/emacs/emacs-28.1.tar.xz",
	Sha256:      "ce6240e6b8d98f41ddc4ad3e74cc856b61e8a3722a3aa9af5e9ae1f0fcf90e56",
	License:     "GPL-3.0-or-later",
	Install: func() error {
		fmt.Println("Downloading emacs...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/emacs/emacs-28.1.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting emacs...")
		cmd = exec.Command("tar", "-xf", "emacs-28.1.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring emacs...")
		cmd = exec.Command("./configure", "--prefix=/usr/local")
		cmd.Dir = "emacs-28.1"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing emacs...")
		cmd = exec.Command("make")
		cmd.Dir = "emacs-28.1"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "emacs-28.1"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing emacs...")
		cmd := exec.Command("emacs", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := emacs.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := emacs.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("emacs installed and tested successfully!")
}
