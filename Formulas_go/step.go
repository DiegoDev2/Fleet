package main

// Step - Crypto and x509 Swiss-Army-Knife
// Homepage: https://smallstep.com

import (
	"fmt"
	
	"os/exec"
)

func installStep() {
	// Método 1: Descargar y extraer .tar.gz
	step_tar_url := "https://github.com/smallstep/cli/releases/download/v0.27.2/step_0.27.2.tar.gz"
	step_cmd_tar := exec.Command("curl", "-L", step_tar_url, "-o", "package.tar.gz")
	err := step_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	step_zip_url := "https://github.com/smallstep/cli/releases/download/v0.27.2/step_0.27.2.zip"
	step_cmd_zip := exec.Command("curl", "-L", step_zip_url, "-o", "package.zip")
	err = step_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	step_bin_url := "https://github.com/smallstep/cli/releases/download/v0.27.2/step_0.27.2.bin"
	step_cmd_bin := exec.Command("curl", "-L", step_bin_url, "-o", "binary.bin")
	err = step_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	step_src_url := "https://github.com/smallstep/cli/releases/download/v0.27.2/step_0.27.2.src.tar.gz"
	step_cmd_src := exec.Command("curl", "-L", step_src_url, "-o", "source.tar.gz")
	err = step_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	step_cmd_direct := exec.Command("./binary")
	err = step_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
