package main

import (
	"fmt"
	"os/exec"
)

func installHtop() {
	url := "https://hisham.hm/htop/releases/3.2.0/htop-3.2.0.tar.gz"
	fileName := "htop-3.2.0.tar.gz"
	extractDir := "htop-3.2.0"

	fmt.Println("Downloading htop...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)
	fmt.Println("Extracting htop...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring htop...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling htop...")
	exec.Command("make").Run()

	fmt.Println("Installing htop...")
	exec.Command("make", "install").Run()
}
