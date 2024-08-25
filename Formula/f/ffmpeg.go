package main

import (
	"fmt"
	"os/exec"
)

func installFfmpeg() {
	url := "https://ffmpeg.org/releases/ffmpeg-5.0.tar.gz"
	fileName := "ffmpeg-5.0.tar.gz"
	extractDir := "ffmpeg-5.0"

	fmt.Println("Downloading ffmpeg...")
	exec.Command("curl", "-LO", url).Run()
	fmt.Println(extractDir)

	fmt.Println("Extracting ffmpeg...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Configuring ffmpeg...")
	exec.Command("sh", "configure", "--prefix=/usr/local").Run()

	fmt.Println("Compiling ffmpeg...")
	exec.Command("make").Run()

	fmt.Println("Installing ffmpeg...")
	exec.Command("make", "install").Run()
}
