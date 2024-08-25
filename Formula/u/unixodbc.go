package main

import (
	"fmt"
	"os/exec"
)

func installUnixODBC() {
	url := "https://downloads.sourceforge.net/project/unixodbc/unixODBC/2.3.11/unixODBC-2.3.11.tar.gz"
	fileName := "unixODBC-2.3.11.tar.gz"
	extractDir := "unixODBC-2.3.11"

	fmt.Println("Downloading unixODBC...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting unixODBC...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring unixODBC...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling unixODBC...")
	exec.Command("make").Run()

	fmt.Println("Installing unixODBC...")
	exec.Command("make", "install").Run()
}
