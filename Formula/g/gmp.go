package main

import (
	"fmt"
	"os"
	"os/exec"
)

var gmp = &Formula{
	Name:        "gmp",
	Description: "GNU Multiple Precision Arithmetic Library",
	Homepage:    "https://gmplib.org/",
	URL:         "https://gmplib.org/download/gmp/gmp-6.2.1.tar.xz",
	Sha256:      "59d12f9b76815fc4d9cf0128c3b6a4d02c85f1cf2a8b0736e0806a8a22f1658f",
	License:     "LGPL-3.1",
	Install: func() error {
		fmt.Println("Downloading gmp...")
		cmd := exec.Command("curl", "-LO", "https://gmplib.org/download/gmp/gmp-6.2.1.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting gmp...")
		cmd = exec.Command("tar", "-xJf", "gmp-6.2.1.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing gmp...")
		cmd = exec.Command("cd", "gmp-6.2.1", "&&", "./configure")
		if err := cmd.Run(); err != nil {
			return err
		}

		cmd = exec.Command("make")
		if err := cmd.Run(); err != nil {
			return err
		}

		cmd = exec.Command("sudo", "make", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing gmp installation...")
		cmd := exec.Command("gmp", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := gmp.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := gmp.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("gmp installed and tested successfully!")
}
