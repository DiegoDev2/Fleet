package main

// Xcenv - Xcode version manager
// Homepage: https://github.com/xcenv/xcenv

import (
	"fmt"
	
	"os/exec"
)

func installXcenv() {
	// Método 1: Descargar y extraer .tar.gz
	xcenv_tar_url := "https://github.com/xcenv/xcenv/archive/refs/tags/v1.2.0.tar.gz"
	xcenv_cmd_tar := exec.Command("curl", "-L", xcenv_tar_url, "-o", "package.tar.gz")
	err := xcenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcenv_zip_url := "https://github.com/xcenv/xcenv/archive/refs/tags/v1.2.0.zip"
	xcenv_cmd_zip := exec.Command("curl", "-L", xcenv_zip_url, "-o", "package.zip")
	err = xcenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcenv_bin_url := "https://github.com/xcenv/xcenv/archive/refs/tags/v1.2.0.bin"
	xcenv_cmd_bin := exec.Command("curl", "-L", xcenv_bin_url, "-o", "binary.bin")
	err = xcenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcenv_src_url := "https://github.com/xcenv/xcenv/archive/refs/tags/v1.2.0.src.tar.gz"
	xcenv_cmd_src := exec.Command("curl", "-L", xcenv_src_url, "-o", "source.tar.gz")
	err = xcenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcenv_cmd_direct := exec.Command("./binary")
	err = xcenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
