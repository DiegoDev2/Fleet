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

var qgis = &Formula{
	Name:        "QGIS",
	Description: "QGIS Geographic Information System",
	Homepage:    "https://qgis.org/",
	URL:         "https://qgis.org/downloads/QGIS-3.28.3-x86_64.tar.bz2",
	Sha256:      "7e0dcb7d8de2fcb42c5829cc90ff4d510c4e8c5b6e764e2ffce3ff15c89b32b9",
	License:     "GPL-2.0",
	Install: func() error {
		fmt.Println("Downloading QGIS...")
		cmd := exec.Command("curl", "-LO", "https://qgis.org/downloads/QGIS-3.28.3-x86_64.tar.bz2")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting QGIS...")
		cmd = exec.Command("tar", "-xjf", "QGIS-3.28.3-x86_64.tar.bz2")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing QGIS...")
		cmd = exec.Command("mv", "QGIS-3.28.3", "/opt/qgis")
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("ln", "-s", "/opt/qgis/bin/qgis", "/usr/local/bin/qgis")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing QGIS...")
		cmd := exec.Command("qgis", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := qgis.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := qgis.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("QGIS installed and tested successfully!")
}
