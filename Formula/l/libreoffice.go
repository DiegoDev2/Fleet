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

var libreOffice = &Formula{
	Name:        "LibreOffice",
	Description: "LibreOffice Suite",
	Homepage:    "https://www.libreoffice.org/",
	URL:         "https://download.documentfoundation.org/libreoffice/stable/7.6.0/rpm/x86_64/LibreOffice_7.6.0.3_Linux_x86-64_rpm.tar.gz",
	Sha256:      "6c1e63d46c348d7df5b5b3a7d6c2f95389f8f5d1f4c5e87e1b5684bfe5a3398a",
	License:     "MPL-2.0",
	Install: func() error {
		fmt.Println("Downloading LibreOffice...")
		cmd := exec.Command("curl", "-LO", "https://download.documentfoundation.org/libreoffice/stable/7.6.0/rpm/x86_64/LibreOffice_7.6.0.3_Linux_x86-64_rpm.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting LibreOffice...")
		cmd = exec.Command("tar", "-xzf", "LibreOffice_7.6.0.3_Linux_x86-64_rpm.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing LibreOffice...")
		cmd = exec.Command("rpm", "-i", "LibreOffice_7.6.0.3_Linux_x86-64_rpm/*.rpm")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing LibreOffice...")
		cmd := exec.Command("libreoffice", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := libreOffice.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := libreOffice.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("LibreOffice installed and tested successfully!")
}
