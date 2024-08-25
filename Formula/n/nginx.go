package main

import (
	"fmt"
	"os/exec"
)

func installNginx() {
	url := "https://nginx.org/download/nginx-1.24.0.tar.gz"
	fileName := "nginx-1.24.0.tar.gz"
	extractDir := "nginx-1.24.0"

	fmt.Println("Downloading Nginx...")
	exec.Command("curl", "-LO", url).Run()

	fmt.Println("Extracting Nginx...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Changing directory to Nginx...")
	exec.Command("cd", extractDir).Run()

	fmt.Println("Configuring Nginx...")
	exec.Command("./configure", "--prefix=/usr/local/nginx").Run()

	fmt.Println("Compiling Nginx...")
	exec.Command("make").Run()

	fmt.Println("Installing Nginx...")
	exec.Command("make", "install").Run()

	fmt.Println("Nginx is ready to use.")
}

func main() {
	installNginx()
}
