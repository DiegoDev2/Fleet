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

func InstallGo() {
	switch runtime.GOOS {
	case "darwin":
		installGoMac()
	case "arm64":
		installGoArm64()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installGoMac() {
	url := "https://go.dev/dl/go1.22.8.darwin-amd64.pkg"

	boldGreen.Println("Starting Go installation 🚀")
	yellow.Println("Downloading Go...")
	download := exec.Command("curl", "-L", url, "-o", "go.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Go:", err)
		return
	}

	yellow.Println("Installing Go...")
	install := exec.Command("sudo", "installer", "-pkg", "go.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Go:", err)
		return
	}

	boldGreen.Println("Go installed successfully 🎉")
}

func installGoArm64() {
	url := "https://go.dev/dl/go1.22.8.darwin-arm64.pkg"

	boldGreen.Println("Starting Go installation 🚀")
	yellow.Println("Downloading Go...")
	download := exec.Command("curl", "-L", url, "-o", "go.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Go:", err)
		return
	}

	yellow.Println("Installing Go...")
	install := exec.Command("sudo", "installer", "-pkg", "go.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Go:", err)
		return
	}

	boldGreen.Println("Go installed successfully 🎉")
}
