package main

import (
	"fmt"
	"os/exec"
)

func installOpenSSL() {
	url := "https://www.openssl.org/source/openssl-3.0.7.tar.gz"
	fileName := "openssl-3.0.7.tar.gz"
	extractDir := "openssl-3.0.7"

	fmt.Println("Downloading OpenSSL...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting OpenSSL...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring OpenSSL...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling OpenSSL...")
	exec.Command("make").Run()

	fmt.Println("Installing OpenSSL...")
	exec.Command("make", "install").Run()
}
