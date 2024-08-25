package main

import (
	"fmt"
	"os/exec"
)

func installJq() {
	url := "https://github.com/stedolan/jq/releases/download/jq-1.6/jq-1.6.tar.gz"
	fileName := "jq-1.6.tar.gz"
	extractDir := "jq-1.6"

	fmt.Println("Downloading jq...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting jq...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring jq...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling jq...")
	exec.Command("make").Run()

	fmt.Println("Installing jq...")
	exec.Command("make", "install").Run()
}
