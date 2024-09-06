package main

// OhMyPosh - Prompt theme engine for any shell
// Homepage: https://ohmyposh.dev

import (
	"fmt"
	
	"os/exec"
)

func installOhMyPosh() {
	// Método 1: Descargar y extraer .tar.gz
	ohmyposh_tar_url := "https://github.com/JanDeDobbeleer/oh-my-posh/archive/refs/tags/v23.10.1.tar.gz"
	ohmyposh_cmd_tar := exec.Command("curl", "-L", ohmyposh_tar_url, "-o", "package.tar.gz")
	err := ohmyposh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ohmyposh_zip_url := "https://github.com/JanDeDobbeleer/oh-my-posh/archive/refs/tags/v23.10.1.zip"
	ohmyposh_cmd_zip := exec.Command("curl", "-L", ohmyposh_zip_url, "-o", "package.zip")
	err = ohmyposh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ohmyposh_bin_url := "https://github.com/JanDeDobbeleer/oh-my-posh/archive/refs/tags/v23.10.1.bin"
	ohmyposh_cmd_bin := exec.Command("curl", "-L", ohmyposh_bin_url, "-o", "binary.bin")
	err = ohmyposh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ohmyposh_src_url := "https://github.com/JanDeDobbeleer/oh-my-posh/archive/refs/tags/v23.10.1.src.tar.gz"
	ohmyposh_cmd_src := exec.Command("curl", "-L", ohmyposh_src_url, "-o", "source.tar.gz")
	err = ohmyposh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ohmyposh_cmd_direct := exec.Command("./binary")
	err = ohmyposh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
