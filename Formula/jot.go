package main

// Jot - Rapid note management for the terminal
// Homepage: https://github.com/shashwatah/jot

import (
	"fmt"
	
	"os/exec"
)

func installJot() {
	// Método 1: Descargar y extraer .tar.gz
	jot_tar_url := "https://github.com/shashwatah/jot/archive/refs/tags/v0.1.2.tar.gz"
	jot_cmd_tar := exec.Command("curl", "-L", jot_tar_url, "-o", "package.tar.gz")
	err := jot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jot_zip_url := "https://github.com/shashwatah/jot/archive/refs/tags/v0.1.2.zip"
	jot_cmd_zip := exec.Command("curl", "-L", jot_zip_url, "-o", "package.zip")
	err = jot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jot_bin_url := "https://github.com/shashwatah/jot/archive/refs/tags/v0.1.2.bin"
	jot_cmd_bin := exec.Command("curl", "-L", jot_bin_url, "-o", "binary.bin")
	err = jot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jot_src_url := "https://github.com/shashwatah/jot/archive/refs/tags/v0.1.2.src.tar.gz"
	jot_cmd_src := exec.Command("curl", "-L", jot_src_url, "-o", "source.tar.gz")
	err = jot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jot_cmd_direct := exec.Command("./binary")
	err = jot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
