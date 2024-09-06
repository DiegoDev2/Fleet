package main

// Dotter - Dotfile manager and templater written in rust
// Homepage: https://github.com/SuperCuber/dotter

import (
	"fmt"
	
	"os/exec"
)

func installDotter() {
	// Método 1: Descargar y extraer .tar.gz
	dotter_tar_url := "https://github.com/SuperCuber/dotter/archive/refs/tags/v0.13.2.tar.gz"
	dotter_cmd_tar := exec.Command("curl", "-L", dotter_tar_url, "-o", "package.tar.gz")
	err := dotter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dotter_zip_url := "https://github.com/SuperCuber/dotter/archive/refs/tags/v0.13.2.zip"
	dotter_cmd_zip := exec.Command("curl", "-L", dotter_zip_url, "-o", "package.zip")
	err = dotter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dotter_bin_url := "https://github.com/SuperCuber/dotter/archive/refs/tags/v0.13.2.bin"
	dotter_cmd_bin := exec.Command("curl", "-L", dotter_bin_url, "-o", "binary.bin")
	err = dotter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dotter_src_url := "https://github.com/SuperCuber/dotter/archive/refs/tags/v0.13.2.src.tar.gz"
	dotter_cmd_src := exec.Command("curl", "-L", dotter_src_url, "-o", "source.tar.gz")
	err = dotter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dotter_cmd_direct := exec.Command("./binary")
	err = dotter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
