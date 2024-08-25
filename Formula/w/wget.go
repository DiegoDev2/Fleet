package main

import (
	"fmt"
	"os/exec"
)

func installWget() {
	url := "https://ftp.gnu.org/gnu/wget/wget-1.21.3.tar.gz"
	fileName := "wget-1.21.3.tar.gz"
	extractDir := "wget-1.21.3"

	fmt.Println("Downloading wget...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting wget...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring wget...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling wget...")
	exec.Command("make").Run()

	fmt.Println("Installing wget...")
	exec.Command("make", "install").Run()
}
