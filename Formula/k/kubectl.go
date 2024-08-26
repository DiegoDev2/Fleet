package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installKubectl() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installKubectlForMac()
	case "linux":
		installKubectlForLinux()
	case "windows":
		installKubectlForWindows()
	}
}

func installKubectlForMac() {
	url := "https://dl.k8s.io/release/v1.21.0/bin/darwin/amd64/kubectl"
	downloadAndInstall(url, "kubectl", "chmod", "+x", "kubectl", "&&", "sudo", "mv", "kubectl", "/usr/local/bin/")
}

func installKubectlForLinux() {
	url := "https://dl.k8s.io/release/v1.21.0/bin/linux/amd64/kubectl"
	downloadAndInstall(url, "kubectl", "chmod", "+x", "kubectl", "&&", "sudo", "mv", "kubectl", "/usr/local/bin/")
}

func installKubectlForWindows() {
	url := "https://dl.k8s.io/release/v1.21.0/bin/windows/amd64/kubectl.exe"
	downloadAndInstall(url, "kubectl.exe", "mv", "kubectl.exe", "C:\\Windows\\System32")
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
	installKubectl()
}
