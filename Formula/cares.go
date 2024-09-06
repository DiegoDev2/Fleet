package main

// CAres - Asynchronous DNS library
// Homepage: https://c-ares.org/

import (
	"fmt"
	
	"os/exec"
)

func installCAres() {
	// Método 1: Descargar y extraer .tar.gz
	cares_tar_url := "https://github.com/c-ares/c-ares/releases/download/v1.33.1/c-ares-1.33.1.tar.gz"
	cares_cmd_tar := exec.Command("curl", "-L", cares_tar_url, "-o", "package.tar.gz")
	err := cares_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cares_zip_url := "https://github.com/c-ares/c-ares/releases/download/v1.33.1/c-ares-1.33.1.zip"
	cares_cmd_zip := exec.Command("curl", "-L", cares_zip_url, "-o", "package.zip")
	err = cares_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cares_bin_url := "https://github.com/c-ares/c-ares/releases/download/v1.33.1/c-ares-1.33.1.bin"
	cares_cmd_bin := exec.Command("curl", "-L", cares_bin_url, "-o", "binary.bin")
	err = cares_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cares_src_url := "https://github.com/c-ares/c-ares/releases/download/v1.33.1/c-ares-1.33.1.src.tar.gz"
	cares_cmd_src := exec.Command("curl", "-L", cares_src_url, "-o", "source.tar.gz")
	err = cares_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cares_cmd_direct := exec.Command("./binary")
	err = cares_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
