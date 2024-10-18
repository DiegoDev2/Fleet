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
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func InstallNuclei() {
	switch runtime.GOOS {
	case "darwin":
		installNucleiMac()
	default:
		fmt.Println("Este script solo está preparado para macOS (darwin).")
	}
}

func installNucleiMac() {
	cmd := exec.Command("go", "install", "github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Error instalando nuclei:", err)
		return
	}
}
