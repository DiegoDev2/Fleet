package main

import (
	"fmt"
	"os/exec"
)

func installDiffutils() {
	url := "https://ftp.gnu.org/gnu/diffutils/diffutils-3.8.tar.xz"
	fileName := "diffutils-3.8.tar.xz"
	extractDir := "diffutils-3.8"

	fmt.Println("Downloading diffutils...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting diffutils...")
	exec.Command("tar", "-xf", fileName).Run()

	fmt.Println("Configuring diffutils...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling diffutils...")
	exec.Command("make").Run()

	fmt.Println("Installing diffutils...")
	exec.Command("make", "install").Run()
}
