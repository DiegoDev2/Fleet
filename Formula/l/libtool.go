package main

import (
	"fmt"
	"os"
	"os/exec"
)

var libtool = &Formula{
	Name:        "libtool",
	Description: "GNU Libtool for managing library creation",
	Homepage:    "https://www.gnu.org/software/libtool/",
	URL:         "https://ftp.gnu.org/gnu/libtool/libtool-2.4.7.tar.gz",
	Sha256:      "7e8c5f8f3b9c7c8f5b0e6c2b2d7463b8724d705b5d6c7f9d5b5c704a282745aa",
	License:     "GPL-3.0-or-later",
	Install: func() error {
		fmt.Println("Downloading libtool...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/libtool/libtool-2.4.7.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting libtool...")
		cmd = exec.Command("tar", "-xzf", "libtool-2.4.7.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing libtool...")
		cmd = exec.Command("cd", "libtool-2.4.7", "&&", "./configure")
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
		fmt.Println("Testing libtool installation...")
		cmd := exec.Command("libtool", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := libtool.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := libtool.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("libtool installed and tested successfully!")
}
