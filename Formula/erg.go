package main

// Erg - Statically typed language that can deeply improve the Python ecosystem
// Homepage: https://github.com/erg-lang/erg

import (
	"fmt"
	
	"os/exec"
)

func installErg() {
	// Método 1: Descargar y extraer .tar.gz
	erg_tar_url := "https://github.com/erg-lang/erg/archive/refs/tags/v0.6.42.tar.gz"
	erg_cmd_tar := exec.Command("curl", "-L", erg_tar_url, "-o", "package.tar.gz")
	err := erg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erg_zip_url := "https://github.com/erg-lang/erg/archive/refs/tags/v0.6.42.zip"
	erg_cmd_zip := exec.Command("curl", "-L", erg_zip_url, "-o", "package.zip")
	err = erg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erg_bin_url := "https://github.com/erg-lang/erg/archive/refs/tags/v0.6.42.bin"
	erg_cmd_bin := exec.Command("curl", "-L", erg_bin_url, "-o", "binary.bin")
	err = erg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erg_src_url := "https://github.com/erg-lang/erg/archive/refs/tags/v0.6.42.src.tar.gz"
	erg_cmd_src := exec.Command("curl", "-L", erg_src_url, "-o", "source.tar.gz")
	err = erg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erg_cmd_direct := exec.Command("./binary")
	err = erg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
