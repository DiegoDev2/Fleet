package main

import (
	"fmt"
	"os"
	"os/exec"
)

var pylint = &Formula{
	Name:        "pylint",
	Description: "Python code static checker",
	Homepage:    "https://pylint.pycqa.org/",
	URL:         "https://files.pythonhosted.org/packages/source/p/pylint/pylint-2.15.0.tar.gz",
	Sha256:      "e6c9d17b914a6e34e1c5d1e79b15b09f07a87289d5a846b8a8d9b4dbd053ea6f",
	License:     "GPL-2.0-or-later",
	Install: func() error {
		fmt.Println("Downloading pylint...")
		cmd := exec.Command("curl", "-LO", "https://files.pythonhosted.org/packages/source/p/pylint/pylint-2.15.0.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting pylint...")
		cmd = exec.Command("tar", "-xzf", "pylint-2.15.0.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Installing pylint...")
		cmd = exec.Command("cd", "pylint-2.15.0", "&&", "python3", "setup.py", "install")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing pylint installation...")
		cmd := exec.Command("pylint", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := pylint.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := pylint.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("pylint installed and tested successfully!")
}
