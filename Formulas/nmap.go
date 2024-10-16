package Formulas

import (
	"os"
	"os/exec"
	"runtime"
)

func installNmap() {
	switch runtime.GOOS {
	//case "windows":
	//	installNmapWin()
	case "darwin":
		installNmapMac()
		//case "linux":
		//	installNmapLinux()
	}
}
func installNmapMac() {
	url := "https://nmap.org/dist/nmap-7.95.dmg"
	install := exec.Command("curl", "-L", url, "-o", "nmap.dmg")
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	install.Run()
	exec.Command("chmod", "777", "nmap.dmg").Run()
	exec.Command("open", "nmap.dmg").Run()
}
