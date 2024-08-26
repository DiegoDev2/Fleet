package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installDocker() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installDockerForMac()
	case "linux":
		installDockerForLinux()
	case "windows":
		installDockerForWindows()
	}
}

func installDockerForMac() {
	url := "https://desktop.docker.com/mac/stable/Docker.dmg"
	downloadAndInstall(url, "Docker.dmg", "sudo", "hdiutil", "attach", "Docker.dmg", "&&", "cp", "-R", "/Volumes/Docker/Docker.app", "/Applications", "&&", "sudo", "hdiutil", "detach", "/Volumes/Docker")
}

func installDockerForLinux() {
	url := "https://download.docker.com/linux/static/stable/x86_64/docker-20.10.7.tgz"
	downloadAndInstall(url, "docker.tgz", "tar", "-xvzf", "docker.tgz", "&&", "sudo", "mv", "docker/*", "/usr/local/bin/", "&&", "sudo", "groupadd", "docker", "&&", "sudo", "usermod", "-aG", "docker", "$USER", "&&", "newgrp", "docker")
}

func installDockerForWindows() {
	url := "https://desktop.docker.com/win/stable/Docker%20Desktop%20Installer.exe"
	downloadAndInstall(url, "DockerDesktopInstaller.exe", "DockerDesktopInstaller.exe", "/install", "/quiet", "/norestart")
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
	installDocker()
}
