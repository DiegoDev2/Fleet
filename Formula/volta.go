package main

// Volta - JavaScript toolchain manager for reproducible environments
// Homepage: https://volta.sh

import (
	"fmt"
	
	"os/exec"
)

func installVolta() {
	// Método 1: Descargar y extraer .tar.gz
	volta_tar_url := "https://github.com/volta-cli/volta/archive/refs/tags/v2.0.1.tar.gz"
	volta_cmd_tar := exec.Command("curl", "-L", volta_tar_url, "-o", "package.tar.gz")
	err := volta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	volta_zip_url := "https://github.com/volta-cli/volta/archive/refs/tags/v2.0.1.zip"
	volta_cmd_zip := exec.Command("curl", "-L", volta_zip_url, "-o", "package.zip")
	err = volta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	volta_bin_url := "https://github.com/volta-cli/volta/archive/refs/tags/v2.0.1.bin"
	volta_cmd_bin := exec.Command("curl", "-L", volta_bin_url, "-o", "binary.bin")
	err = volta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	volta_src_url := "https://github.com/volta-cli/volta/archive/refs/tags/v2.0.1.src.tar.gz"
	volta_cmd_src := exec.Command("curl", "-L", volta_src_url, "-o", "source.tar.gz")
	err = volta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	volta_cmd_direct := exec.Command("./binary")
	err = volta_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
