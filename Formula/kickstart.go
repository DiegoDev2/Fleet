package main

// Kickstart - Scaffolding tool to get new projects up and running quickly
// Homepage: https://github.com/Keats/kickstart

import (
	"fmt"
	
	"os/exec"
)

func installKickstart() {
	// Método 1: Descargar y extraer .tar.gz
	kickstart_tar_url := "https://github.com/Keats/kickstart/archive/refs/tags/v0.4.0.tar.gz"
	kickstart_cmd_tar := exec.Command("curl", "-L", kickstart_tar_url, "-o", "package.tar.gz")
	err := kickstart_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kickstart_zip_url := "https://github.com/Keats/kickstart/archive/refs/tags/v0.4.0.zip"
	kickstart_cmd_zip := exec.Command("curl", "-L", kickstart_zip_url, "-o", "package.zip")
	err = kickstart_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kickstart_bin_url := "https://github.com/Keats/kickstart/archive/refs/tags/v0.4.0.bin"
	kickstart_cmd_bin := exec.Command("curl", "-L", kickstart_bin_url, "-o", "binary.bin")
	err = kickstart_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kickstart_src_url := "https://github.com/Keats/kickstart/archive/refs/tags/v0.4.0.src.tar.gz"
	kickstart_cmd_src := exec.Command("curl", "-L", kickstart_src_url, "-o", "source.tar.gz")
	err = kickstart_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kickstart_cmd_direct := exec.Command("./binary")
	err = kickstart_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
