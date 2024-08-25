package main

import (
	"fmt"
	"os/exec"
)

func installPcre2() {
	url := "https://sourceforge.net/projects/pcre/files/pcre2/10.42/pcre2-10.42.tar.bz2"
	fileName := "pcre2-10.42.tar.bz2"
	extractDir := "pcre2-10.42"

	fmt.Println("Downloading PCRE2...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting PCRE2...")
	exec.Command("tar", "-xjf", fileName).Run()

	fmt.Println("Configuring PCRE2...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling PCRE2...")
	exec.Command("make").Run()

	fmt.Println("Installing PCRE2...")
	exec.Command("make", "install").Run()
}
