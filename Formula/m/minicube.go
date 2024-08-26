package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installMinikube() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installMinikubeForMac()
	case "linux":
		installMinikubeForLinux()
	case "windows":
		installMinikubeForWindows()
	}
}

func installMinikubeForMac() {
	url := "https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64"
	downloadAndInstall(url, "minikube", "chmod", "+x", "minikube", "&&", "sudo", "mv", "minikube", "/usr/local/bin/")
}

func installMinikubeForLinux() {
	url := "https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64"
	downloadAndInstall(url, "minikube", "chmod", "+x", "minikube", "&&", "sudo", "mv", "minikube", "/usr/local/bin/")
}

func installMinikubeForWindows() {
	url := "https://storage.googleapis.com/minikube/releases/latest/minikube-windows-amd64.exe"
	downloadAndInstall(url, "minikube.exe", "mv", "minikube.exe", "C:\\Windows\\System32")
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
	installMinikube()
}
