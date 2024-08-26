package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installBazel() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installBazelForMac()
	case "linux":
		installBazelForLinux()
	case "windows":
		installBazelForWindows()
	}
}

func installBazelForMac() {
	url := "https://github.com/bazelbuild/bazel/releases/download/5.0.0/bazel-5.0.0-darwin-x86_64"
	downloadAndInstall(url, "bazel", "chmod", "+x", "bazel", "&&", "sudo", "mv", "bazel", "/usr/local/bin/")
}

func installBazelForLinux() {
	url := "https://github.com/bazelbuild/bazel/releases/download/5.0.0/bazel-5.0.0-linux-x86_64"
	downloadAndInstall(url, "bazel", "chmod", "+x", "bazel", "&&", "sudo", "mv", "bazel", "/usr/local/bin/")
}

func installBazelForWindows() {
	url := "https://github.com/bazelbuild/bazel/releases/download/5.0.0/bazel-5.0.0-windows-x86_64.exe"
	downloadAndInstall(url, "bazel.exe", "mv", "bazel.exe", "C:\\Windows\\System32")
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
	installBazel()
}
