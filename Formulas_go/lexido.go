package main

// Lexido - Innovative assistant for the command-line
// Homepage: https://github.com/micr0-dev/lexido

import (
	"fmt"
	
	"os/exec"
)

func installLexido() {
	// Método 1: Descargar y extraer .tar.gz
	lexido_tar_url := "https://github.com/micr0-dev/lexido/archive/refs/tags/v1.4.3.tar.gz"
	lexido_cmd_tar := exec.Command("curl", "-L", lexido_tar_url, "-o", "package.tar.gz")
	err := lexido_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lexido_zip_url := "https://github.com/micr0-dev/lexido/archive/refs/tags/v1.4.3.zip"
	lexido_cmd_zip := exec.Command("curl", "-L", lexido_zip_url, "-o", "package.zip")
	err = lexido_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lexido_bin_url := "https://github.com/micr0-dev/lexido/archive/refs/tags/v1.4.3.bin"
	lexido_cmd_bin := exec.Command("curl", "-L", lexido_bin_url, "-o", "binary.bin")
	err = lexido_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lexido_src_url := "https://github.com/micr0-dev/lexido/archive/refs/tags/v1.4.3.src.tar.gz"
	lexido_cmd_src := exec.Command("curl", "-L", lexido_src_url, "-o", "source.tar.gz")
	err = lexido_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lexido_cmd_direct := exec.Command("./binary")
	err = lexido_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
