package main

// ActionValidator - Tool to validate GitHub Action and Workflow YAML files
// Homepage: https://github.com/mpalmer/action-validator

import (
	"fmt"
	
	"os/exec"
)

func installActionValidator() {
	// Método 1: Descargar y extraer .tar.gz
	actionvalidator_tar_url := "https://github.com/mpalmer/action-validator/archive/refs/tags/v0.6.0.tar.gz"
	actionvalidator_cmd_tar := exec.Command("curl", "-L", actionvalidator_tar_url, "-o", "package.tar.gz")
	err := actionvalidator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	actionvalidator_zip_url := "https://github.com/mpalmer/action-validator/archive/refs/tags/v0.6.0.zip"
	actionvalidator_cmd_zip := exec.Command("curl", "-L", actionvalidator_zip_url, "-o", "package.zip")
	err = actionvalidator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	actionvalidator_bin_url := "https://github.com/mpalmer/action-validator/archive/refs/tags/v0.6.0.bin"
	actionvalidator_cmd_bin := exec.Command("curl", "-L", actionvalidator_bin_url, "-o", "binary.bin")
	err = actionvalidator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	actionvalidator_src_url := "https://github.com/mpalmer/action-validator/archive/refs/tags/v0.6.0.src.tar.gz"
	actionvalidator_cmd_src := exec.Command("curl", "-L", actionvalidator_src_url, "-o", "source.tar.gz")
	err = actionvalidator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	actionvalidator_cmd_direct := exec.Command("./binary")
	err = actionvalidator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
