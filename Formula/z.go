package main

// Z - Tracks most-used directories to make cd smarter
// Homepage: https://github.com/rupa/z

import (
	"fmt"
	
	"os/exec"
)

func installZ() {
	// Método 1: Descargar y extraer .tar.gz
	z_tar_url := "https://github.com/rupa/z/archive/refs/tags/v1.12.tar.gz"
	z_cmd_tar := exec.Command("curl", "-L", z_tar_url, "-o", "package.tar.gz")
	err := z_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	z_zip_url := "https://github.com/rupa/z/archive/refs/tags/v1.12.zip"
	z_cmd_zip := exec.Command("curl", "-L", z_zip_url, "-o", "package.zip")
	err = z_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	z_bin_url := "https://github.com/rupa/z/archive/refs/tags/v1.12.bin"
	z_cmd_bin := exec.Command("curl", "-L", z_bin_url, "-o", "binary.bin")
	err = z_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	z_src_url := "https://github.com/rupa/z/archive/refs/tags/v1.12.src.tar.gz"
	z_cmd_src := exec.Command("curl", "-L", z_src_url, "-o", "source.tar.gz")
	err = z_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	z_cmd_direct := exec.Command("./binary")
	err = z_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
