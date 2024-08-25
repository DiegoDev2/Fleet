package main

import (
	"fmt"
	"os/exec"
)

func installCurl() {
	url := "https://curl.se/download/curl-7.79.1.tar.xz"
	fileName := "curl-7.79.1.tar.xz"
	extractDir := "curl-7.79.1"

	fmt.Println("Downloading curl...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting curl...")
	exec.Command("tar", "-xf", fileName).Run()

	fmt.Println("Configuring curl...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling curl...")
	exec.Command("make").Run()

	fmt.Println("Installing curl...")
	exec.Command("make", "install").Run()
}
