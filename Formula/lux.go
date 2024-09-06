package main

// Lux - Fast and simple video downloader
// Homepage: https://github.com/iawia002/lux

import (
	"fmt"
	
	"os/exec"
)

func installLux() {
	// Método 1: Descargar y extraer .tar.gz
	lux_tar_url := "https://github.com/iawia002/lux/archive/refs/tags/v0.24.1.tar.gz"
	lux_cmd_tar := exec.Command("curl", "-L", lux_tar_url, "-o", "package.tar.gz")
	err := lux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lux_zip_url := "https://github.com/iawia002/lux/archive/refs/tags/v0.24.1.zip"
	lux_cmd_zip := exec.Command("curl", "-L", lux_zip_url, "-o", "package.zip")
	err = lux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lux_bin_url := "https://github.com/iawia002/lux/archive/refs/tags/v0.24.1.bin"
	lux_cmd_bin := exec.Command("curl", "-L", lux_bin_url, "-o", "binary.bin")
	err = lux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lux_src_url := "https://github.com/iawia002/lux/archive/refs/tags/v0.24.1.src.tar.gz"
	lux_cmd_src := exec.Command("curl", "-L", lux_src_url, "-o", "source.tar.gz")
	err = lux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lux_cmd_direct := exec.Command("./binary")
	err = lux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
