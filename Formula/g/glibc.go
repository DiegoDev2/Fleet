package main

import (
	"fmt"
	"os"
	"os/exec"
)

var glibc = &Formula{
	Name:        "glibc",
	Description: "GNU C Library",
	Homepage:    "https://www.gnu.org/software/libc/",
	URL:         "https://ftp.gnu.org/gnu/libc/glibc-2.33.tar.gz",
	Sha256:      "3b6d2f9f9b8c5e3f7d7f9f91b1c4623c07f2632d748706a0a9477a68b4a22d70",
	License:     "LGPL-2.1",
	Install: func() error {
		fmt.Println("Downloading glibc...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/libc/glibc-2.33.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting glibc...")
		cmd = exec.Command("tar", "-xzf", "glibc-2.33.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing glibc...")
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
		fmt.Println("Testing glibc installation...")
		cmd := exec.Command("ldd", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := glibc.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := glibc.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("glibc installed and tested successfully!")
}
