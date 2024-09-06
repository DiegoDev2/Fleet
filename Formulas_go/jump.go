package main

// Jump - Helps you navigate your file system faster by learning your habits
// Homepage: https://github.com/gsamokovarov/jump

import (
	"fmt"
	
	"os/exec"
)

func installJump() {
	// Método 1: Descargar y extraer .tar.gz
	jump_tar_url := "https://github.com/gsamokovarov/jump/archive/refs/tags/v0.51.0.tar.gz"
	jump_cmd_tar := exec.Command("curl", "-L", jump_tar_url, "-o", "package.tar.gz")
	err := jump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jump_zip_url := "https://github.com/gsamokovarov/jump/archive/refs/tags/v0.51.0.zip"
	jump_cmd_zip := exec.Command("curl", "-L", jump_zip_url, "-o", "package.zip")
	err = jump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jump_bin_url := "https://github.com/gsamokovarov/jump/archive/refs/tags/v0.51.0.bin"
	jump_cmd_bin := exec.Command("curl", "-L", jump_bin_url, "-o", "binary.bin")
	err = jump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jump_src_url := "https://github.com/gsamokovarov/jump/archive/refs/tags/v0.51.0.src.tar.gz"
	jump_cmd_src := exec.Command("curl", "-L", jump_src_url, "-o", "source.tar.gz")
	err = jump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jump_cmd_direct := exec.Command("./binary")
	err = jump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
