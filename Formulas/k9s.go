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

	"github.com/fatih/color"
)

func InstallK9s() {
	switch runtime.GOOS {
	case "darwin":
		installK9sMac()
	default:
		fmt.Println("Este script solo está preparado para macOS (darwin).")
	}
}

func installK9sMac() {
	fmt.Println("Clonando el repositorio de k9s...")
	download := exec.Command("git", "clone", "https://github.com/derailed/k9s.git")
	download.Stdout = os.Stdout
	download.Stderr = os.Stderr
	if err := download.Run(); err != nil {
		fmt.Println("Error descargando k9s:", err)
		return
	}

	if err := os.Chdir("k9s"); err != nil {
		fmt.Println("Error cambiando al directorio k9s:", err)
		return
	}

	color.Magenta("Compilando k9s...")
	build := exec.Command("go", "build", "-o", "k9s", "./cmd")
	build.Stdout = os.Stdout
	build.Stderr = os.Stderr
	if err := build.Run(); err != nil {
		fmt.Println("Error compilando k9s:", err)
		return
	}

	fmt.Println("Moviendo k9s a /usr/local/bin...")
	if err := os.Rename("k9s", "/usr/local/bin/k9s"); err != nil {
		fmt.Println("Error moviendo el binario k9s:", err)
		return
	}

}
