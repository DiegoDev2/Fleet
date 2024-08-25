package main

import (
	"fmt"
	"os/exec"
)

func installQt() {
	url := "https://download.qt.io/official_releases/qt/5.15/5.15.4/single/qt-everywhere-src-5.15.4.tar.xz"
	fileName := "qt-everywhere-src-5.15.4.tar.xz"
	extractDir := "qt-everywhere-src-5.15.4"

	fmt.Println("Downloading Qt...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting Qt...")
	exec.Command("tar", "-xf", fileName).Run()

	fmt.Println("Configuring Qt...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling Qt...")
	exec.Command("make").Run()

	fmt.Println("Installing Qt...")
	exec.Command("make", "install").Run()
}
