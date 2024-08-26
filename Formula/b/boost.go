package main

import (
	"os/exec"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installBoostMac()
	case "linux":
		installBoostLinux()
	case "windows":
		installBoostWindows()
	}
}

func installBoostMac() {
	url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.gz"
	filename := "boost_1_76_0.tar.gz"
	downloadFile(url, filename)
	extractTarGz(filename, "/usr/local/boost")
}

func installBoostLinux() {
	url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.gz"
	filename := "boost_1_76_0.tar.gz"
	downloadFile(url, filename)
	extractTarGz(filename, "/usr/local/boost")
}

func installBoostWindows() {
	url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.zip"
	filename := "boost_1_76_0.zip"
	downloadFile(url, filename)
	extractZip(filename, "C:\\Boost")
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
