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

var openOffice = &Formula{
	Name:        "Apache OpenOffice",
	Description: "Apache OpenOffice Suite",
	Homepage:    "https://www.openoffice.org/",
	URL:         "https://downloads.apache.org/openoffice/4.1.14/binaries/en-US/Apache_OpenOffice_4.1.14_Linux_x64_install_en-US.tar.gz",
	Sha256:      "5f6e0c8e7d26eb0a813d7485f3d4a1f7c5e3ea703c657d80b8d9e0bfeac053b8",
	License:     "Apache-2.0",
	Install: func() error {
		fmt.Println("Downloading Apache OpenOffice...")
		cmd := exec.Command("curl", "-LO", "https://downloads.apache.org/openoffice/4.1.14/binaries/en-US/Apache_OpenOffice_4.1.14_Linux_x64_install_en-US.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Apache OpenOffice...")
		cmd = exec.Command("tar", "-xzf", "Apache_OpenOffice_4.1.14_Linux_x64_install_en-US.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing Apache OpenOffice...")
		cmd = exec.Command("cd", "Apache_OpenOffice_4.1.14_Linux_x64_install_en-US", "&&", "./install.sh")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing Apache OpenOffice...")
		cmd := exec.Command("openoffice", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := openOffice.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := openOffice.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Apache OpenOffice installed and tested successfully!")
}
