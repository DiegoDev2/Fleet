package main

// Stdman - Formatted C++ stdlib man pages from cppreference.com
// Homepage: https://github.com/jeaye/stdman

import (
	"fmt"
	
	"os/exec"
)

func installStdman() {
	// Método 1: Descargar y extraer .tar.gz
	stdman_tar_url := "https://github.com/jeaye/stdman/archive/refs/tags/2024.07.05.tar.gz"
	stdman_cmd_tar := exec.Command("curl", "-L", stdman_tar_url, "-o", "package.tar.gz")
	err := stdman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stdman_zip_url := "https://github.com/jeaye/stdman/archive/refs/tags/2024.07.05.zip"
	stdman_cmd_zip := exec.Command("curl", "-L", stdman_zip_url, "-o", "package.zip")
	err = stdman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stdman_bin_url := "https://github.com/jeaye/stdman/archive/refs/tags/2024.07.05.bin"
	stdman_cmd_bin := exec.Command("curl", "-L", stdman_bin_url, "-o", "binary.bin")
	err = stdman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stdman_src_url := "https://github.com/jeaye/stdman/archive/refs/tags/2024.07.05.src.tar.gz"
	stdman_cmd_src := exec.Command("curl", "-L", stdman_src_url, "-o", "source.tar.gz")
	err = stdman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stdman_cmd_direct := exec.Command("./binary")
	err = stdman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: man-db")
	exec.Command("latte", "install", "man-db").Run()
}
