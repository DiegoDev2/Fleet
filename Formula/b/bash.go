package main

import (
	"fmt"
	"os/exec"
)

func installBash() {
	url := "https://ftp.gnu.org/gnu/bash/bash-5.1.tar.gz"
	fileName := "bash-5.1.tar.gz"
	extractDir := "bash-5.1"

	fmt.Println("Downloading bash...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting bash...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring bash...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling bash...")
	exec.Command("make").Run()

	fmt.Println("Installing bash...")
	exec.Command("make", "install").Run()
}
