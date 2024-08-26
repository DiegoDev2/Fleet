package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installVagrant() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installVagrantForMac()
	case "linux":
		installVagrantForLinux()
	case "windows":
		installVagrantForWindows()
	}
}

func installVagrantForMac() {
	url := "https://releases.hashicorp.com/vagrant/2.2.18/vagrant_2.2.18_x86_64.dmg"
	downloadAndInstall(url, "vagrant.dmg", "sudo", "hdiutil", "attach", "vagrant.dmg", "&&", "sudo", "installer", "-pkg", "/Volumes/Vagrant/vagrant.pkg", "-target", "/")
}

func installVagrantForLinux() {
	url := "https://releases.hashicorp.com/vagrant/2.2.18/vagrant_2.2.18_linux_amd64.zip"
	downloadAndInstall(url, "vagrant.zip", "unzip", "vagrant.zip", "&&", "sudo", "mv", "vagrant", "/usr/local/bin/")
}

func installVagrantForWindows() {
	url := "https://releases.hashicorp.com/vagrant/2.2.18/vagrant_2.2.18_x86_64.msi"
	downloadAndInstall(url, "vagrant.msi", "start", "/wait", "vagrant.msi", "/quiet", "/norestart")
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
	installVagrant()
}
