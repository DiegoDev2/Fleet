package main

// RbenvAliases - Make aliases for Ruby versions
// Homepage: https://github.com/tpope/rbenv-aliases

import (
	"fmt"
	
	"os/exec"
)

func installRbenvAliases() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvaliases_tar_url := "https://github.com/tpope/rbenv-aliases/archive/refs/tags/v1.1.0.tar.gz"
	rbenvaliases_cmd_tar := exec.Command("curl", "-L", rbenvaliases_tar_url, "-o", "package.tar.gz")
	err := rbenvaliases_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvaliases_zip_url := "https://github.com/tpope/rbenv-aliases/archive/refs/tags/v1.1.0.zip"
	rbenvaliases_cmd_zip := exec.Command("curl", "-L", rbenvaliases_zip_url, "-o", "package.zip")
	err = rbenvaliases_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvaliases_bin_url := "https://github.com/tpope/rbenv-aliases/archive/refs/tags/v1.1.0.bin"
	rbenvaliases_cmd_bin := exec.Command("curl", "-L", rbenvaliases_bin_url, "-o", "binary.bin")
	err = rbenvaliases_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvaliases_src_url := "https://github.com/tpope/rbenv-aliases/archive/refs/tags/v1.1.0.src.tar.gz"
	rbenvaliases_cmd_src := exec.Command("curl", "-L", rbenvaliases_src_url, "-o", "source.tar.gz")
	err = rbenvaliases_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvaliases_cmd_direct := exec.Command("./binary")
	err = rbenvaliases_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
	exec.Command("latte", "install", "rbenv").Run()
}
