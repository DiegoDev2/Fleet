package main

// Chezscheme - Implementation of the Chez Scheme language
// Homepage: https://cisco.github.io/ChezScheme/

import (
	"fmt"
	
	"os/exec"
)

func installChezscheme() {
	// Método 1: Descargar y extraer .tar.gz
	chezscheme_tar_url := "https://github.com/cisco/ChezScheme/releases/download/v10.0.0/csv10.0.0.tar.gz"
	chezscheme_cmd_tar := exec.Command("curl", "-L", chezscheme_tar_url, "-o", "package.tar.gz")
	err := chezscheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chezscheme_zip_url := "https://github.com/cisco/ChezScheme/releases/download/v10.0.0/csv10.0.0.zip"
	chezscheme_cmd_zip := exec.Command("curl", "-L", chezscheme_zip_url, "-o", "package.zip")
	err = chezscheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chezscheme_bin_url := "https://github.com/cisco/ChezScheme/releases/download/v10.0.0/csv10.0.0.bin"
	chezscheme_cmd_bin := exec.Command("curl", "-L", chezscheme_bin_url, "-o", "binary.bin")
	err = chezscheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chezscheme_src_url := "https://github.com/cisco/ChezScheme/releases/download/v10.0.0/csv10.0.0.src.tar.gz"
	chezscheme_cmd_src := exec.Command("curl", "-L", chezscheme_src_url, "-o", "source.tar.gz")
	err = chezscheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chezscheme_cmd_direct := exec.Command("./binary")
	err = chezscheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: xterm")
	exec.Command("latte", "install", "xterm").Run()
}
