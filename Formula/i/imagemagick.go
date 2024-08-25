package main

import (
	"fmt"
	"os/exec"
)

func installImagemagick() {
	url := "https://imagemagick.org/download/ImageMagick-7.1.0-25.tar.xz"
	fileName := "ImageMagick-7.1.0-25.tar.xz"
	extractDir := "ImageMagick-7.1.0-25"

	fmt.Println("Downloading ImageMagick...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting ImageMagick...")
	exec.Command("tar", "-xf", fileName).Run()

	fmt.Println("Configuring ImageMagick...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling ImageMagick...")
	exec.Command("make").Run()

	fmt.Println("Installing ImageMagick...")
	exec.Command("make", "install").Run()
}
