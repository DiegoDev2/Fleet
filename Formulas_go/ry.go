package main

// Ry - Ruby virtual env tool
// Homepage: https://github.com/jneen/ry

import (
	"fmt"
	
	"os/exec"
)

func installRy() {
	// Método 1: Descargar y extraer .tar.gz
	ry_tar_url := "https://github.com/jneen/ry/archive/refs/tags/v0.5.2.tar.gz"
	ry_cmd_tar := exec.Command("curl", "-L", ry_tar_url, "-o", "package.tar.gz")
	err := ry_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ry_zip_url := "https://github.com/jneen/ry/archive/refs/tags/v0.5.2.zip"
	ry_cmd_zip := exec.Command("curl", "-L", ry_zip_url, "-o", "package.zip")
	err = ry_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ry_bin_url := "https://github.com/jneen/ry/archive/refs/tags/v0.5.2.bin"
	ry_cmd_bin := exec.Command("curl", "-L", ry_bin_url, "-o", "binary.bin")
	err = ry_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ry_src_url := "https://github.com/jneen/ry/archive/refs/tags/v0.5.2.src.tar.gz"
	ry_cmd_src := exec.Command("curl", "-L", ry_src_url, "-o", "source.tar.gz")
	err = ry_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ry_cmd_direct := exec.Command("./binary")
	err = ry_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash-completion")
exec.Command("latte", "install", "bash-completion")
	fmt.Println("Instalando dependencia: ruby-build")
exec.Command("latte", "install", "ruby-build")
}
