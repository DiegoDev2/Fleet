package main

import (
	"fmt"
	"os/exec"
)

func installZsh() {
	url := "https://sourceforge.net/projects/zsh/files/zsh/5.9/zsh-5.9.tar.xz"
	fileName := "zsh-5.9.tar.xz"
	extractDir := "zsh-5.9"

	fmt.Println("Downloading zsh...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting zsh...")
	exec.Command("tar", "-xf", fileName).Run()

	fmt.Println("Configuring zsh...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling zsh...")
	exec.Command("make").Run()

	fmt.Println("Installing zsh...")
	exec.Command("make", "install").Run()
}
