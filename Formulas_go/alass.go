package main

// Alass - Automatic Language-Agnostic Subtitle Synchronization
// Homepage: https://github.com/kaegi/alass

import (
	"fmt"
	
	"os/exec"
)

func installAlass() {
	// Método 1: Descargar y extraer .tar.gz
	alass_tar_url := "https://github.com/kaegi/alass/archive/refs/tags/v2.0.0.tar.gz"
	alass_cmd_tar := exec.Command("curl", "-L", alass_tar_url, "-o", "package.tar.gz")
	err := alass_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alass_zip_url := "https://github.com/kaegi/alass/archive/refs/tags/v2.0.0.zip"
	alass_cmd_zip := exec.Command("curl", "-L", alass_zip_url, "-o", "package.zip")
	err = alass_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alass_bin_url := "https://github.com/kaegi/alass/archive/refs/tags/v2.0.0.bin"
	alass_cmd_bin := exec.Command("curl", "-L", alass_bin_url, "-o", "binary.bin")
	err = alass_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alass_src_url := "https://github.com/kaegi/alass/archive/refs/tags/v2.0.0.src.tar.gz"
	alass_cmd_src := exec.Command("curl", "-L", alass_src_url, "-o", "source.tar.gz")
	err = alass_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alass_cmd_direct := exec.Command("./binary")
	err = alass_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
