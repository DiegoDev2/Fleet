package main

// Redo - Implements djb's redo: an alternative to make
// Homepage: https://redo.rtfd.io/

import (
	"fmt"
	
	"os/exec"
)

func installRedo() {
	// Método 1: Descargar y extraer .tar.gz
	redo_tar_url := "https://github.com/apenwarr/redo/archive/refs/tags/redo-0.42d.tar.gz"
	redo_cmd_tar := exec.Command("curl", "-L", redo_tar_url, "-o", "package.tar.gz")
	err := redo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redo_zip_url := "https://github.com/apenwarr/redo/archive/refs/tags/redo-0.42d.zip"
	redo_cmd_zip := exec.Command("curl", "-L", redo_zip_url, "-o", "package.zip")
	err = redo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redo_bin_url := "https://github.com/apenwarr/redo/archive/refs/tags/redo-0.42d.bin"
	redo_cmd_bin := exec.Command("curl", "-L", redo_bin_url, "-o", "binary.bin")
	err = redo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redo_src_url := "https://github.com/apenwarr/redo/archive/refs/tags/redo-0.42d.src.tar.gz"
	redo_cmd_src := exec.Command("curl", "-L", redo_src_url, "-o", "source.tar.gz")
	err = redo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redo_cmd_direct := exec.Command("./binary")
	err = redo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
