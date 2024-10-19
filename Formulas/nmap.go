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

func InstallNmap() {
	switch runtime.GOOS {
	case "darwin":
		installNmapMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installNmapMac() {
	url := "https://nmap.org/dist/nmap-7.95.tar.bz2"

	boldGreen.Println("Starting Nmap installation 🚀")
	yellow.Println("Recommended to have XCode installed")
	yellow.Println("Downloading Nmap...")
	download := exec.Command("curl", "-L", url, "-o", "nmap.tar.bz2")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Nmap:", err)
		return
	}

	yellow.Println("Extracting Nmap...")
	extract := exec.Command("tar", "-xjf", "nmap.tar.bz2")
	if err := extract.Run(); err != nil {
		redBold.Println("Error extracting Nmap:", err)
		return
	}

	yellow.Println("Entering Nmap directory...")
	if err := os.Chdir("nmap-7.95"); err != nil {
		redBold.Println("Error changing directory:", err)
		return
	}

	yellow.Println("Compiling Nmap...")
	configure := exec.Command("./configure")
	if err := configure.Run(); err != nil {
		redBold.Println("Error configuring Nmap:", err)
		return
	}

	make := exec.Command("make")
	if err := make.Run(); err != nil {
		redBold.Println("Error compiling Nmap:", err)
		return
	}

	yellow.Println("Installing Nmap...")
	install := exec.Command("sudo", "make", "install")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Nmap:", err)
		return
	}

	boldGreen.Println("Nmap installed successfully 🎉")
}
