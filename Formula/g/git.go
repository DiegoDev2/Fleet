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

var git = &Formula{
	Name:        "git",
	Description: "Distributed version control system",
	Homepage:    "https://git-scm.com/",
	URL:         "https://www.kernel.org/pub/software/scm/git/git-2.41.0.tar.xz",
	Sha256:      "c0c64a2e1124d4b3a3e9fa159eef8570ed5b058ff4c2be3bb10343b86f0219c1",
	License:     "GPL-2.0-or-later",
	Install: func() error {
		fmt.Println("Downloading git...")
		cmd := exec.Command("curl", "-LO", "https://www.kernel.org/pub/software/scm/git/git-2.41.0.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting git...")
		cmd = exec.Command("tar", "-xf", "git-2.41.0.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring git...")
		cmd = exec.Command("./configure", "--prefix=/usr/local")
		cmd.Dir = "git-2.41.0"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing git...")
		cmd = exec.Command("make")
		cmd.Dir = "git-2.41.0"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "git-2.41.0"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing git...")
		cmd := exec.Command("git", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := git.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := git.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("git installed and tested successfully!")
}
