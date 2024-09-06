package main

// Purr - Versatile zsh CLI tool for viewing and searching through Android logcat output
// Homepage: https://github.com/google/purr

import (
	"fmt"
	
	"os/exec"
)

func installPurr() {
	// Método 1: Descargar y extraer .tar.gz
	purr_tar_url := "https://github.com/google/purr/archive/refs/tags/2.0.4.tar.gz"
	purr_cmd_tar := exec.Command("curl", "-L", purr_tar_url, "-o", "package.tar.gz")
	err := purr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	purr_zip_url := "https://github.com/google/purr/archive/refs/tags/2.0.4.zip"
	purr_cmd_zip := exec.Command("curl", "-L", purr_zip_url, "-o", "package.zip")
	err = purr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	purr_bin_url := "https://github.com/google/purr/archive/refs/tags/2.0.4.bin"
	purr_cmd_bin := exec.Command("curl", "-L", purr_bin_url, "-o", "binary.bin")
	err = purr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	purr_src_url := "https://github.com/google/purr/archive/refs/tags/2.0.4.src.tar.gz"
	purr_cmd_src := exec.Command("curl", "-L", purr_src_url, "-o", "source.tar.gz")
	err = purr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	purr_cmd_direct := exec.Command("./binary")
	err = purr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fzf")
	exec.Command("latte", "install", "fzf").Run()
}
