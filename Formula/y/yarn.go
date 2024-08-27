package main

import (
	"fmt"
	"os"
	"os/exec"
)

var yarn = &Formula{
	Name:        "yarn",
	Description: "Fast, reliable, and secure dependency management",
	Homepage:    "https://yarnpkg.com/",
	URL:         "https://github.com/yarnpkg/yarn/releases/download/v1.22.19/yarn-v1.22.19.tar.gz",
	Sha256:      "b78a6e80fcfd2f4e5e56e7e1c357eab8bbf82d5f89cf415c2fa78b478e1e1a84",
	License:     "BSD-3-Clause",
	Install: func() error {
		fmt.Println("Downloading yarn...")
		cmd := exec.Command("curl", "-LO", "https://github.com/yarnpkg/yarn/releases/download/v1.22.19/yarn-v1.22.19.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting yarn...")
		cmd = exec.Command("tar", "-xzf", "yarn-v1.22.19.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing yarn...")
		cmd = exec.Command("sudo", "mv", "yarn-v1.22.19/bin/yarn", "/usr/local/bin/yarn")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing yarn installation...")
		cmd := exec.Command("yarn", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := yarn.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := yarn.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("yarn installed and tested successfully!")
}
