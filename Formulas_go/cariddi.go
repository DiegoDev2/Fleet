package main

// Cariddi - Scan for endpoints, secrets, API keys, file extensions, tokens and more
// Homepage: https://github.com/edoardottt/cariddi

import (
	"fmt"
	
	"os/exec"
)

func installCariddi() {
	// Método 1: Descargar y extraer .tar.gz
	cariddi_tar_url := "https://github.com/edoardottt/cariddi/archive/refs/tags/v1.3.5.tar.gz"
	cariddi_cmd_tar := exec.Command("curl", "-L", cariddi_tar_url, "-o", "package.tar.gz")
	err := cariddi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cariddi_zip_url := "https://github.com/edoardottt/cariddi/archive/refs/tags/v1.3.5.zip"
	cariddi_cmd_zip := exec.Command("curl", "-L", cariddi_zip_url, "-o", "package.zip")
	err = cariddi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cariddi_bin_url := "https://github.com/edoardottt/cariddi/archive/refs/tags/v1.3.5.bin"
	cariddi_cmd_bin := exec.Command("curl", "-L", cariddi_bin_url, "-o", "binary.bin")
	err = cariddi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cariddi_src_url := "https://github.com/edoardottt/cariddi/archive/refs/tags/v1.3.5.src.tar.gz"
	cariddi_cmd_src := exec.Command("curl", "-L", cariddi_src_url, "-o", "source.tar.gz")
	err = cariddi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cariddi_cmd_direct := exec.Command("./binary")
	err = cariddi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
