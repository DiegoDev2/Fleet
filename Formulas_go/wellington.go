package main

// Wellington - Project-focused tool to manage Sass and spriting
// Homepage: https://github.com/wellington/wellington

import (
	"fmt"
	
	"os/exec"
)

func installWellington() {
	// Método 1: Descargar y extraer .tar.gz
	wellington_tar_url := "https://github.com/wellington/wellington/archive/refs/tags/v1.0.5.tar.gz"
	wellington_cmd_tar := exec.Command("curl", "-L", wellington_tar_url, "-o", "package.tar.gz")
	err := wellington_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wellington_zip_url := "https://github.com/wellington/wellington/archive/refs/tags/v1.0.5.zip"
	wellington_cmd_zip := exec.Command("curl", "-L", wellington_zip_url, "-o", "package.zip")
	err = wellington_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wellington_bin_url := "https://github.com/wellington/wellington/archive/refs/tags/v1.0.5.bin"
	wellington_cmd_bin := exec.Command("curl", "-L", wellington_bin_url, "-o", "binary.bin")
	err = wellington_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wellington_src_url := "https://github.com/wellington/wellington/archive/refs/tags/v1.0.5.src.tar.gz"
	wellington_cmd_src := exec.Command("curl", "-L", wellington_src_url, "-o", "source.tar.gz")
	err = wellington_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wellington_cmd_direct := exec.Command("./binary")
	err = wellington_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.19")
exec.Command("latte", "install", "go@1.19")
}
