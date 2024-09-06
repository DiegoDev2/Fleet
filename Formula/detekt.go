package main

// Detekt - Static code analysis for Kotlin
// Homepage: https://github.com/detekt/detekt

import (
	"fmt"
	
	"os/exec"
)

func installDetekt() {
	// Método 1: Descargar y extraer .tar.gz
	detekt_tar_url := "https://github.com/detekt/detekt/releases/download/v1.23.6/detekt-cli-1.23.6-all.jar"
	detekt_cmd_tar := exec.Command("curl", "-L", detekt_tar_url, "-o", "package.tar.gz")
	err := detekt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	detekt_zip_url := "https://github.com/detekt/detekt/releases/download/v1.23.6/detekt-cli-1.23.6-all.jar"
	detekt_cmd_zip := exec.Command("curl", "-L", detekt_zip_url, "-o", "package.zip")
	err = detekt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	detekt_bin_url := "https://github.com/detekt/detekt/releases/download/v1.23.6/detekt-cli-1.23.6-all.jar"
	detekt_cmd_bin := exec.Command("curl", "-L", detekt_bin_url, "-o", "binary.bin")
	err = detekt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	detekt_src_url := "https://github.com/detekt/detekt/releases/download/v1.23.6/detekt-cli-1.23.6-all.jar"
	detekt_cmd_src := exec.Command("curl", "-L", detekt_src_url, "-o", "source.tar.gz")
	err = detekt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	detekt_cmd_direct := exec.Command("./binary")
	err = detekt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}
