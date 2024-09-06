package main

// Rhino - JavaScript engine
// Homepage: https://mozilla.github.io/rhino/

import (
	"fmt"
	
	"os/exec"
)

func installRhino() {
	// Método 1: Descargar y extraer .tar.gz
	rhino_tar_url := "https://github.com/mozilla/rhino/releases/download/Rhino1_7_15_Release/rhino-1.7.15.zip"
	rhino_cmd_tar := exec.Command("curl", "-L", rhino_tar_url, "-o", "package.tar.gz")
	err := rhino_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rhino_zip_url := "https://github.com/mozilla/rhino/releases/download/Rhino1_7_15_Release/rhino-1.7.15.zip"
	rhino_cmd_zip := exec.Command("curl", "-L", rhino_zip_url, "-o", "package.zip")
	err = rhino_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rhino_bin_url := "https://github.com/mozilla/rhino/releases/download/Rhino1_7_15_Release/rhino-1.7.15.zip"
	rhino_cmd_bin := exec.Command("curl", "-L", rhino_bin_url, "-o", "binary.bin")
	err = rhino_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rhino_src_url := "https://github.com/mozilla/rhino/releases/download/Rhino1_7_15_Release/rhino-1.7.15.zip"
	rhino_cmd_src := exec.Command("curl", "-L", rhino_src_url, "-o", "source.tar.gz")
	err = rhino_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rhino_cmd_direct := exec.Command("./binary")
	err = rhino_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
