package main

import (
	"fmt"
	"os/exec"
)

func installAutoconf() {
	url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.71.tar.gz"
	fileName := "autoconf-2.71.tar.gz"
	extractDir := "autoconf-2.71"

	fmt.Println("Downloading Autoconf...")
	exec.Command("curl", "-LO", url).Run()

	fmt.Println("Extracting Autoconf...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Changing directory to Autoconf...")
	exec.Command("cd", extractDir).Run()

	fmt.Println("Configuring Autoconf...")
	exec.Command("./configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling Autoconf...")
	exec.Command("make").Run()

	fmt.Println("Installing Autoconf...")
	exec.Command("make", "install").Run()

	fmt.Println("Autoconf is ready to use.")
}

func main() {
	installAutoconf()
}
