package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installOpenssl() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installOpensslForMac()
	case "linux":
		installOpensslForLinux()
	case "windows":
		installOpensslForWindows()
	}
}

func installOpensslForMac() {
	url := "https://www.openssl.org/source/openssl-1.1.1l.tar.gz"
	downloadAndInstall(url, "openssl.tar.gz", "tar", "-xvzf", "openssl.tar.gz", "&&", "cd", "openssl-1.1.1l", "&&", "./config", "&&", "make", "&&", "sudo", "make", "install")
}

func installOpensslForLinux() {
	url := "https://www.openssl.org/source/openssl-1.1.1l.tar.gz"
	downloadAndInstall(url, "openssl.tar.gz", "tar", "-xvzf", "openssl.tar.gz", "&&", "cd", "openssl-1.1.1l", "&&", "./config", "&&", "make", "&&", "sudo", "make", "install")
}

func installOpensslForWindows() {
	url := "https://slproweb.com/download/Win64OpenSSL-1_1_1l.exe"
	downloadAndInstall(url, "Win64OpenSSL.exe", "start", "/wait", "Win64OpenSSL.exe", "/SILENT")
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
	installOpenssl()
}
