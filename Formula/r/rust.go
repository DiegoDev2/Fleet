package main

import (
	"fmt"
	"os/exec"
)

func installRust() {
	url := "https://static.rust-lang.org/dist/rust-1.70.0-src.tar.gz"
	fileName := "rust-1.70.0-src.tar.gz"
	extractDir := "rust-1.70.0"

	fmt.Println("Downloading Rust...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting Rust...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Building Rust...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling Rust...")
	exec.Command("make").Run()

	fmt.Println("Installing Rust...")
	exec.Command("make", "install").Run()
}
