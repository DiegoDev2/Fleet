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

func InstallKubectl() {
	switch runtime.GOARCH {
	case "amd64":
		installKubectlMac("https://dl.k8s.io/release/v1.26.0/bin/darwin/amd64/kubectl")
	case "arm64":
		installKubectlMac("https://dl.k8s.io/release/v1.26.0/bin/darwin/arm64/kubectl")
	default:
		redBold.Println("This script only supports macOS.")
	}
}

func installKubectlMac(url string) {
	boldGreen.Println("Starting kubectl installation 🚀")

	yellow.Println("Downloading kubectl...")

	cmd := exec.Command("curl", "-LO", url)
	if err := cmd.Run(); err != nil {
		redBold.Println("Error downloading kubectl:", err)
		return
	}

	yellow.Println("Making kubectl executable...")

	cmd = exec.Command("chmod", "+x", "kubectl")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error making kubectl executable:", err)
		return
	}

	yellow.Println("Moving kubectl to /usr/local/bin...")

	cmd = exec.Command("sudo", "mv", "kubectl", "/usr/local/bin/kubectl")
	if err := cmd.Run(); err != nil {
		redBold.Println("Error moving kubectl:", err)
		return
	}

	boldGreen.Println("kubectl installed successfully 🎉")
}
