package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installGit() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installGitForMac()
	case "linux":
		installGitForLinux()
	case "windows":
		installGitForWindows()
	}
}

func installGitForMac() {
	url := "https://sourceforge.net/projects/git-osx-installer/files/git-2.33.0-intel-universal-mavericks.dmg"
	downloadAndInstall(url, "git.dmg", "sudo", "hdiutil", "attach", "git.dmg", "&&", "sudo", "installer", "-pkg", "/Volumes/Git 2.33.0 Intel Universal/git-2.33.0-intel-universal-mavericks.pkg", "-target", "/")
}

func installGitForLinux() {
	runCommand(exec.Command("sudo", "apt-get", "install", "-y", "git"))
}

func installGitForWindows() {
	url := "https://github.com/git-for-windows/git/releases/download/v2.33.0.windows.1/Git-2.33.0-64-bit.exe"
	downloadAndInstall(url, "git.exe", "start", "/wait", "git.exe", "/SILENT")
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
	installGit()
}
