package main

// Eless - Better `less` using Emacs view-mode and Bash
// Homepage: https://eless.scripter.co/

import (
	"fmt"
	
	"os/exec"
)

func installEless() {
	// Método 1: Descargar y extraer .tar.gz
	eless_tar_url := "https://github.com/kaushalmodi/eless/archive/refs/tags/v0.7.tar.gz"
	eless_cmd_tar := exec.Command("curl", "-L", eless_tar_url, "-o", "package.tar.gz")
	err := eless_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eless_zip_url := "https://github.com/kaushalmodi/eless/archive/refs/tags/v0.7.zip"
	eless_cmd_zip := exec.Command("curl", "-L", eless_zip_url, "-o", "package.zip")
	err = eless_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eless_bin_url := "https://github.com/kaushalmodi/eless/archive/refs/tags/v0.7.bin"
	eless_cmd_bin := exec.Command("curl", "-L", eless_bin_url, "-o", "binary.bin")
	err = eless_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eless_src_url := "https://github.com/kaushalmodi/eless/archive/refs/tags/v0.7.src.tar.gz"
	eless_cmd_src := exec.Command("curl", "-L", eless_src_url, "-o", "source.tar.gz")
	err = eless_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eless_cmd_direct := exec.Command("./binary")
	err = eless_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
}
