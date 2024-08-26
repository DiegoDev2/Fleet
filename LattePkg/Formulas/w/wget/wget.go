package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installWget() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installWgetForMac()
	case "linux":
		installWgetForLinux()
	case "windows":
		installWgetForWindows()
	}
}

func installWgetForMac() {
	runCommand(exec.Command("brew", "install", "wget"))
}

func installWgetForLinux() {
	runCommand(exec.Command("sudo", "apt-get", "install", "-y", "wget"))
}

func installWgetForWindows() {
	url := "https://eternallybored.org/misc/wget/releases/wget-1.21.1-win64.zip"
	downloadAndInstall(url, "wget.zip", "unzip", "wget.zip", "-d", "C:\\wget", "&&", "set", "PATH=%PATH%;C:\\wget")
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
	installWget()
}
