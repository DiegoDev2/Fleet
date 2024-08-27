package main

import (
	"fmt"
	"os"
	"os/exec"
)

var zlib = &Formula{
	Name:        "zlib",
	Description: "Compression library",
	Homepage:    "https://www.zlib.net/",
	URL:         "https://zlib.net/zlib-1.2.12.tar.gz",
	Sha256:      "d02c3c74f3a09a4dc6a6d3fc86e4579c68c3c07ad69f1b74131d1a396580e7d0b",
	License:     "Zlib",
	Install: func() error {
		fmt.Println("Downloading zlib...")
		cmd := exec.Command("curl", "-LO", "https://zlib.net/zlib-1.2.12.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Extracting zlib...")
		cmd = exec.Command("tar", "-xzf", "zlib-1.2.12.tar.gz")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Configuring and installing zlib...")
		cmd = exec.Command("cd", "zlib-1.2.12", "&&", "./configure")
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
		fmt.Println("Testing zlib installation...")
		cmd := exec.Command("zlib", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := zlib.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := zlib.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("zlib installed and tested successfully!")
}
