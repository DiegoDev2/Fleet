package main

// GitDelta - Syntax-highlighting pager for git and diff output
// Homepage: https://github.com/dandavison/delta

import (
	"fmt"
	
	"os/exec"
)

func installGitDelta() {
	// Método 1: Descargar y extraer .tar.gz
	gitdelta_tar_url := "https://github.com/dandavison/delta/archive/refs/tags/0.18.1.tar.gz"
	gitdelta_cmd_tar := exec.Command("curl", "-L", gitdelta_tar_url, "-o", "package.tar.gz")
	err := gitdelta_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitdelta_zip_url := "https://github.com/dandavison/delta/archive/refs/tags/0.18.1.zip"
	gitdelta_cmd_zip := exec.Command("curl", "-L", gitdelta_zip_url, "-o", "package.zip")
	err = gitdelta_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitdelta_bin_url := "https://github.com/dandavison/delta/archive/refs/tags/0.18.1.bin"
	gitdelta_cmd_bin := exec.Command("curl", "-L", gitdelta_bin_url, "-o", "binary.bin")
	err = gitdelta_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitdelta_src_url := "https://github.com/dandavison/delta/archive/refs/tags/0.18.1.src.tar.gz"
	gitdelta_cmd_src := exec.Command("curl", "-L", gitdelta_src_url, "-o", "source.tar.gz")
	err = gitdelta_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitdelta_cmd_direct := exec.Command("./binary")
	err = gitdelta_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
