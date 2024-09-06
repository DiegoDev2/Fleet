package main

// Dprint - Pluggable and configurable code formatting platform written in Rust
// Homepage: https://dprint.dev/

import (
	"fmt"
	
	"os/exec"
)

func installDprint() {
	// Método 1: Descargar y extraer .tar.gz
	dprint_tar_url := "https://github.com/dprint/dprint/archive/refs/tags/0.47.2.tar.gz"
	dprint_cmd_tar := exec.Command("curl", "-L", dprint_tar_url, "-o", "package.tar.gz")
	err := dprint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dprint_zip_url := "https://github.com/dprint/dprint/archive/refs/tags/0.47.2.zip"
	dprint_cmd_zip := exec.Command("curl", "-L", dprint_zip_url, "-o", "package.zip")
	err = dprint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dprint_bin_url := "https://github.com/dprint/dprint/archive/refs/tags/0.47.2.bin"
	dprint_cmd_bin := exec.Command("curl", "-L", dprint_bin_url, "-o", "binary.bin")
	err = dprint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dprint_src_url := "https://github.com/dprint/dprint/archive/refs/tags/0.47.2.src.tar.gz"
	dprint_cmd_src := exec.Command("curl", "-L", dprint_src_url, "-o", "source.tar.gz")
	err = dprint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dprint_cmd_direct := exec.Command("./binary")
	err = dprint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
