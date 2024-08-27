package main

import (
	"fmt"
	"os"
	"os/exec"
)

var coreutils = &Formula{
	Name:        "coreutils",
	Description: "GNU core utilities",
	Homepage:    "https://www.gnu.org/software/coreutils/",
	URL:         "https://ftp.gnu.org/gnu/coreutils/coreutils-9.3.tar.xz",
	Sha256:      "45e9e3d236f5865a90e6a94b56a274d7328b13c4d9c09bfeab1d28d8c2dc96d0",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading coreutils...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/coreutils/coreutils-9.3.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting coreutils...")
		cmd = exec.Command("tar", "-xJf", "coreutils-9.3.tar.xz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing coreutils...")
		cmd = exec.Command("cd", "coreutils-9.3", "&&", "./configure")
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
		fmt.Println("Testing coreutils installation...")
		cmd := exec.Command("ls", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := coreutils.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := coreutils.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("coreutils installed and tested successfully!")
}
