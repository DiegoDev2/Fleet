package main

import (
	"fmt"
	"os/exec"
)

func installYarn() {
	url := "https://github.com/yarnpkg/yarn/releases/download/v1.22.19/yarn-v1.22.19.tar.gz"
	fileName := "yarn-v1.22.19.tar.gz"
	extractDir := "yarn-v1.22.19"

	fmt.Println("Downloading yarn...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting yarn...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Yarn is ready to use.")
}
