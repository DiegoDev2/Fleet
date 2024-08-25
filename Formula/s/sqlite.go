package main

import (
	"fmt"
	"os/exec"
)

func installSQLite() {
	url := "https://www.sqlite.org/2024/sqlite-autoconf-3400000.tar.gz"
	fileName := "sqlite-autoconf-3400000.tar.gz"
	extractDir := "sqlite-autoconf-3400000"

	fmt.Println("Downloading SQLite...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting SQLite...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring SQLite...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling SQLite...")
	exec.Command("make").Run()

	fmt.Println("Installing SQLite...")
	exec.Command("make", "install").Run()
}
