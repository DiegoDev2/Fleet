package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installAWSCLI() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installAWSCLIForMac()
	case "linux":
		installAWSCLIForLinux()
	case "windows":
		installAWSCLIForWindows()
	}
}

func installAWSCLIForMac() {
	url := "https://awscli.amazonaws.com/AWSCLIV2.pkg"
	downloadAndInstall(url, "AWSCLIV2.pkg", "sudo", "installer", "-pkg", "AWSCLIV2.pkg", "-target", "/")
}

func installAWSCLIForLinux() {
	url := "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip"
	downloadAndInstall(url, "awscli-exe-linux-x86_64.zip", "unzip", "awscli-exe-linux-x86_64.zip")
	runCommand(exec.Command("sudo", "./aws/install"))
}

func installAWSCLIForWindows() {
	url := "https://awscli.amazonaws.com/AWSCLIV2.msi"
	downloadAndInstall(url, "AWSCLIV2.msi", "msiexec", "/i", "AWSCLIV2.msi", "/qn")
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
	installAWSCLI()
}
