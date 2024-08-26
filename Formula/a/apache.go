package main

import (
	"os/exec"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installApacheMac()
	case "linux":
		installApacheLinux()
	case "windows":
		installApacheWindows()
	}
}

func installApacheMac() {
	url := "https://downloads.apache.org/httpd/httpd-2.4.57.tar.gz"
	filename := "httpd-2.4.57.tar.gz"
	downloadFile(url, filename)
	extractTarGz(filename, "/usr/local/apache")
}

func installApacheLinux() {
	url := "https://downloads.apache.org/httpd/httpd-2.4.57.tar.gz"
	filename := "httpd-2.4.57.tar.gz"
	downloadFile(url, filename)
	extractTarGz(filename, "/usr/local/apache")
}

func installApacheWindows() {
	url := "https://downloads.apache.org/httpd/binaries/win32/httpd-2.4.57-win32-VC15.zip"
	filename := "httpd-2.4.57-win32.zip"
	downloadFile(url, filename)
	extractZip(filename, "C:\\Apache24")
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
