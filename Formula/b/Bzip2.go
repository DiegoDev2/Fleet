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

var bzip2 = &Formula{
	Name:        "bzip2",
	Description: "A high-quality data compressor",
	Homepage:    "https://www.sourceware.org/bzip2/",
	URL:         "https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz",
	Sha256:      "ab5a03176ee106d3f0fa90e381da478ddae405918153cca248e682cd0c4a2269",
	License:     "BSD-4-Clause",
	Install: func() error {
		fmt.Println("Downloading bzip2...")
		cmd := exec.Command("curl", "-LO", "https://sourceware.org/pub/bzip2/bzip2-1.0.8.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting bzip2...")
		cmd = exec.Command("tar", "-xzf", "bzip2-1.0.8.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building bzip2...")
		cmd = exec.Command("make")
		cmd.Dir = "bzip2-1.0.8"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing bzip2...")
		cmd = exec.Command("make", "install")
		cmd.Dir = "bzip2-1.0.8"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing bzip2...")
		cmd := exec.Command("bzip2", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := bzip2.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := bzip2.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("bzip2 installed and tested successfully!")
}
