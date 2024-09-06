package main

// Copilot - CLI tool for Amazon ECS and AWS Fargate
// Homepage: https://aws.github.io/copilot-cli/

import (
	"fmt"
	
	"os/exec"
)

func installCopilot() {
	// Método 1: Descargar y extraer .tar.gz
	copilot_tar_url := "https://github.com/aws/copilot-cli/archive/refs/tags/v1.34.0.tar.gz"
	copilot_cmd_tar := exec.Command("curl", "-L", copilot_tar_url, "-o", "package.tar.gz")
	err := copilot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	copilot_zip_url := "https://github.com/aws/copilot-cli/archive/refs/tags/v1.34.0.zip"
	copilot_cmd_zip := exec.Command("curl", "-L", copilot_zip_url, "-o", "package.zip")
	err = copilot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	copilot_bin_url := "https://github.com/aws/copilot-cli/archive/refs/tags/v1.34.0.bin"
	copilot_cmd_bin := exec.Command("curl", "-L", copilot_bin_url, "-o", "binary.bin")
	err = copilot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	copilot_src_url := "https://github.com/aws/copilot-cli/archive/refs/tags/v1.34.0.src.tar.gz"
	copilot_cmd_src := exec.Command("curl", "-L", copilot_src_url, "-o", "source.tar.gz")
	err = copilot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	copilot_cmd_direct := exec.Command("./binary")
	err = copilot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
