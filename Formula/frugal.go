package main

// Frugal - Cross language code generator for creating scalable microservices
// Homepage: https://github.com/Workiva/frugal

import (
	"fmt"
	
	"os/exec"
)

func installFrugal() {
	// Método 1: Descargar y extraer .tar.gz
	frugal_tar_url := "https://github.com/Workiva/frugal/archive/refs/tags/v3.17.13.tar.gz"
	frugal_cmd_tar := exec.Command("curl", "-L", frugal_tar_url, "-o", "package.tar.gz")
	err := frugal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frugal_zip_url := "https://github.com/Workiva/frugal/archive/refs/tags/v3.17.13.zip"
	frugal_cmd_zip := exec.Command("curl", "-L", frugal_zip_url, "-o", "package.zip")
	err = frugal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frugal_bin_url := "https://github.com/Workiva/frugal/archive/refs/tags/v3.17.13.bin"
	frugal_cmd_bin := exec.Command("curl", "-L", frugal_bin_url, "-o", "binary.bin")
	err = frugal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frugal_src_url := "https://github.com/Workiva/frugal/archive/refs/tags/v3.17.13.src.tar.gz"
	frugal_cmd_src := exec.Command("curl", "-L", frugal_src_url, "-o", "source.tar.gz")
	err = frugal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frugal_cmd_direct := exec.Command("./binary")
	err = frugal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
