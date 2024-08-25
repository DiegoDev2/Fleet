package main

import (
	"fmt"
	"os/exec"
)

func installNmap() {
	url := "https://nmap.org/dist/nmap-7.95.tar.bz2"
	fileName := "nmap-7.95.tar.bz2"
	extractDir := "nmap-7.95"

	fmt.Println("Downloading nmap...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting nmap...")
	exec.Command("tar", "-xjf", fileName).Run()

	fmt.Println("Configuring nmap...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling nmap...")
	exec.Command("make").Run()

	fmt.Println("Installing nmap...")
	exec.Command("make", "install").Run()
}
