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

func InstallAircrackNg() {
	switch runtime.GOOS {
	case "darwin":
		installAircrackNgMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installAircrackNgMac() {
	boldGreen.Println("Starting Aircrack-ng installation 🚀")

	yellow.Println("Cloning aircrack-ng repository...")
	download := exec.Command("git", "clone", "https://github.com/aircrack-ng/aircrack-ng")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading K9s:", err)
		return
	}

	yellow.Println("Entering aircrack-ng directory...")
	if err := os.Chdir("aircrack-ng"); err != nil {
		redBold.Println("Error changing to K9s directory:", err)
		return
	}

	yellow.Println("Building aircrack-ng...")
	build := exec.Command("autoreconf", "-i", "&&", "./configure", "--with-experimental", "&&", "make", "&&", "make", "install", "&&", "ldconfig")
	if err := build.Run(); err != nil {
		redBold.Println("Error building aircrack-ng:", err)
		return
	}

	yellow.Println("Moving aircrack-ng binary to /usr/local/bin...")
	if err := os.Rename("aircrack-ng", "/usr/local/bin/aircrack-ng"); err != nil {
		redBold.Println("Error moving aircrack-ng binary:", err)
		return
	}

	boldGreen.Println("aircrack-ng installed successfully 🎉")
}
