package main

// Statix - Lints and suggestions for the nix programming language
// Homepage: https://github.com/oppiliappan/statix

import (
	"fmt"
	
	"os/exec"
)

func installStatix() {
	// Método 1: Descargar y extraer .tar.gz
	statix_tar_url := "https://github.com/oppiliappan/statix/archive/refs/tags/v0.5.8.tar.gz"
	statix_cmd_tar := exec.Command("curl", "-L", statix_tar_url, "-o", "package.tar.gz")
	err := statix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	statix_zip_url := "https://github.com/oppiliappan/statix/archive/refs/tags/v0.5.8.zip"
	statix_cmd_zip := exec.Command("curl", "-L", statix_zip_url, "-o", "package.zip")
	err = statix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	statix_bin_url := "https://github.com/oppiliappan/statix/archive/refs/tags/v0.5.8.bin"
	statix_cmd_bin := exec.Command("curl", "-L", statix_bin_url, "-o", "binary.bin")
	err = statix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	statix_src_url := "https://github.com/oppiliappan/statix/archive/refs/tags/v0.5.8.src.tar.gz"
	statix_cmd_src := exec.Command("curl", "-L", statix_src_url, "-o", "source.tar.gz")
	err = statix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	statix_cmd_direct := exec.Command("./binary")
	err = statix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
