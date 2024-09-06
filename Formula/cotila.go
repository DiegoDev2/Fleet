package main

// Cotila - Compile-time linear algebra system for C++
// Homepage: https://github.com/calebzulawski/cotila

import (
	"fmt"
	
	"os/exec"
)

func installCotila() {
	// Método 1: Descargar y extraer .tar.gz
	cotila_tar_url := "https://github.com/calebzulawski/cotila/archive/refs/tags/1.2.1.tar.gz"
	cotila_cmd_tar := exec.Command("curl", "-L", cotila_tar_url, "-o", "package.tar.gz")
	err := cotila_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cotila_zip_url := "https://github.com/calebzulawski/cotila/archive/refs/tags/1.2.1.zip"
	cotila_cmd_zip := exec.Command("curl", "-L", cotila_zip_url, "-o", "package.zip")
	err = cotila_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cotila_bin_url := "https://github.com/calebzulawski/cotila/archive/refs/tags/1.2.1.bin"
	cotila_cmd_bin := exec.Command("curl", "-L", cotila_bin_url, "-o", "binary.bin")
	err = cotila_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cotila_src_url := "https://github.com/calebzulawski/cotila/archive/refs/tags/1.2.1.src.tar.gz"
	cotila_cmd_src := exec.Command("curl", "-L", cotila_src_url, "-o", "source.tar.gz")
	err = cotila_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cotila_cmd_direct := exec.Command("./binary")
	err = cotila_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
