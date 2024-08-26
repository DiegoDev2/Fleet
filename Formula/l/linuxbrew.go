package main

import (
	"os"
	"os/exec"
	"runtime"
)

func installLinuxbrew() {
	switch os := runtime.GOOS; os {
	case "darwin":
		installLinuxbrewForMac()
	case "linux":
		installLinuxbrewForLinux()
	case "windows":
		installLinuxbrewForWindows()
	}
}

func installLinuxbrewForMac() {
	runCommand(exec.Command("/bin/bash", "-c", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""))
}

func installLinuxbrewForLinux() {
	runCommand(exec.Command("/bin/bash", "-c", "\"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""))
}

func installLinuxbrewForWindows() {
	runCommand(exec.Command("curl", "-fsSL", "https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh", "-o", "install.sh", "&&", "bash", "install.sh"))
}

func runCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func main() {
	installLinuxbrew()
}
