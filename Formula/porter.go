package main

// Porter - App artifacts, tools, configs, and logic packaged as distributable installer
// Homepage: https://porter.sh

import (
	"fmt"
	
	"os/exec"
)

func installPorter() {
	// Método 1: Descargar y extraer .tar.gz
	porter_tar_url := "https://github.com/getporter/porter/archive/refs/tags/v1.1.0.tar.gz"
	porter_cmd_tar := exec.Command("curl", "-L", porter_tar_url, "-o", "package.tar.gz")
	err := porter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	porter_zip_url := "https://github.com/getporter/porter/archive/refs/tags/v1.1.0.zip"
	porter_cmd_zip := exec.Command("curl", "-L", porter_zip_url, "-o", "package.zip")
	err = porter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	porter_bin_url := "https://github.com/getporter/porter/archive/refs/tags/v1.1.0.bin"
	porter_cmd_bin := exec.Command("curl", "-L", porter_bin_url, "-o", "binary.bin")
	err = porter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	porter_src_url := "https://github.com/getporter/porter/archive/refs/tags/v1.1.0.src.tar.gz"
	porter_cmd_src := exec.Command("curl", "-L", porter_src_url, "-o", "source.tar.gz")
	err = porter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	porter_cmd_direct := exec.Command("./binary")
	err = porter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
