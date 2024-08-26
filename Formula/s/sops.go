package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installSops() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installSopsForMac()
	case "linux":
		installSopsForLinux()
	case "windows":
		installSopsForWindows()
	}
}

func installSopsForMac() {
	url := "https://github.com/mozilla/sops/releases/download/v3.7.1/sops-v3.7.1.darwin"
	downloadAndInstall(url, "sops", "chmod", "+x", "sops", "&&", "sudo", "mv", "sops", "/usr/local/bin/")
}

func installSopsForLinux() {
	url := "https://github.com/mozilla/sops/releases/download/v3.7.1/sops-v3.7.1.linux"
	downloadAndInstall(url, "sops", "chmod", "+x", "sops", "&&", "sudo", "mv", "sops", "/usr/local/bin/")
}

func installSopsForWindows() {
	url := "https://github.com/mozilla/sops/releases/download/v3.7.1/sops-v3.7.1.exe"
	downloadAndInstall(url, "sops.exe", "mv", "sops.exe", "C:\\Windows\\System32")
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
	installSops()
}
