// Copyright (C) 2024 Fleet Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package formulas

import (
	"os"
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
		if err := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/").Run(); err != nil {
			redBold.Println("Error detaching the disk image:", err)
			return
		}
		return
	}

	yellow.Println("Installing Git from the package...")
	if err := exec.Command("sudo", "installer", "-pkg", pkgPath, "-target", "/").Run(); err != nil {
		redBold.Println("Error installing Git:", err)
		if err := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/").Run(); err != nil {
			redBold.Println("Error detaching the disk image:", err)
			return
		}
		return
	}

	zipFile := "git.zip"
	download = exec.Command("curl", "-L", "https://github.com/git-for-windows/git/releases/download/v2.38.1.windows.1/Git-2.38.1-64-bit.zip", "-o", zipFile)
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Git:", err)
		if err := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/").Run(); err != nil {
			redBold.Println("Error detaching the disk image:", err)
			return
		}
		return
	}
	defer os.Remove(zipFile)

	yellow.Println("Extracting Git...")
	unzip := exec.Command("unzip", zipFile)
	if err := unzip.Run(); err != nil {
		redBold.Println("Error extracting Git:", err)
		if err := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/").Run(); err != nil {
			redBold.Println("Error detaching the disk image:", err)
			return
		}
		return
	}

	install := exec.Command("sudo", "installer", "-pkg", "/Volumes/Git/Git.app/Contents/Resources/git-installer.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Git:", err)
		if err := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/").Run(); err != nil {
			redBold.Println("Error detaching the disk image:", err)
			return
		}
	}

	yellow.Println("Unmounting the disk image...")
	detach := exec.Command("hdiutil", "detach", "/Volumes/Git 2.2.1 Mavericks Intel Universal/")
	if err := detach.Run(); err != nil {
		redBold.Println("Error unmounting the disk image:", err)
		return
	}

	boldGreen.Println("Git installed successfully 🎉")
}
