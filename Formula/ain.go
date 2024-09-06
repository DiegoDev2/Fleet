package main

// Ain - HTTP API client for the terminal
// Homepage: https://github.com/jonaslu/ain

import (
	"fmt"
	
	"os/exec"
)

func installAin() {
	// Método 1: Descargar y extraer .tar.gz
	ain_tar_url := "https://github.com/jonaslu/ain/archive/refs/tags/v1.5.0.tar.gz"
	ain_cmd_tar := exec.Command("curl", "-L", ain_tar_url, "-o", "package.tar.gz")
	err := ain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ain_zip_url := "https://github.com/jonaslu/ain/archive/refs/tags/v1.5.0.zip"
	ain_cmd_zip := exec.Command("curl", "-L", ain_zip_url, "-o", "package.zip")
	err = ain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ain_bin_url := "https://github.com/jonaslu/ain/archive/refs/tags/v1.5.0.bin"
	ain_cmd_bin := exec.Command("curl", "-L", ain_bin_url, "-o", "binary.bin")
	err = ain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ain_src_url := "https://github.com/jonaslu/ain/archive/refs/tags/v1.5.0.src.tar.gz"
	ain_cmd_src := exec.Command("curl", "-L", ain_src_url, "-o", "source.tar.gz")
	err = ain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ain_cmd_direct := exec.Command("./binary")
	err = ain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
