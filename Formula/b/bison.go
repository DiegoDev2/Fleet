package main

import (
	"fmt"
	"os"
	"os/exec"
)

var bison = &Formula{
	Name:        "bison",
	Description: "Parser generator, part of the GNU Project",
	Homepage:    "https://www.gnu.org/software/bison/",
	URL:         "https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.gz",
	Sha256:      "14b68d8572ae99a26dd7e15ed1d33a3e4219a933f7f21b6de7c0cb8e0f27a41d",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading Bison...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/bison/bison-3.8.2.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting Bison...")
		cmd = exec.Command("tar", "-xzf", "bison-3.8.2.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing Bison...")
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
		fmt.Println("Testing Bison installation...")
		cmd := exec.Command("bison", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := bison.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := bison.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("Bison installed and tested successfully!")
}
