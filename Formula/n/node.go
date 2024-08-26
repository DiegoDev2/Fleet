package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installNode() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installNodeForMac()
	case "linux":
		installNodeForLinux()
	case "windows":
		installNodeForWindows()
	}
}

func installNodeForMac() {
	url := "https://nodejs.org/dist/v14.17.6/node-v14.17.6-darwin-x64.tar.gz"
	downloadAndInstall(url, "node.tar.gz", "tar", "-xvzf", "node.tar.gz", "&&", "sudo", "mv", "node-v14.17.6-darwin-x64/bin/node", "/usr/local/bin/")
}

func installNodeForLinux() {
	url := "https://nodejs.org/dist/v14.17.6/node-v14.17.6-linux-x64.tar.gz"
	downloadAndInstall(url, "node.tar.gz", "tar", "-xvzf", "node.tar.gz", "&&", "sudo", "mv", "node-v14.17.6-linux-x64/bin/node", "/usr/local/bin/")
}

func installNodeForWindows() {
	url := "https://nodejs.org/dist/v14.17.6/node-v14.17.6-win-x64.zip"
	downloadAndInstall(url, "node.zip", "unzip", "node.zip", "-d", "C:\\node", "&&", "set", "PATH=%PATH%;C:\\node")
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
	installNode()
}
