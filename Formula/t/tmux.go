package main

import (
	"fmt"
	"os/exec"
)

func installTmux() {
	url := "https://github.com/tmux/tmux/releases/download/3.3/tmux-3.3.tar.gz"
	fileName := "tmux-3.3.tar.gz"
	extractDir := "tmux-3.3"

	fmt.Println("Downloading tmux...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting tmux...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring tmux...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling tmux...")
	exec.Command("make").Run()

	fmt.Println("Installing tmux...")
	exec.Command("make", "install").Run()
}
