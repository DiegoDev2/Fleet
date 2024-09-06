package main

// Xcv - Cut, copy and paste files with Bash
// Homepage: https://github.com/busterc/xcv

import (
	"fmt"
	
	"os/exec"
)

func installXcv() {
	// Método 1: Descargar y extraer .tar.gz
	xcv_tar_url := "https://github.com/busterc/xcv/archive/refs/tags/v1.0.1.tar.gz"
	xcv_cmd_tar := exec.Command("curl", "-L", xcv_tar_url, "-o", "package.tar.gz")
	err := xcv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcv_zip_url := "https://github.com/busterc/xcv/archive/refs/tags/v1.0.1.zip"
	xcv_cmd_zip := exec.Command("curl", "-L", xcv_zip_url, "-o", "package.zip")
	err = xcv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcv_bin_url := "https://github.com/busterc/xcv/archive/refs/tags/v1.0.1.bin"
	xcv_cmd_bin := exec.Command("curl", "-L", xcv_bin_url, "-o", "binary.bin")
	err = xcv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcv_src_url := "https://github.com/busterc/xcv/archive/refs/tags/v1.0.1.src.tar.gz"
	xcv_cmd_src := exec.Command("curl", "-L", xcv_src_url, "-o", "source.tar.gz")
	err = xcv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcv_cmd_direct := exec.Command("./binary")
	err = xcv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
