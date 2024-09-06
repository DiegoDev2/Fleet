package main

// ChrubyFish - Thin wrapper around chruby to make it work with the Fish shell
// Homepage: https://github.com/JeanMertz/chruby-fish

import (
	"fmt"
	
	"os/exec"
)

func installChrubyFish() {
	// Método 1: Descargar y extraer .tar.gz
	chrubyfish_tar_url := "https://github.com/JeanMertz/chruby-fish/archive/refs/tags/v1.0.0.tar.gz"
	chrubyfish_cmd_tar := exec.Command("curl", "-L", chrubyfish_tar_url, "-o", "package.tar.gz")
	err := chrubyfish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chrubyfish_zip_url := "https://github.com/JeanMertz/chruby-fish/archive/refs/tags/v1.0.0.zip"
	chrubyfish_cmd_zip := exec.Command("curl", "-L", chrubyfish_zip_url, "-o", "package.zip")
	err = chrubyfish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chrubyfish_bin_url := "https://github.com/JeanMertz/chruby-fish/archive/refs/tags/v1.0.0.bin"
	chrubyfish_cmd_bin := exec.Command("curl", "-L", chrubyfish_bin_url, "-o", "binary.bin")
	err = chrubyfish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chrubyfish_src_url := "https://github.com/JeanMertz/chruby-fish/archive/refs/tags/v1.0.0.src.tar.gz"
	chrubyfish_cmd_src := exec.Command("curl", "-L", chrubyfish_src_url, "-o", "source.tar.gz")
	err = chrubyfish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chrubyfish_cmd_direct := exec.Command("./binary")
	err = chrubyfish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: chruby")
	exec.Command("latte", "install", "chruby").Run()
	fmt.Println("Instalando dependencia: fish")
	exec.Command("latte", "install", "fish").Run()
}
