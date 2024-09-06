package main

// Deadfinder - Finds broken links
// Homepage: https://rubygems.org/gems/deadfinder

import (
	"fmt"
	
	"os/exec"
)

func installDeadfinder() {
	// Método 1: Descargar y extraer .tar.gz
	deadfinder_tar_url := "https://github.com/hahwul/deadfinder/archive/refs/tags/1.3.4.tar.gz"
	deadfinder_cmd_tar := exec.Command("curl", "-L", deadfinder_tar_url, "-o", "package.tar.gz")
	err := deadfinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	deadfinder_zip_url := "https://github.com/hahwul/deadfinder/archive/refs/tags/1.3.4.zip"
	deadfinder_cmd_zip := exec.Command("curl", "-L", deadfinder_zip_url, "-o", "package.zip")
	err = deadfinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	deadfinder_bin_url := "https://github.com/hahwul/deadfinder/archive/refs/tags/1.3.4.bin"
	deadfinder_cmd_bin := exec.Command("curl", "-L", deadfinder_bin_url, "-o", "binary.bin")
	err = deadfinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	deadfinder_src_url := "https://github.com/hahwul/deadfinder/archive/refs/tags/1.3.4.src.tar.gz"
	deadfinder_cmd_src := exec.Command("curl", "-L", deadfinder_src_url, "-o", "source.tar.gz")
	err = deadfinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	deadfinder_cmd_direct := exec.Command("./binary")
	err = deadfinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
