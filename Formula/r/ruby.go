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

var ruby = &Formula{
	Name:        "Ruby",
	Description: "Ruby Programming Language",
	Homepage:    "https://www.ruby-lang.org/",
	URL:         "https://cache.ruby-lang.org/pub/ruby/3.2/ruby-3.2.2.tar.gz",
	Sha256:      "dcff7e7f5fdb9a5a27e1b6d38b2ff71a5a13cbef80f17f54b7f544b8f7a68f8c",
	License:     "BSD-2-Clause",
	Install: func() error {
		fmt.Println("Downloading Ruby...")
		cmd := exec.Command("curl", "-LO", "https://cache.ruby-lang.org/pub/ruby/3.2/ruby-3.2.2.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Ruby...")
		cmd = exec.Command("tar", "-xzf", "ruby-3.2.2.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Ruby...")
		cmd = exec.Command("cd", "ruby-3.2.2", "&&", "./configure", "&&", "make", "&&", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Ruby...")
		cmd := exec.Command("ruby", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := ruby.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := ruby.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Ruby installed and tested successfully!")
}
