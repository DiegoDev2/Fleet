package main

// Jove - Emacs-style editor with vi-like memory, CPU, and size requirements
// Homepage: https://directory.fsf.org/wiki/Jove

import (
	"fmt"
	
	"os/exec"
)

func installJove() {
	// Método 1: Descargar y extraer .tar.gz
	jove_tar_url := "https://github.com/jonmacs/jove/archive/refs/tags/4.17.5.3.tar.gz"
	jove_cmd_tar := exec.Command("curl", "-L", jove_tar_url, "-o", "package.tar.gz")
	err := jove_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jove_zip_url := "https://github.com/jonmacs/jove/archive/refs/tags/4.17.5.3.zip"
	jove_cmd_zip := exec.Command("curl", "-L", jove_zip_url, "-o", "package.zip")
	err = jove_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jove_bin_url := "https://github.com/jonmacs/jove/archive/refs/tags/4.17.5.3.bin"
	jove_cmd_bin := exec.Command("curl", "-L", jove_bin_url, "-o", "binary.bin")
	err = jove_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jove_src_url := "https://github.com/jonmacs/jove/archive/refs/tags/4.17.5.3.src.tar.gz"
	jove_cmd_src := exec.Command("curl", "-L", jove_src_url, "-o", "source.tar.gz")
	err = jove_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jove_cmd_direct := exec.Command("./binary")
	err = jove_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: groff")
	exec.Command("latte", "install", "groff").Run()
}
