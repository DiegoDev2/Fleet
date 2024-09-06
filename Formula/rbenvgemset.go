package main

// RbenvGemset - Adds basic gemset support to rbenv
// Homepage: https://github.com/jf/rbenv-gemset

import (
	"fmt"
	
	"os/exec"
)

func installRbenvGemset() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvgemset_tar_url := "https://github.com/jf/rbenv-gemset/archive/refs/tags/v0.5.10.tar.gz"
	rbenvgemset_cmd_tar := exec.Command("curl", "-L", rbenvgemset_tar_url, "-o", "package.tar.gz")
	err := rbenvgemset_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvgemset_zip_url := "https://github.com/jf/rbenv-gemset/archive/refs/tags/v0.5.10.zip"
	rbenvgemset_cmd_zip := exec.Command("curl", "-L", rbenvgemset_zip_url, "-o", "package.zip")
	err = rbenvgemset_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvgemset_bin_url := "https://github.com/jf/rbenv-gemset/archive/refs/tags/v0.5.10.bin"
	rbenvgemset_cmd_bin := exec.Command("curl", "-L", rbenvgemset_bin_url, "-o", "binary.bin")
	err = rbenvgemset_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvgemset_src_url := "https://github.com/jf/rbenv-gemset/archive/refs/tags/v0.5.10.src.tar.gz"
	rbenvgemset_cmd_src := exec.Command("curl", "-L", rbenvgemset_src_url, "-o", "source.tar.gz")
	err = rbenvgemset_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvgemset_cmd_direct := exec.Command("./binary")
	err = rbenvgemset_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
	exec.Command("latte", "install", "rbenv").Run()
}
