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

func InstallPython2() {
	switch runtime.GOOS {
	case "darwin":
		installPython2Mac()
	default:
		redBold.Println("This script only supports macOS (darwin).")
	}
}

func installPython2Mac() {
	boldGreen.Println("Starting Python2 installation 🚀")
	yellow.Println("Downloading Python2...")
	download := exec.Command("curl", "-L", "https://www.python.org/ftp/python/2.7.18/python-2.7.18-macosx10.9.pkg", "-o", "python2.pkg")
	if err := download.Run(); err != nil {
		redBold.Println("Error downloading Python2:", err)
		return
	}

	yellow.Println("Installing Python2...")
	install := exec.Command("sudo", "installer", "-pkg", "python2.pkg", "-target", "/")
	if err := install.Run(); err != nil {
		redBold.Println("Error installing Python2:", err)
		return
	}

	boldGreen.Println("Python2 installed successfully 🎉")
}
