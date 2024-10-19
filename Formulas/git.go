package formulas

import (
	"os/exec"
	"runtime"
)

func InstallGit() {
	switch runtime.GOOS {
	case "darwin":
		installGitMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installGitMac() {
	url := "https://github.com/timcharper/git_osx_installer/releases/download/2.2.1/git-2.2.1-intel-universal-mavericks.dmg"

	boldGreen.Println("Starting Git installation 🚀")
	yellow.Println("Downloading Git...")

	download := exec.Command("curl", "-L", url, "-o", "git.dmg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Git:", err)
		return
	}

	yellow.Println("Mounting the disk image...")
	mount := exec.Command("hdiutil", "attach", "git.dmg")
	if err := mount.Run(); err != nil {
		redBold.Println("Error mounting the disk image:", err)
		return
	}

	pkgPath := "/Volumes/Git 2.2.1 Mavericks Intel Universal/git-2.2.1-intel-universal-mavericks.pkg"
	if err := exec.Command("test", "-f", pkgPath).Run(); err != nil {
		redBold.Println("Error finding the .pkg file:", err)
		exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Intel Universal Mavericks").Run() // Desmontar en caso de error
		return
	}

	yellow.Println("Installing Git from the package...")
	install := exec.Command("sudo", "installer", "-pkg", pkgPath, "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Git:", err)
		exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Intel Universal Mavericks").Run() // Desmontar en caso de error
		return
	}

	yellow.Println("Unmounting the disk image...")
	detach := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/")
	if err := detach.Run(); err != nil {
		redBold.Println("Error unmounting the disk image:", err)
		return
	}

	boldGreen.Println("Git installed successfully 🎉")
}
