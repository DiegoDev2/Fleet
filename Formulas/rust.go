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

func InstallCargo() {
	switch runtime.GOOS {
	case "darwin":
		installCargoMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}
func installCargoMac() {
	boldGreen.Println("Starting Rust installation 🚀")

	yellow.Println("Downloading Rust...")

	cmd := exec.Command("curl", "--proto", "'=https'", "--tlsv1.2", "-sSf", "https://sh.rustup.rs", "-o", "rustup.sh")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error downloading Rust:", err)
		return
	}

	yellow.Println("Installing Rust...")

	cmd = exec.Command("sh", "rustup.sh", "-y")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error installing Cargo:", err)
		return
	}

	boldGreen.Println("Rust installed successfully 🎉")
}
