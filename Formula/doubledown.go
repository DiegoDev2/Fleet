package main

// Doubledown - Sync local changes to a remote directory
// Homepage: https://github.com/devstructure/doubledown

import (
	"fmt"
	
	"os/exec"
)

func installDoubledown() {
	// Método 1: Descargar y extraer .tar.gz
	doubledown_tar_url := "https://github.com/devstructure/doubledown/archive/refs/tags/v0.0.2.tar.gz"
	doubledown_cmd_tar := exec.Command("curl", "-L", doubledown_tar_url, "-o", "package.tar.gz")
	err := doubledown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doubledown_zip_url := "https://github.com/devstructure/doubledown/archive/refs/tags/v0.0.2.zip"
	doubledown_cmd_zip := exec.Command("curl", "-L", doubledown_zip_url, "-o", "package.zip")
	err = doubledown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doubledown_bin_url := "https://github.com/devstructure/doubledown/archive/refs/tags/v0.0.2.bin"
	doubledown_cmd_bin := exec.Command("curl", "-L", doubledown_bin_url, "-o", "binary.bin")
	err = doubledown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doubledown_src_url := "https://github.com/devstructure/doubledown/archive/refs/tags/v0.0.2.src.tar.gz"
	doubledown_cmd_src := exec.Command("curl", "-L", doubledown_src_url, "-o", "source.tar.gz")
	err = doubledown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doubledown_cmd_direct := exec.Command("./binary")
	err = doubledown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
