package main

import (
	"fmt"
	"os/exec"
)

func installVim() {
	url := "https://github.com/vim/vim/archive/refs/tags/v9.0.0000.tar.gz"
	fileName := "vim-9.0.0000.tar.gz"
	extractDir := "vim-9.0.0000"

	fmt.Println("Downloading vim...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)
	fmt.Println("Extracting vim...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring vim...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling vim...")
	exec.Command("make").Run()

	fmt.Println("Installing vim...")
	exec.Command("make", "install").Run()
}
