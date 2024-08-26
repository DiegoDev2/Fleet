package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installCurl() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installCurlForMac()
	case "linux":
		installCurlForLinux()
	case "windows":
		installCurlForWindows()
	}
}

func installCurlForMac() {
	url := "https://curl.se/download/curl-7.80.0.tar.gz"
	downloadAndInstall(url, "curl.tar.gz", "tar", "-xvzf", "curl.tar.gz", "&&", "cd", "curl-7.80.0", "&&", "./configure", "&&", "make", "&&", "sudo", "make", "install")
}

func installCurlForLinux() {
	url := "https://curl.se/download/curl-7.80.0.tar.gz"
	downloadAndInstall(url, "curl.tar.gz", "tar", "-xvzf", "curl.tar.gz", "&&", "cd", "curl-7.80.0", "&&", "./configure", "&&", "make", "&&", "sudo", "make", "install")
}

func installCurlForWindows() {
	url := "https://curl.se/windows/dl-7.80.0_1/curl-7.80.0-win64-mingw.zip"
	downloadAndInstall(url, "curl.zip", "unzip", "curl.zip", "-d", "C:\\curl", "&&", "set", "PATH=%PATH%;C:\\curl")
}

func downloadAndInstall(url, filename string, installCommand ...string) {
	downloadCmd := exec.Command("curl", "-o", filename, url)
	runCommand(downloadCmd)
	runCommand(exec.Command(installCommand[0], installCommand[1:]...))
}

func runCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	installCurl()
}
