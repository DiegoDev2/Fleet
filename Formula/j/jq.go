package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installJq() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installJqForMac()
	case "linux":
		installJqForLinux()
	case "windows":
		installJqForWindows()
	}
}

func installJqForMac() {
	url := "https://github.com/stedolan/jq/releases/download/jq-1.6/jq-osx-amd64"
	downloadAndInstall(url, "jq", "chmod", "+x", "jq", "&&", "sudo", "mv", "jq", "/usr/local/bin/")
}

func installJqForLinux() {
	url := "https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64"
	downloadAndInstall(url, "jq", "chmod", "+x", "jq", "&&", "sudo", "mv", "jq", "/usr/local/bin/")
}

func installJqForWindows() {
	url := "https://github.com/stedolan/jq/releases/download/jq-1.6/jq-win64.exe"
	downloadAndInstall(url, "jq.exe", "mv", "jq.exe", "C:\\Windows\\System32")
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
	installJq()
}
