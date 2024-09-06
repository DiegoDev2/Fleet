package main

// Reshape - Easy-to-use, zero-downtime schema migration tool for Postgres
// Homepage: https://github.com/fabianlindfors/reshape

import (
	"fmt"
	
	"os/exec"
)

func installReshape() {
	// Método 1: Descargar y extraer .tar.gz
	reshape_tar_url := "https://github.com/fabianlindfors/reshape/archive/refs/tags/v0.7.0.tar.gz"
	reshape_cmd_tar := exec.Command("curl", "-L", reshape_tar_url, "-o", "package.tar.gz")
	err := reshape_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reshape_zip_url := "https://github.com/fabianlindfors/reshape/archive/refs/tags/v0.7.0.zip"
	reshape_cmd_zip := exec.Command("curl", "-L", reshape_zip_url, "-o", "package.zip")
	err = reshape_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reshape_bin_url := "https://github.com/fabianlindfors/reshape/archive/refs/tags/v0.7.0.bin"
	reshape_cmd_bin := exec.Command("curl", "-L", reshape_bin_url, "-o", "binary.bin")
	err = reshape_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reshape_src_url := "https://github.com/fabianlindfors/reshape/archive/refs/tags/v0.7.0.src.tar.gz"
	reshape_cmd_src := exec.Command("curl", "-L", reshape_src_url, "-o", "source.tar.gz")
	err = reshape_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reshape_cmd_direct := exec.Command("./binary")
	err = reshape_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
