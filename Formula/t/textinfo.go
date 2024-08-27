package main

import (
	"fmt"
	"os"
	"os/exec"
)

var texinfo = &Formula{
	Name:        "texinfo",
	Description: "Documentation system for software packages",
	Homepage:    "https://www.gnu.org/software/texinfo/",
	URL:         "https://ftp.gnu.org/gnu/texinfo/texinfo-6.8.tar.gz",
	Sha256:      "f9a9ef72c36b0d1adf89f5278c062e3889a2059dc09b63ff0bbd6fbb58203d7d",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading texinfo...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/texinfo/texinfo-6.8.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting texinfo...")
		cmd = exec.Command("tar", "-xzf", "texinfo-6.8.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing texinfo...")
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
		fmt.Println("Testing texinfo installation...")
		cmd := exec.Command("info", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := texinfo.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := texinfo.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("texinfo installed and tested successfully!")
}
