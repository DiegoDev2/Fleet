package main

// Fileicon - macOS CLI for managing custom icons for files and folders
// Homepage: https://github.com/mklement0/fileicon

import (
	"fmt"
	
	"os/exec"
)

func installFileicon() {
	// Método 1: Descargar y extraer .tar.gz
	fileicon_tar_url := "https://github.com/mklement0/fileicon/archive/refs/tags/v0.3.4.tar.gz"
	fileicon_cmd_tar := exec.Command("curl", "-L", fileicon_tar_url, "-o", "package.tar.gz")
	err := fileicon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fileicon_zip_url := "https://github.com/mklement0/fileicon/archive/refs/tags/v0.3.4.zip"
	fileicon_cmd_zip := exec.Command("curl", "-L", fileicon_zip_url, "-o", "package.zip")
	err = fileicon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fileicon_bin_url := "https://github.com/mklement0/fileicon/archive/refs/tags/v0.3.4.bin"
	fileicon_cmd_bin := exec.Command("curl", "-L", fileicon_bin_url, "-o", "binary.bin")
	err = fileicon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fileicon_src_url := "https://github.com/mklement0/fileicon/archive/refs/tags/v0.3.4.src.tar.gz"
	fileicon_cmd_src := exec.Command("curl", "-L", fileicon_src_url, "-o", "source.tar.gz")
	err = fileicon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fileicon_cmd_direct := exec.Command("./binary")
	err = fileicon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
