package main

import (
	"os/exec"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installCMakeMac()
	case "linux":
		installCMakeLinux()
	case "windows":
		installCMakeWindows()
	}
}

func installCMakeMac() {
	url := "https://github.com/Kitware/CMake/releases/download/v3.21.3/cmake-3.21.3-macos-universal.tar.gz"
	filename := "cmake-3.21.3-macos-universal.tar.gz"
	downloadFile(url, filename)
	extractTarGz(filename, "/usr/local/cmake")
}

func installCMakeLinux() {
	url := "https://github.com/Kitware/CMake/releases/download/v3.21.3/cmake-3.21.3-linux-x86_64.tar.gz"
	filename := "cmake-3.21.3-linux-x86_64.tar.gz"
	downloadFile(url, filename)
	extractTarGz(filename, "/usr/local/cmake")
}

func installCMakeWindows() {
	url := "https://github.com/Kitware/CMake/releases/download/v3.21.3/cmake-3.21.3-windows-x86_64.zip"
	filename := "cmake-3.21.3-windows-x86_64.zip"
	downloadFile(url, filename)
	extractZip(filename, "C:\\CMake")
}

func downloadFile(url, filename string) {
	cmd := exec.Command("curl", "-L", "-o", filename, url)
	cmd.Run()
}

func extractTarGz(filename, destination string) {
	cmd := exec.Command("tar", "-xzf", filename, "-C", destination)
	cmd.Run()
}

func extractZip(filename, destination string) {
	cmd := exec.Command("unzip", filename, "-d", destination)
	cmd.Run()
}
