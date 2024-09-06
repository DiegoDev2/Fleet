package main

// Rgxg - C library and command-line tool to generate (extended) regular expressions
// Homepage: https://rgxg.github.io

import (
	"fmt"
	
	"os/exec"
)

func installRgxg() {
	// Método 1: Descargar y extraer .tar.gz
	rgxg_tar_url := "https://github.com/rgxg/rgxg/releases/download/v0.1.2/rgxg-0.1.2.tar.gz"
	rgxg_cmd_tar := exec.Command("curl", "-L", rgxg_tar_url, "-o", "package.tar.gz")
	err := rgxg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rgxg_zip_url := "https://github.com/rgxg/rgxg/releases/download/v0.1.2/rgxg-0.1.2.zip"
	rgxg_cmd_zip := exec.Command("curl", "-L", rgxg_zip_url, "-o", "package.zip")
	err = rgxg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rgxg_bin_url := "https://github.com/rgxg/rgxg/releases/download/v0.1.2/rgxg-0.1.2.bin"
	rgxg_cmd_bin := exec.Command("curl", "-L", rgxg_bin_url, "-o", "binary.bin")
	err = rgxg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rgxg_src_url := "https://github.com/rgxg/rgxg/releases/download/v0.1.2/rgxg-0.1.2.src.tar.gz"
	rgxg_cmd_src := exec.Command("curl", "-L", rgxg_src_url, "-o", "source.tar.gz")
	err = rgxg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rgxg_cmd_direct := exec.Command("./binary")
	err = rgxg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
