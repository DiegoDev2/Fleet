package main

// Lab - Git wrapper for GitLab
// Homepage: https://zaquestion.github.io/lab

import (
	"fmt"
	
	"os/exec"
)

func installLab() {
	// Método 1: Descargar y extraer .tar.gz
	lab_tar_url := "https://github.com/zaquestion/lab/archive/refs/tags/v0.25.1.tar.gz"
	lab_cmd_tar := exec.Command("curl", "-L", lab_tar_url, "-o", "package.tar.gz")
	err := lab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lab_zip_url := "https://github.com/zaquestion/lab/archive/refs/tags/v0.25.1.zip"
	lab_cmd_zip := exec.Command("curl", "-L", lab_zip_url, "-o", "package.zip")
	err = lab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lab_bin_url := "https://github.com/zaquestion/lab/archive/refs/tags/v0.25.1.bin"
	lab_cmd_bin := exec.Command("curl", "-L", lab_bin_url, "-o", "binary.bin")
	err = lab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lab_src_url := "https://github.com/zaquestion/lab/archive/refs/tags/v0.25.1.src.tar.gz"
	lab_cmd_src := exec.Command("curl", "-L", lab_src_url, "-o", "source.tar.gz")
	err = lab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lab_cmd_direct := exec.Command("./binary")
	err = lab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
