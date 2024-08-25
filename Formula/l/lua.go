package main

import (
	"fmt"
	"os/exec"
)

func installLua() {
	url := "https://www.lua.org/ftp/lua-5.4.4.tar.gz"
	fileName := "lua-5.4.4.tar.gz"
	extractDir := "lua-5.4.4"

	fmt.Println("Downloading Lua...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting Lua...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Building Lua...")
	exec.Command("make", "linux").Run()

	fmt.Println("Installing Lua...")
	exec.Command("make", "install").Run()
}
