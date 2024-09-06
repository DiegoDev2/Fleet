package main

// Jsign - Tool for signing Windows executable files, installers and scripts
// Homepage: https://ebourg.github.io/jsign/

import (
	"fmt"
	
	"os/exec"
)

func installJsign() {
	// Método 1: Descargar y extraer .tar.gz
	jsign_tar_url := "https://github.com/ebourg/jsign/archive/refs/tags/6.0.tar.gz"
	jsign_cmd_tar := exec.Command("curl", "-L", jsign_tar_url, "-o", "package.tar.gz")
	err := jsign_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsign_zip_url := "https://github.com/ebourg/jsign/archive/refs/tags/6.0.zip"
	jsign_cmd_zip := exec.Command("curl", "-L", jsign_zip_url, "-o", "package.zip")
	err = jsign_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsign_bin_url := "https://github.com/ebourg/jsign/archive/refs/tags/6.0.bin"
	jsign_cmd_bin := exec.Command("curl", "-L", jsign_bin_url, "-o", "binary.bin")
	err = jsign_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsign_src_url := "https://github.com/ebourg/jsign/archive/refs/tags/6.0.src.tar.gz"
	jsign_cmd_src := exec.Command("curl", "-L", jsign_src_url, "-o", "source.tar.gz")
	err = jsign_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsign_cmd_direct := exec.Command("./binary")
	err = jsign_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}
