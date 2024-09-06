package main

// Unibilium - Very basic terminfo library
// Homepage: https://github.com/neovim/unibilium

import (
	"fmt"
	
	"os/exec"
)

func installUnibilium() {
	// Método 1: Descargar y extraer .tar.gz
	unibilium_tar_url := "https://github.com/neovim/unibilium/archive/refs/tags/v2.1.1.tar.gz"
	unibilium_cmd_tar := exec.Command("curl", "-L", unibilium_tar_url, "-o", "package.tar.gz")
	err := unibilium_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unibilium_zip_url := "https://github.com/neovim/unibilium/archive/refs/tags/v2.1.1.zip"
	unibilium_cmd_zip := exec.Command("curl", "-L", unibilium_zip_url, "-o", "package.zip")
	err = unibilium_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unibilium_bin_url := "https://github.com/neovim/unibilium/archive/refs/tags/v2.1.1.bin"
	unibilium_cmd_bin := exec.Command("curl", "-L", unibilium_bin_url, "-o", "binary.bin")
	err = unibilium_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unibilium_src_url := "https://github.com/neovim/unibilium/archive/refs/tags/v2.1.1.src.tar.gz"
	unibilium_cmd_src := exec.Command("curl", "-L", unibilium_src_url, "-o", "source.tar.gz")
	err = unibilium_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unibilium_cmd_direct := exec.Command("./binary")
	err = unibilium_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
