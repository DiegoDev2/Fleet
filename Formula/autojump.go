package main

// Autojump - Shell extension to jump to frequently used directories
// Homepage: https://github.com/wting/autojump

import (
	"fmt"
	
	"os/exec"
)

func installAutojump() {
	// Método 1: Descargar y extraer .tar.gz
	autojump_tar_url := "https://github.com/wting/autojump/archive/refs/tags/release-v22.5.3.tar.gz"
	autojump_cmd_tar := exec.Command("curl", "-L", autojump_tar_url, "-o", "package.tar.gz")
	err := autojump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autojump_zip_url := "https://github.com/wting/autojump/archive/refs/tags/release-v22.5.3.zip"
	autojump_cmd_zip := exec.Command("curl", "-L", autojump_zip_url, "-o", "package.zip")
	err = autojump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autojump_bin_url := "https://github.com/wting/autojump/archive/refs/tags/release-v22.5.3.bin"
	autojump_cmd_bin := exec.Command("curl", "-L", autojump_bin_url, "-o", "binary.bin")
	err = autojump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autojump_src_url := "https://github.com/wting/autojump/archive/refs/tags/release-v22.5.3.src.tar.gz"
	autojump_cmd_src := exec.Command("curl", "-L", autojump_src_url, "-o", "source.tar.gz")
	err = autojump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autojump_cmd_direct := exec.Command("./binary")
	err = autojump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
