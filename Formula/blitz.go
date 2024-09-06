package main

// Blitz - Multi-dimensional array library for C++
// Homepage: https://github.com/blitzpp/blitz/wiki

import (
	"fmt"
	
	"os/exec"
)

func installBlitz() {
	// Método 1: Descargar y extraer .tar.gz
	blitz_tar_url := "https://github.com/blitzpp/blitz/archive/refs/tags/1.0.2.tar.gz"
	blitz_cmd_tar := exec.Command("curl", "-L", blitz_tar_url, "-o", "package.tar.gz")
	err := blitz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blitz_zip_url := "https://github.com/blitzpp/blitz/archive/refs/tags/1.0.2.zip"
	blitz_cmd_zip := exec.Command("curl", "-L", blitz_zip_url, "-o", "package.zip")
	err = blitz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blitz_bin_url := "https://github.com/blitzpp/blitz/archive/refs/tags/1.0.2.bin"
	blitz_cmd_bin := exec.Command("curl", "-L", blitz_bin_url, "-o", "binary.bin")
	err = blitz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blitz_src_url := "https://github.com/blitzpp/blitz/archive/refs/tags/1.0.2.src.tar.gz"
	blitz_cmd_src := exec.Command("curl", "-L", blitz_src_url, "-o", "source.tar.gz")
	err = blitz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blitz_cmd_direct := exec.Command("./binary")
	err = blitz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
