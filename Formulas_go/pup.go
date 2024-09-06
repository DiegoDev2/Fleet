package main

// Pup - Parse HTML at the command-line
// Homepage: https://github.com/EricChiang/pup

import (
	"fmt"
	
	"os/exec"
)

func installPup() {
	// Método 1: Descargar y extraer .tar.gz
	pup_tar_url := "https://github.com/ericchiang/pup/archive/refs/tags/v0.4.0.tar.gz"
	pup_cmd_tar := exec.Command("curl", "-L", pup_tar_url, "-o", "package.tar.gz")
	err := pup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pup_zip_url := "https://github.com/ericchiang/pup/archive/refs/tags/v0.4.0.zip"
	pup_cmd_zip := exec.Command("curl", "-L", pup_zip_url, "-o", "package.zip")
	err = pup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pup_bin_url := "https://github.com/ericchiang/pup/archive/refs/tags/v0.4.0.bin"
	pup_cmd_bin := exec.Command("curl", "-L", pup_bin_url, "-o", "binary.bin")
	err = pup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pup_src_url := "https://github.com/ericchiang/pup/archive/refs/tags/v0.4.0.src.tar.gz"
	pup_cmd_src := exec.Command("curl", "-L", pup_src_url, "-o", "source.tar.gz")
	err = pup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pup_cmd_direct := exec.Command("./binary")
	err = pup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: gox")
exec.Command("latte", "install", "gox")
}
