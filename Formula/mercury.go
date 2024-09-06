package main

// Mercury - Logic/functional programming language
// Homepage: https://mercurylang.org/

import (
	"fmt"
	
	"os/exec"
)

func installMercury() {
	// Método 1: Descargar y extraer .tar.gz
	mercury_tar_url := "https://dl.mercurylang.org/release/mercury-srcdist-22.01.8.tar.gz"
	mercury_cmd_tar := exec.Command("curl", "-L", mercury_tar_url, "-o", "package.tar.gz")
	err := mercury_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mercury_zip_url := "https://dl.mercurylang.org/release/mercury-srcdist-22.01.8.zip"
	mercury_cmd_zip := exec.Command("curl", "-L", mercury_zip_url, "-o", "package.zip")
	err = mercury_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mercury_bin_url := "https://dl.mercurylang.org/release/mercury-srcdist-22.01.8.bin"
	mercury_cmd_bin := exec.Command("curl", "-L", mercury_bin_url, "-o", "binary.bin")
	err = mercury_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mercury_src_url := "https://dl.mercurylang.org/release/mercury-srcdist-22.01.8.src.tar.gz"
	mercury_cmd_src := exec.Command("curl", "-L", mercury_src_url, "-o", "source.tar.gz")
	err = mercury_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mercury_cmd_direct := exec.Command("./binary")
	err = mercury_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
