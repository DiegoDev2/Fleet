package main

// Desk - Lightweight workspace manager for the shell
// Homepage: https://github.com/jamesob/desk

import (
	"fmt"
	
	"os/exec"
)

func installDesk() {
	// Método 1: Descargar y extraer .tar.gz
	desk_tar_url := "https://github.com/jamesob/desk/archive/refs/tags/v0.6.0.tar.gz"
	desk_cmd_tar := exec.Command("curl", "-L", desk_tar_url, "-o", "package.tar.gz")
	err := desk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	desk_zip_url := "https://github.com/jamesob/desk/archive/refs/tags/v0.6.0.zip"
	desk_cmd_zip := exec.Command("curl", "-L", desk_zip_url, "-o", "package.zip")
	err = desk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	desk_bin_url := "https://github.com/jamesob/desk/archive/refs/tags/v0.6.0.bin"
	desk_cmd_bin := exec.Command("curl", "-L", desk_bin_url, "-o", "binary.bin")
	err = desk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	desk_src_url := "https://github.com/jamesob/desk/archive/refs/tags/v0.6.0.src.tar.gz"
	desk_cmd_src := exec.Command("curl", "-L", desk_src_url, "-o", "source.tar.gz")
	err = desk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	desk_cmd_direct := exec.Command("./binary")
	err = desk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
