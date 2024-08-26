package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installIstioctl() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installIstioctlForMac()
	case "linux":
		installIstioctlForLinux()
	case "windows":
		installIstioctlForWindows()
	}
}

func installIstioctlForMac() {
	url := "https://github.com/istio/istio/releases/download/1.10.0/istioctl-1.10.0-osx.tar.gz"
	downloadAndInstall(url, "istioctl.tar.gz", "tar", "-xvzf", "istioctl.tar.gz", "&&", "sudo", "mv", "istioctl", "/usr/local/bin/")
}

func installIstioctlForLinux() {
	url := "https://github.com/istio/istio/releases/download/1.10.0/istioctl-1.10.0-linux-amd64.tar.gz"
	downloadAndInstall(url, "istioctl.tar.gz", "tar", "-xvzf", "istioctl.tar.gz", "&&", "sudo", "mv", "istioctl", "/usr/local/bin/")
}

func installIstioctlForWindows() {
	url := "https://github.com/istio/istio/releases/download/1.10.0/istioctl-1.10.0-win.zip"
	downloadAndInstall(url, "istioctl.zip", "unzip", "istioctl.zip", "-d", "C:\\istioctl", "&&", "set", "PATH=%PATH%;C:\\istioctl")
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
	installIstioctl()
}
