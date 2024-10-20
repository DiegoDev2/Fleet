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
	"os/exec"
	"runtime"
)

func InstallDocker() {
	switch runtime.GOARCH {
	case "amd64":
		installDockerMac("https://desktop.docker.com/mac/main/amd64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-amd64")
	case "arm64":
		installDockerMac("https://desktop.docker.com/mac/main/arm64/Docker.dmg?utm_source=docker&utm_medium=webreferral&utm_campaign=docs-driven-download-mac-arm64")
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installDockerMac(url string) {
	boldGreen.Println("Starting Docker installation 🚀")
	yellow.Println("Downloading Docker...")

	download := exec.Command("curl", "-L", url, "-o", "Docker.dmg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Docker:", err)
		return
	}

	yellow.Println("Mounting the disk image...")
	mount := exec.Command("sudo", "hdiutil", "attach", "Docker.dmg")
	if err := mount.Run(); err != nil {
		redBold.Println("Error mounting the disk image:", err)
		return
	}

	yellow.Println("Running the Docker installer...")
	install := exec.Command("sudo", "/Volumes/Docker/Docker.app/Contents/MacOS/install")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Docker:", err)
		exec.Command("hdiutil", "detach", "/Volumes/Docker").Run() // Desmontar en caso de error
		return
	}

	yellow.Println("Unmounting the disk image...")
	detach := exec.Command("sudo", "hdiutil", "detach", "/Volumes/Docker")
	if err := detach.Run(); err != nil {
		redBold.Println("Error unmounting the disk image:", err)
		return
	}

	boldGreen.Println("Docker installed successfully 🎉")
}
