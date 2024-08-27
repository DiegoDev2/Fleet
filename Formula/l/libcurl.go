package main

import (
	"fmt"
	"os"
	"os/exec"
)

var libcurl = &Formula{
	Name:        "libcurl",
	Description: "Multiprotocol file transfer library",
	Homepage:    "https://curl.se/libcurl/",
	URL:         "https://curl.se/download/curl-7.79.1.tar.gz",
	Sha256:      "d607e74aa0e8a58d2f340f23ff093baea10d12088a9298ddf59a8f7b30eaaedf",
	License:     "curl",
	Install: func() error {
		fmt.Println("Downloading libcurl...")
		cmd := exec.Command("curl", "-LO", "https://curl.se/download/curl-7.79.1.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting libcurl...")
		cmd = exec.Command("tar", "-xzf", "curl-7.79.1.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing libcurl...")
		cmd = exec.Command("./configure", "--disable-shared", "--enable-static")
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
		fmt.Println("Testing libcurl installation...")
		cmd := exec.Command("curl", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := libcurl.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := libcurl.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("libcurl installed and tested successfully!")
}
