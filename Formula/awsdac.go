package main

// Awsdac - CLI tool for drawing AWS architecture
// Homepage: https://github.com/awslabs/diagram-as-code

import (
	"fmt"
	
	"os/exec"
)

func installAwsdac() {
	// Método 1: Descargar y extraer .tar.gz
	awsdac_tar_url := "https://github.com/awslabs/diagram-as-code/archive/refs/tags/v0.21.4.tar.gz"
	awsdac_cmd_tar := exec.Command("curl", "-L", awsdac_tar_url, "-o", "package.tar.gz")
	err := awsdac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsdac_zip_url := "https://github.com/awslabs/diagram-as-code/archive/refs/tags/v0.21.4.zip"
	awsdac_cmd_zip := exec.Command("curl", "-L", awsdac_zip_url, "-o", "package.zip")
	err = awsdac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsdac_bin_url := "https://github.com/awslabs/diagram-as-code/archive/refs/tags/v0.21.4.bin"
	awsdac_cmd_bin := exec.Command("curl", "-L", awsdac_bin_url, "-o", "binary.bin")
	err = awsdac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsdac_src_url := "https://github.com/awslabs/diagram-as-code/archive/refs/tags/v0.21.4.src.tar.gz"
	awsdac_cmd_src := exec.Command("curl", "-L", awsdac_src_url, "-o", "source.tar.gz")
	err = awsdac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsdac_cmd_direct := exec.Command("./binary")
	err = awsdac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
