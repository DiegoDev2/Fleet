package main

// Argc - Easily create and use cli based on bash script
// Homepage: https://github.com/sigoden/argc

import (
	"fmt"
	
	"os/exec"
)

func installArgc() {
	// Método 1: Descargar y extraer .tar.gz
	argc_tar_url := "https://github.com/sigoden/argc/archive/refs/tags/v1.20.1.tar.gz"
	argc_cmd_tar := exec.Command("curl", "-L", argc_tar_url, "-o", "package.tar.gz")
	err := argc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argc_zip_url := "https://github.com/sigoden/argc/archive/refs/tags/v1.20.1.zip"
	argc_cmd_zip := exec.Command("curl", "-L", argc_zip_url, "-o", "package.zip")
	err = argc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argc_bin_url := "https://github.com/sigoden/argc/archive/refs/tags/v1.20.1.bin"
	argc_cmd_bin := exec.Command("curl", "-L", argc_bin_url, "-o", "binary.bin")
	err = argc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argc_src_url := "https://github.com/sigoden/argc/archive/refs/tags/v1.20.1.src.tar.gz"
	argc_cmd_src := exec.Command("curl", "-L", argc_src_url, "-o", "source.tar.gz")
	err = argc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argc_cmd_direct := exec.Command("./binary")
	err = argc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
