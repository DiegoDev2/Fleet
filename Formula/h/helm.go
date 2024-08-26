package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installHelm() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installHelmForMac()
	case "linux":
		installHelmForLinux()
	case "windows":
		installHelmForWindows()
	}
}

func installHelmForMac() {
	url := "https://get.helm.sh/helm-v3.6.3-darwin-amd64.tar.gz"
	downloadAndInstall(url, "helm.tar.gz", "tar", "-xvzf", "helm.tar.gz", "&&", "sudo", "mv", "darwin-amd64/helm", "/usr/local/bin/")
}

func installHelmForLinux() {
	url := "https://get.helm.sh/helm-v3.6.3-linux-amd64.tar.gz"
	downloadAndInstall(url, "helm.tar.gz", "tar", "-xvzf", "helm.tar.gz", "&&", "sudo", "mv", "linux-amd64/helm", "/usr/local/bin/")
}

func installHelmForWindows() {
	url := "https://get.helm.sh/helm-v3.6.3-windows-amd64.zip"
	downloadAndInstall(url, "helm.zip", "unzip", "helm.zip", "-d", "C:\\helm", "&&", "set", "PATH=%PATH%;C:\\helm")
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
	installHelm()
}
