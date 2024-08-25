package main

import (
	"fmt"
	"os/exec"
)

func installMake() {
	url := "https://ftp.gnu.org/gnu/make/make-4.3.tar.gz"
	fileName := "make-4.3.tar.gz"
	extractDir := "make-4.3"

	fmt.Println("Downloading make...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting make...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring make...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling make...")
	exec.Command("make").Run()

	fmt.Println("Installing make...")
	exec.Command("make", "install").Run()
}
