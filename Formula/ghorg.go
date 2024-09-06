package main

// Ghorg - Quickly clone an entire org's or user's repositories into one directory
// Homepage: https://github.com/gabrie30/ghorg

import (
	"fmt"
	
	"os/exec"
)

func installGhorg() {
	// Método 1: Descargar y extraer .tar.gz
	ghorg_tar_url := "https://github.com/gabrie30/ghorg/archive/refs/tags/v1.9.13.tar.gz"
	ghorg_cmd_tar := exec.Command("curl", "-L", ghorg_tar_url, "-o", "package.tar.gz")
	err := ghorg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghorg_zip_url := "https://github.com/gabrie30/ghorg/archive/refs/tags/v1.9.13.zip"
	ghorg_cmd_zip := exec.Command("curl", "-L", ghorg_zip_url, "-o", "package.zip")
	err = ghorg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghorg_bin_url := "https://github.com/gabrie30/ghorg/archive/refs/tags/v1.9.13.bin"
	ghorg_cmd_bin := exec.Command("curl", "-L", ghorg_bin_url, "-o", "binary.bin")
	err = ghorg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghorg_src_url := "https://github.com/gabrie30/ghorg/archive/refs/tags/v1.9.13.src.tar.gz"
	ghorg_cmd_src := exec.Command("curl", "-L", ghorg_src_url, "-o", "source.tar.gz")
	err = ghorg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghorg_cmd_direct := exec.Command("./binary")
	err = ghorg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
