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

func InstallBazel() {
	switch runtime.GOOS {
	case "darwin":
		installBazelMac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installBazelMac() {
	boldGreen.Println("Starting Bazel installation 🚀")
	yellow.Println("Downloading Bazel...")

	download := exec.Command("curl", "-fLO", "https://github.com/bazelbuild/bazel/releases/download/5.0.0/bazel-5.0.0-installer-darwin-x86_64.sh")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Bazel:", err)
		return
	}

	yellow.Println("Setting execute permissions for Bazel installer...")
	chmod := exec.Command("chmod", "+x", "bazel-5.0.0-installer-darwin-x86_64.sh")
	if err := chmod.Run(); err != nil {
		redBold.Println("Error setting permissions for Bazel installer:", err)
		return
	}

	yellow.Println("Installing Bazel...")
	install := exec.Command("bash", "bazel-5.0.0-installer-darwin-x86_64.sh")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Bazel:", err)
		return
	}
	yellow.Println("Deleting Bazel installer...")
	delete := exec.Command("rm", "bazel-5.0.0-installer-darwin-x86_64.sh")
	if err := delete.Run(); err != nil {
		redBold.Println("Error deleting Bazel installer:", err)
		return
	}

	boldGreen.Println("Bazel installed successfully 🎉")
}
