package main

// Gpm - Barebones dependency manager for Go
// Homepage: https://github.com/pote/gpm

import (
	"fmt"
	
	"os/exec"
)

func installGpm() {
	// Método 1: Descargar y extraer .tar.gz
	gpm_tar_url := "https://github.com/pote/gpm/archive/refs/tags/v1.4.0.tar.gz"
	gpm_cmd_tar := exec.Command("curl", "-L", gpm_tar_url, "-o", "package.tar.gz")
	err := gpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpm_zip_url := "https://github.com/pote/gpm/archive/refs/tags/v1.4.0.zip"
	gpm_cmd_zip := exec.Command("curl", "-L", gpm_zip_url, "-o", "package.zip")
	err = gpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpm_bin_url := "https://github.com/pote/gpm/archive/refs/tags/v1.4.0.bin"
	gpm_cmd_bin := exec.Command("curl", "-L", gpm_bin_url, "-o", "binary.bin")
	err = gpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpm_src_url := "https://github.com/pote/gpm/archive/refs/tags/v1.4.0.src.tar.gz"
	gpm_cmd_src := exec.Command("curl", "-L", gpm_src_url, "-o", "source.tar.gz")
	err = gpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpm_cmd_direct := exec.Command("./binary")
	err = gpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
