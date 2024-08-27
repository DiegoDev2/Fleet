package main

import (
	"fmt"
	"os"
	"os/exec"
)

var autoconf = &Formula{
	Name:        "autoconf",
	Description: "Tool for generating configure scripts",
	Homepage:    "https://www.gnu.org/software/autoconf/",
	URL:         "https://ftp.gnu.org/gnu/autoconf/autoconf-2.71.tar.gz",
	Sha256:      "431075ad0bf529ef139a247df6e0f45f9e6b13b4a7b5e4a1e57b383e4d49e533",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading Autoconf...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/autoconf/autoconf-2.71.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Autoconf...")
		cmd = exec.Command("tar", "-xzf", "autoconf-2.71.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing Autoconf...")
		cmd = exec.Command("./configure")
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
		fmt.Println("Testing Autoconf installation...")
		cmd := exec.Command("autoconf", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := autoconf.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := autoconf.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Autoconf installed and tested successfully!")
}
