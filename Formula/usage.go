package main

// Usage - Tool for working with usage-spec CLIs
// Homepage: https://usage.jdx.dev/

import (
	"fmt"
	
	"os/exec"
)

func installUsage() {
	// Método 1: Descargar y extraer .tar.gz
	usage_tar_url := "https://github.com/jdx/usage/archive/refs/tags/v0.3.1.tar.gz"
	usage_cmd_tar := exec.Command("curl", "-L", usage_tar_url, "-o", "package.tar.gz")
	err := usage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	usage_zip_url := "https://github.com/jdx/usage/archive/refs/tags/v0.3.1.zip"
	usage_cmd_zip := exec.Command("curl", "-L", usage_zip_url, "-o", "package.zip")
	err = usage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	usage_bin_url := "https://github.com/jdx/usage/archive/refs/tags/v0.3.1.bin"
	usage_cmd_bin := exec.Command("curl", "-L", usage_bin_url, "-o", "binary.bin")
	err = usage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	usage_src_url := "https://github.com/jdx/usage/archive/refs/tags/v0.3.1.src.tar.gz"
	usage_cmd_src := exec.Command("curl", "-L", usage_src_url, "-o", "source.tar.gz")
	err = usage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	usage_cmd_direct := exec.Command("./binary")
	err = usage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
