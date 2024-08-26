package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Definición de la estructura Formula
type Formula struct {
	Name        string
	Description string
	Homepage    string
	URL         string
	Sha256      string
	License     string
	Install     func() error
	Test        func() error
}

func (f *Formula) InstallPackage() error {
	fmt.Printf("Installing %s...\n", f.Name)
	if err := f.Install(); err != nil {
		return fmt.Errorf("installation failed: %v", err)
	}
	return nil
}

func (f *Formula) TestPackage() error {
	fmt.Printf("Testing %s...\n", f.Name)
	if err := f.Test(); err != nil {
		return fmt.Errorf("testing failed: %v", err)
	}
	return nil
}

var youtubeDL = &Formula{
	Name:        "YouTube-DL",
	Description: "YouTube-DL Download Tool",
	Homepage:    "https://github.com/ytdl-org/youtube-dl",
	URL:         "https://github.com/ytdl-org/youtube-dl/releases/download/2024.08.25/youtube-dl",
	Sha256:      "a3cfb5e748bcfbf5c7563c6e72f35baf5566b2f6211c4de0a4288ecf9c5e6a16",
	License:     "Unlicense",
	Install: func() error {
		fmt.Println("Downloading YouTube-DL...")
		cmd := exec.Command("curl", "-LO", "https://github.com/ytdl-org/youtube-dl/releases/download/2024.08.25/youtube-dl")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Making YouTube-DL executable...")
		cmd = exec.Command("chmod", "+x", "youtube-dl")
		if err := cmd.Run(); err != nil {
			return err
		}

		fmt.Println("Moving YouTube-DL to /usr/local/bin...")
		cmd = exec.Command("mv", "youtube-dl", "/usr/local/bin/youtube-dl")
		if err := cmd.Run(); err != nil {
			return err
		}

		return nil
	},
	Test: func() error {
		fmt.Println("Testing YouTube-DL...")
		cmd := exec.Command("youtube-dl", "--version")
		output, err := cmd.CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(output))
		return nil
	},
}

func main() {
	if err := youtubeDL.InstallPackage(); err != nil {
		fmt.Println("Installation failed:", err)
		os.Exit(1)
	}

	if err := youtubeDL.TestPackage(); err != nil {
		fmt.Println("Testing failed:", err)
		os.Exit(1)
	}

	fmt.Println("YouTube-DL installed and tested successfully!")
}
