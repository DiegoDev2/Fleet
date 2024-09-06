package main

// Protolint - Pluggable linter and fixer to enforce Protocol Buffer style and conventions
// Homepage: https://github.com/yoheimuta/protolint

import (
	"fmt"
	
	"os/exec"
)

func installProtolint() {
	// Método 1: Descargar y extraer .tar.gz
	protolint_tar_url := "https://github.com/yoheimuta/protolint/archive/refs/tags/v0.50.5.tar.gz"
	protolint_cmd_tar := exec.Command("curl", "-L", protolint_tar_url, "-o", "package.tar.gz")
	err := protolint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protolint_zip_url := "https://github.com/yoheimuta/protolint/archive/refs/tags/v0.50.5.zip"
	protolint_cmd_zip := exec.Command("curl", "-L", protolint_zip_url, "-o", "package.zip")
	err = protolint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protolint_bin_url := "https://github.com/yoheimuta/protolint/archive/refs/tags/v0.50.5.bin"
	protolint_cmd_bin := exec.Command("curl", "-L", protolint_bin_url, "-o", "binary.bin")
	err = protolint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protolint_src_url := "https://github.com/yoheimuta/protolint/archive/refs/tags/v0.50.5.src.tar.gz"
	protolint_cmd_src := exec.Command("curl", "-L", protolint_src_url, "-o", "source.tar.gz")
	err = protolint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protolint_cmd_direct := exec.Command("./binary")
	err = protolint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
