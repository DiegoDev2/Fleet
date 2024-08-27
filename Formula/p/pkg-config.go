package main

import (
	"fmt"
	"os"
	"os/exec"
)

var pkgConfig = &Formula{
	Name:        "pkg-config",
	Description: "Helper tool for compiling applications and libraries",
	Homepage:    "https://www.freedesktop.org/wiki/Software/pkg-config/",
	URL:         "https://pkg-config.freedesktop.org/releases/pkg-config-0.29.2.tar.gz",
	Sha256:      "9ddcb05d31b9b42e4b1388e8e067e43c37834851f66c7b6d5644d2d03413b4bb",
	License:     "GPL-2.0",
	Install: func() error {
		fmt.Println("Downloading pkg-config...")
		cmd := exec.Command("curl", "-LO", "https://pkg-config.freedesktop.org/releases/pkg-config-0.29.2.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting pkg-config...")
		cmd = exec.Command("tar", "-xzf", "pkg-config-0.29.2.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing pkg-config...")
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
		fmt.Println("Testing pkg-config installation...")
		cmd := exec.Command("pkg-config", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := pkgConfig.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := pkgConfig.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("pkg-config installed and tested successfully!")
}
