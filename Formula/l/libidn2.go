package main

import (
	"fmt"
	"os"
	"os/exec"
)

var libidn2 = &Formula{
	Name:        "libidn2",
	Description: "Library to handle internationalized domain names (IDN)",
	Homepage:    "https://www.gnu.org/software/libidn/#libidn2",
	URL:         "https://ftp.gnu.org/gnu/libidn/libidn2-2.3.0.tar.gz",
	Sha256:      "ac788dde67b7125b8b3131ac514bdf6856e967e63d1db5d1f5f8d3a6189ebf21",
	License:     "GPL-3.0",
	Install: func() error {
		fmt.Println("Downloading libidn2...")
		cmd := exec.Command("curl", "-LO", "https://ftp.gnu.org/gnu/libidn/libidn2-2.3.0.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting libidn2...")
		cmd = exec.Command("tar", "-xzf", "libidn2-2.3.0.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing libidn2...")
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
		fmt.Println("Testing libidn2 installation...")
		cmd := exec.Command("idn2", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := libidn2.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := libidn2.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("libidn2 installed and tested successfully!")
}
