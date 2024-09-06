package main

// Entityx - Fast, type-safe C++ Entity Component System
// Homepage: https://github.com/alecthomas/entityx

import (
	"fmt"
	
	"os/exec"
)

func installEntityx() {
	// Método 1: Descargar y extraer .tar.gz
	entityx_tar_url := "https://github.com/alecthomas/entityx/archive/refs/tags/1.3.0.tar.gz"
	entityx_cmd_tar := exec.Command("curl", "-L", entityx_tar_url, "-o", "package.tar.gz")
	err := entityx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	entityx_zip_url := "https://github.com/alecthomas/entityx/archive/refs/tags/1.3.0.zip"
	entityx_cmd_zip := exec.Command("curl", "-L", entityx_zip_url, "-o", "package.zip")
	err = entityx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	entityx_bin_url := "https://github.com/alecthomas/entityx/archive/refs/tags/1.3.0.bin"
	entityx_cmd_bin := exec.Command("curl", "-L", entityx_bin_url, "-o", "binary.bin")
	err = entityx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	entityx_src_url := "https://github.com/alecthomas/entityx/archive/refs/tags/1.3.0.src.tar.gz"
	entityx_cmd_src := exec.Command("curl", "-L", entityx_src_url, "-o", "source.tar.gz")
	err = entityx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	entityx_cmd_direct := exec.Command("./binary")
	err = entityx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
