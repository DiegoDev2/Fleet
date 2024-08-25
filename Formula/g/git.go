package main

import (
	"fmt"
	"os/exec"
)

func installGit() {
	url := "https://www.kernel.org/pub/software/scm/git/git-2.34.1.tar.gz"
	fileName := "git-2.34.1.tar.gz"
	extractDir := "git-2.34.1"

	fmt.Println("Downloading git...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)
	fmt.Println("Extracting git...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring git...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling git...")
	exec.Command("make").Run()

	fmt.Println("Installing git...")
	exec.Command("make", "install").Run()
}
