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

// InstallAws instala AWS CLI en macOS
func InstallAws() {
	switch runtime.GOOS {
	case "darwin":
		installAwsMac()
	default:
		fmt.Println("Este script solo está preparado para macOS (darwin).")
	}
}

func installAwsMac() {
	url := "https://awscli.amazonaws.com/AWSCLIV2.pkg"

	fmt.Println("Descargando AWS CLI...")
	download := exec.Command("curl", "-L", url, "-o", "AWSCLIV2.pkg")
	download.Stdout = os.Stdout
	download.Stderr = os.Stderr
	if err := download.Run(); err != nil {
		fmt.Println("Error descargando AWS CLI:", err)
		return
	}

	fmt.Println("Instalando AWS CLI...")
	install := exec.Command("sudo", "installer", "-pkg", "AWSCLIV2.pkg", "-target", "/")
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	if err := install.Run(); err != nil {
		fmt.Println("Error instalando AWS CLI:", err)
		return
	}

	fmt.Println("AWS CLI instalado correctamente.")

	// Opcional: Limpiar el archivo descargado
	if err := os.Remove("AWSCLIV2.pkg"); err != nil {
		fmt.Println("Error eliminando el archivo .pkg:", err)
		return
	}

	fmt.Println("Limpieza completa.")
}
