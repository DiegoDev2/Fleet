package main

import (
	"fmt"
	"os/exec"
)

func installXclip() {
	url := "https://sourceforge.net/projects/xclip/files/xclip/0.13/xclip-0.13.tar.gz"
	fileName := "xclip-0.13.tar.gz"
	extractDir := "xclip-0.13"

	fmt.Println("Downloading xclip...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting xclip...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring xclip...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling xclip...")
	exec.Command("make").Run()

	fmt.Println("Installing xclip...")
	exec.Command("make", "install").Run()
}
