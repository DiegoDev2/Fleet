package main

import (
	"fmt"
	"os/exec"
)

func installEmacs() {
	url := "https://ftp.gnu.org/gnu/emacs/emacs-27.2.tar.xz"
	fileName := "emacs-27.2.tar.xz"
	extractDir := "emacs-27.2"

	fmt.Println("Downloading emacs...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)
	fmt.Println("Extracting emacs...")
	exec.Command("tar", "-xf", fileName).Run()

	fmt.Println("Configuring emacs...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling emacs...")
	exec.Command("make").Run()

	fmt.Println("Installing emacs...")
	exec.Command("make", "install").Run()
}
