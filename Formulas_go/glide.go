package main

// Glide - Simplified Go project management, dependency management, and vendoring
// Homepage: https://github.com/Masterminds/glide

import (
	"fmt"
	
	"os/exec"
)

func installGlide() {
	// Método 1: Descargar y extraer .tar.gz
	glide_tar_url := "https://github.com/Masterminds/glide/archive/refs/tags/v0.13.3.tar.gz"
	glide_cmd_tar := exec.Command("curl", "-L", glide_tar_url, "-o", "package.tar.gz")
	err := glide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glide_zip_url := "https://github.com/Masterminds/glide/archive/refs/tags/v0.13.3.zip"
	glide_cmd_zip := exec.Command("curl", "-L", glide_zip_url, "-o", "package.zip")
	err = glide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glide_bin_url := "https://github.com/Masterminds/glide/archive/refs/tags/v0.13.3.bin"
	glide_cmd_bin := exec.Command("curl", "-L", glide_bin_url, "-o", "binary.bin")
	err = glide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glide_src_url := "https://github.com/Masterminds/glide/archive/refs/tags/v0.13.3.src.tar.gz"
	glide_cmd_src := exec.Command("curl", "-L", glide_src_url, "-o", "source.tar.gz")
	err = glide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glide_cmd_direct := exec.Command("./binary")
	err = glide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
