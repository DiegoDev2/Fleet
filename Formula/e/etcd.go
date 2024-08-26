package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installEtcd() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installEtcdForMac()
	case "linux":
		installEtcdForLinux()
	case "windows":
		installEtcdForWindows()
	}
}

func installEtcdForMac() {
	url := "https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-darwin-amd64.zip"
	downloadAndInstall(url, "etcd.zip", "unzip", "etcd.zip", "&&", "sudo", "mv", "etcd-v3.5.1-darwin-amd64/etcdctl", "/usr/local/bin/")

}

func installEtcdForLinux() {
	url := "https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-linux-amd64.tar.gz"
	downloadAndInstall(url, "etcd.tar.gz", "tar", "-xvzf", "etcd.tar.gz", "&&", "sudo", "mv", "etcd-v3.5.1-linux-amd64/etcdctl", "/usr/local/bin/")
}

func installEtcdForWindows() {
	url := "https://github.com/etcd-io/etcd/releases/download/v3.5.1/etcd-v3.5.1-windows-amd64.zip"
	downloadAndInstall(url, "etcd.zip", "unzip", "etcd.zip", "-d", "C:\\etcd", "&&", "set", "PATH=%PATH%;C:\\etcd")
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
	installEtcd()
}
