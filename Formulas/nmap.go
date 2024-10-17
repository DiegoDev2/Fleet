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

func InstallNmap() {
	switch runtime.GOOS {
	case "darwin":
		installNmapMac()
	default:
		fmt.Println("Este script solo está preparado para macOS (darwin).")
	}
}

func installNmapMac() {
	url := "https://nmap.org/dist/nmap-7.95.tar.bz2"

	fmt.Println("Descargando Nmap...")
	download := exec.Command("curl", "-L", url, "-o", "nmap.tar.bz2")
	download.Stdout = os.Stdout
	download.Stderr = os.Stderr
	if err := download.Run(); err != nil {
		fmt.Println("Error descargando Nmap:", err)
		return
	}

	fmt.Println("Descomprimiendo Nmap...")
	extract := exec.Command("tar", "-xjf", "nmap.tar.bz2")
	extract.Stdout = os.Stdout
	extract.Stderr = os.Stderr
	if err := extract.Run(); err != nil {
		fmt.Println("Error descomprimiendo Nmap:", err)
		return
	}

	fmt.Println("Entrando en el directorio de Nmap...")
	if err := os.Chdir("nmap-7.95"); err != nil {
		fmt.Println("Error cambiando de directorio:", err)
		return
	}

	fmt.Println("Compilando Nmap...")
	configure := exec.Command("./configure")
	configure.Stdout = os.Stdout
	configure.Stderr = os.Stderr
	if err := configure.Run(); err != nil {
		fmt.Println("Error configurando Nmap:", err)
		return
	}

	make := exec.Command("make")
	make.Stdout = os.Stdout
	make.Stderr = os.Stderr
	if err := make.Run(); err != nil {
		fmt.Println("Error compilando Nmap:", err)
		return
	}

	fmt.Println("Instalando Nmap...")
	install := exec.Command("sudo", "make", "install")
	install.Stdout = os.Stdout
	install.Stderr = os.Stderr
	if err := install.Run(); err != nil {
		fmt.Println("Error instalando Nmap:", err)
		return
	}

	fmt.Println("Nmap instalado correctamente.")
}
