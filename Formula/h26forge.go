package main

// H26forge - Tool for making syntactically valid but semantically spec-noncompliant videos
// Homepage: https://github.com/h26forge/h26forge

import (
	"fmt"
	
	"os/exec"
)

func installH26forge() {
	// Método 1: Descargar y extraer .tar.gz
	h26forge_tar_url := "https://github.com/h26forge/h26forge/archive/refs/tags/2024-07-06.tar.gz"
	h26forge_cmd_tar := exec.Command("curl", "-L", h26forge_tar_url, "-o", "package.tar.gz")
	err := h26forge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	h26forge_zip_url := "https://github.com/h26forge/h26forge/archive/refs/tags/2024-07-06.zip"
	h26forge_cmd_zip := exec.Command("curl", "-L", h26forge_zip_url, "-o", "package.zip")
	err = h26forge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	h26forge_bin_url := "https://github.com/h26forge/h26forge/archive/refs/tags/2024-07-06.bin"
	h26forge_cmd_bin := exec.Command("curl", "-L", h26forge_bin_url, "-o", "binary.bin")
	err = h26forge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	h26forge_src_url := "https://github.com/h26forge/h26forge/archive/refs/tags/2024-07-06.src.tar.gz"
	h26forge_cmd_src := exec.Command("curl", "-L", h26forge_src_url, "-o", "source.tar.gz")
	err = h26forge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	h26forge_cmd_direct := exec.Command("./binary")
	err = h26forge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
