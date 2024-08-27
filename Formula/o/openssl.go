package main

import (
	"fmt"
	"os"
	"os/exec"
)

var openssl = &Formula{
	Name:        "openssl",
	Description: "Toolkit for Secure Sockets Layer (SSL) and Transport Layer Security (TLS)",
	Homepage:    "https://www.openssl.org/",
	URL:         "https://www.openssl.org/source/openssl-3.0.6.tar.gz",
	Sha256:      "d1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
	License:     "OpenSSL",
	Install: func() error {
		fmt.Println("Downloading openssl...")
		cmd := exec.Command("curl", "-LO", "https://www.openssl.org/source/openssl-3.0.6.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting openssl...")
		cmd = exec.Command("tar", "-xzf", "openssl-3.0.6.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing openssl...")
		cmd = exec.Command("cd", "openssl-3.0.6", "&&", "./configure")
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
		fmt.Println("Testing openssl installation...")
		cmd := exec.Command("openssl", "version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := openssl.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := openssl.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("openssl installed and tested successfully!")
}
