package main

// Gomplate - Command-line Golang template processor
// Homepage: https://gomplate.ca/

import (
	"fmt"
	
	"os/exec"
)

func installGomplate() {
	// Método 1: Descargar y extraer .tar.gz
	gomplate_tar_url := "https://github.com/hairyhenderson/gomplate/archive/refs/tags/v4.1.0.tar.gz"
	gomplate_cmd_tar := exec.Command("curl", "-L", gomplate_tar_url, "-o", "package.tar.gz")
	err := gomplate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gomplate_zip_url := "https://github.com/hairyhenderson/gomplate/archive/refs/tags/v4.1.0.zip"
	gomplate_cmd_zip := exec.Command("curl", "-L", gomplate_zip_url, "-o", "package.zip")
	err = gomplate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gomplate_bin_url := "https://github.com/hairyhenderson/gomplate/archive/refs/tags/v4.1.0.bin"
	gomplate_cmd_bin := exec.Command("curl", "-L", gomplate_bin_url, "-o", "binary.bin")
	err = gomplate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gomplate_src_url := "https://github.com/hairyhenderson/gomplate/archive/refs/tags/v4.1.0.src.tar.gz"
	gomplate_cmd_src := exec.Command("curl", "-L", gomplate_src_url, "-o", "source.tar.gz")
	err = gomplate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gomplate_cmd_direct := exec.Command("./binary")
	err = gomplate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
