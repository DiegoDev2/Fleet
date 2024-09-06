package main

// RmImproved - Command-line deletion tool focused on safety, ergonomics, and performance
// Homepage: https://github.com/nivekuil/rip

import (
	"fmt"
	
	"os/exec"
)

func installRmImproved() {
	// Método 1: Descargar y extraer .tar.gz
	rmimproved_tar_url := "https://github.com/nivekuil/rip/archive/refs/tags/0.13.1.tar.gz"
	rmimproved_cmd_tar := exec.Command("curl", "-L", rmimproved_tar_url, "-o", "package.tar.gz")
	err := rmimproved_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rmimproved_zip_url := "https://github.com/nivekuil/rip/archive/refs/tags/0.13.1.zip"
	rmimproved_cmd_zip := exec.Command("curl", "-L", rmimproved_zip_url, "-o", "package.zip")
	err = rmimproved_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rmimproved_bin_url := "https://github.com/nivekuil/rip/archive/refs/tags/0.13.1.bin"
	rmimproved_cmd_bin := exec.Command("curl", "-L", rmimproved_bin_url, "-o", "binary.bin")
	err = rmimproved_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rmimproved_src_url := "https://github.com/nivekuil/rip/archive/refs/tags/0.13.1.src.tar.gz"
	rmimproved_cmd_src := exec.Command("curl", "-L", rmimproved_src_url, "-o", "source.tar.gz")
	err = rmimproved_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rmimproved_cmd_direct := exec.Command("./binary")
	err = rmimproved_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
