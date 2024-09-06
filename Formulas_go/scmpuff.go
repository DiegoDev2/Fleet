package main

// Scmpuff - Adds numbered shortcuts for common git commands
// Homepage: https://mroth.github.io/scmpuff/

import (
	"fmt"
	
	"os/exec"
)

func installScmpuff() {
	// Método 1: Descargar y extraer .tar.gz
	scmpuff_tar_url := "https://github.com/mroth/scmpuff/archive/refs/tags/v0.5.0.tar.gz"
	scmpuff_cmd_tar := exec.Command("curl", "-L", scmpuff_tar_url, "-o", "package.tar.gz")
	err := scmpuff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scmpuff_zip_url := "https://github.com/mroth/scmpuff/archive/refs/tags/v0.5.0.zip"
	scmpuff_cmd_zip := exec.Command("curl", "-L", scmpuff_zip_url, "-o", "package.zip")
	err = scmpuff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scmpuff_bin_url := "https://github.com/mroth/scmpuff/archive/refs/tags/v0.5.0.bin"
	scmpuff_cmd_bin := exec.Command("curl", "-L", scmpuff_bin_url, "-o", "binary.bin")
	err = scmpuff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scmpuff_src_url := "https://github.com/mroth/scmpuff/archive/refs/tags/v0.5.0.src.tar.gz"
	scmpuff_cmd_src := exec.Command("curl", "-L", scmpuff_src_url, "-o", "source.tar.gz")
	err = scmpuff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scmpuff_cmd_direct := exec.Command("./binary")
	err = scmpuff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
