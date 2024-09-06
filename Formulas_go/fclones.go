package main

// Fclones - Efficient Duplicate File Finder
// Homepage: https://github.com/pkolaczk/fclones

import (
	"fmt"
	
	"os/exec"
)

func installFclones() {
	// Método 1: Descargar y extraer .tar.gz
	fclones_tar_url := "https://github.com/pkolaczk/fclones/archive/refs/tags/v0.34.0.tar.gz"
	fclones_cmd_tar := exec.Command("curl", "-L", fclones_tar_url, "-o", "package.tar.gz")
	err := fclones_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fclones_zip_url := "https://github.com/pkolaczk/fclones/archive/refs/tags/v0.34.0.zip"
	fclones_cmd_zip := exec.Command("curl", "-L", fclones_zip_url, "-o", "package.zip")
	err = fclones_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fclones_bin_url := "https://github.com/pkolaczk/fclones/archive/refs/tags/v0.34.0.bin"
	fclones_cmd_bin := exec.Command("curl", "-L", fclones_bin_url, "-o", "binary.bin")
	err = fclones_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fclones_src_url := "https://github.com/pkolaczk/fclones/archive/refs/tags/v0.34.0.src.tar.gz"
	fclones_cmd_src := exec.Command("curl", "-L", fclones_src_url, "-o", "source.tar.gz")
	err = fclones_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fclones_cmd_direct := exec.Command("./binary")
	err = fclones_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
