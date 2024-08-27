package main

import (
	"fmt"
	"os"
	"os/exec"
)

var pcre = &Formula{
	Name:        "pcre",
	Description: "Perl Compatible Regular Expressions library",
	Homepage:    "https://pcre.org/",
	URL:         "https://sourceforge.net/projects/pcre/files/pcre/8.45/pcre-8.45.tar.gz",
	Sha256:      "1efc342cf7c0137e2d1a0c248e66062e8a7a7a470a6cb2c9f4d4a1cf74f7b84a",
	License:     "BSD",
	Install: func() error {
		fmt.Println("Downloading pcre...")
		cmd := exec.Command("curl", "-LO", "https://sourceforge.net/projects/pcre/files/pcre/8.45/pcre-8.45.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting pcre...")
		cmd = exec.Command("tar", "-xzf", "pcre-8.45.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing pcre...")
		cmd = exec.Command("cd", "pcre-8.45", "&&", "./configure")
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
		fmt.Println("Testing pcre installation...")
		cmd := exec.Command("pcre-config", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := pcre.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := pcre.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("pcre installed and tested successfully!")
}
