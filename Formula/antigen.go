package main

// Antigen - Plugin manager for zsh, inspired by oh-my-zsh and vundle
// Homepage: https://antigen.sharats.me/

import (
	"fmt"
	
	"os/exec"
)

func installAntigen() {
	// Método 1: Descargar y extraer .tar.gz
	antigen_tar_url := "https://github.com/zsh-users/antigen/releases/download/v2.2.3/v2.2.3.tar.gz"
	antigen_cmd_tar := exec.Command("curl", "-L", antigen_tar_url, "-o", "package.tar.gz")
	err := antigen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	antigen_zip_url := "https://github.com/zsh-users/antigen/releases/download/v2.2.3/v2.2.3.zip"
	antigen_cmd_zip := exec.Command("curl", "-L", antigen_zip_url, "-o", "package.zip")
	err = antigen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	antigen_bin_url := "https://github.com/zsh-users/antigen/releases/download/v2.2.3/v2.2.3.bin"
	antigen_cmd_bin := exec.Command("curl", "-L", antigen_bin_url, "-o", "binary.bin")
	err = antigen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	antigen_src_url := "https://github.com/zsh-users/antigen/releases/download/v2.2.3/v2.2.3.src.tar.gz"
	antigen_cmd_src := exec.Command("curl", "-L", antigen_src_url, "-o", "source.tar.gz")
	err = antigen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	antigen_cmd_direct := exec.Command("./binary")
	err = antigen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
