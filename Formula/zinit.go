package main

// Zinit - Flexible and fast Zsh plugin manager
// Homepage: https://zdharma-continuum.github.io/zinit/wiki/

import (
	"fmt"
	
	"os/exec"
)

func installZinit() {
	// Método 1: Descargar y extraer .tar.gz
	zinit_tar_url := "https://github.com/zdharma-continuum/zinit/archive/refs/tags/v3.13.1.tar.gz"
	zinit_cmd_tar := exec.Command("curl", "-L", zinit_tar_url, "-o", "package.tar.gz")
	err := zinit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zinit_zip_url := "https://github.com/zdharma-continuum/zinit/archive/refs/tags/v3.13.1.zip"
	zinit_cmd_zip := exec.Command("curl", "-L", zinit_zip_url, "-o", "package.zip")
	err = zinit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zinit_bin_url := "https://github.com/zdharma-continuum/zinit/archive/refs/tags/v3.13.1.bin"
	zinit_cmd_bin := exec.Command("curl", "-L", zinit_bin_url, "-o", "binary.bin")
	err = zinit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zinit_src_url := "https://github.com/zdharma-continuum/zinit/archive/refs/tags/v3.13.1.src.tar.gz"
	zinit_cmd_src := exec.Command("curl", "-L", zinit_src_url, "-o", "source.tar.gz")
	err = zinit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zinit_cmd_direct := exec.Command("./binary")
	err = zinit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
