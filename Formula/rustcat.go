package main

// Rustcat - Modern Port listener and Reverse shell
// Homepage: https://github.com/robiot/rustcat

import (
	"fmt"
	
	"os/exec"
)

func installRustcat() {
	// Método 1: Descargar y extraer .tar.gz
	rustcat_tar_url := "https://github.com/robiot/rustcat/archive/refs/tags/v3.0.0.tar.gz"
	rustcat_cmd_tar := exec.Command("curl", "-L", rustcat_tar_url, "-o", "package.tar.gz")
	err := rustcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustcat_zip_url := "https://github.com/robiot/rustcat/archive/refs/tags/v3.0.0.zip"
	rustcat_cmd_zip := exec.Command("curl", "-L", rustcat_zip_url, "-o", "package.zip")
	err = rustcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustcat_bin_url := "https://github.com/robiot/rustcat/archive/refs/tags/v3.0.0.bin"
	rustcat_cmd_bin := exec.Command("curl", "-L", rustcat_bin_url, "-o", "binary.bin")
	err = rustcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustcat_src_url := "https://github.com/robiot/rustcat/archive/refs/tags/v3.0.0.src.tar.gz"
	rustcat_cmd_src := exec.Command("curl", "-L", rustcat_src_url, "-o", "source.tar.gz")
	err = rustcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustcat_cmd_direct := exec.Command("./binary")
	err = rustcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
