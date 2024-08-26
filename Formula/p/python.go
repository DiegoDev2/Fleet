package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installPython() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installPythonForMac()
	case "linux":
		installPythonForLinux()
	case "windows":
		installPythonForWindows()
	}
}

func installPythonForMac() {
	url := "https://www.python.org/ftp/python/3.9.7/python-3.9.7-macosx10.9.pkg"
	downloadAndInstall(url, "python.pkg", "sudo", "installer", "-pkg", "python.pkg", "-target", "/")
}

func installPythonForLinux() {
	runCommand(exec.Command("sudo", "apt-get", "install", "-y", "python3"))
}

func installPythonForWindows() {
	url := "https://www.python.org/ftp/python/3.9.7/python-3.9.7-amd64.exe"
	downloadAndInstall(url, "python.exe", "start", "/wait", "python.exe", "/quiet", "InstallAllUsers=1", "PrependPath=1")
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
	installPython()
}
