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

var cmake = &Formula{
	Name:        "cmake",
	Description: "Cross-platform make",
	Homepage:    "https://cmake.org/",
	URL:         "https://github.com/Kitware/CMake/releases/download/v3.27.0-rc1/cmake-3.27.0-rc1.tar.gz",
	Sha256:      "a85d1d8b59f50a3ea57b197ec0df2a0e0c6c6bfa0c17cbb6c243ab5c601dbda4",
	License:     "BSD-3-Clause",
	Install: func() error {
		fmt.Println("Downloading cmake...")
		cmd := exec.Command("curl", "-LO", "https://github.com/Kitware/CMake/releases/download/v3.27.0-rc1/cmake-3.27.0-rc1.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting cmake...")
		cmd = exec.Command("tar", "-xzf", "cmake-3.27.0-rc1.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring cmake...")
		cmd = exec.Command("./bootstrap")
		cmd.Dir = "cmake-3.27.0-rc1"
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Building and installing cmake...")
		cmd = exec.Command("make")
		cmd.Dir = "cmake-3.27.0-rc1"
		if err := cmd.Run(); err != nil {
			return err
		}
		cmd = exec.Command("make", "install")
		cmd.Dir = "cmake-3.27.0-rc1"
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing cmake...")
		cmd := exec.Command("cmake", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := cmake.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := cmake.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("cmake installed and tested successfully!")
}
