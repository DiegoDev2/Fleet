package main

// Fzf - Command-line fuzzy finder written in Go
// Homepage: https://github.com/junegunn/fzf

import (
	"fmt"
	
	"os/exec"
)

func installFzf() {
	// Método 1: Descargar y extraer .tar.gz
	fzf_tar_url := "https://github.com/junegunn/fzf/archive/refs/tags/v0.55.0.tar.gz"
	fzf_cmd_tar := exec.Command("curl", "-L", fzf_tar_url, "-o", "package.tar.gz")
	err := fzf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fzf_zip_url := "https://github.com/junegunn/fzf/archive/refs/tags/v0.55.0.zip"
	fzf_cmd_zip := exec.Command("curl", "-L", fzf_zip_url, "-o", "package.zip")
	err = fzf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fzf_bin_url := "https://github.com/junegunn/fzf/archive/refs/tags/v0.55.0.bin"
	fzf_cmd_bin := exec.Command("curl", "-L", fzf_bin_url, "-o", "binary.bin")
	err = fzf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fzf_src_url := "https://github.com/junegunn/fzf/archive/refs/tags/v0.55.0.src.tar.gz"
	fzf_cmd_src := exec.Command("curl", "-L", fzf_src_url, "-o", "source.tar.gz")
	err = fzf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fzf_cmd_direct := exec.Command("./binary")
	err = fzf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
