package main

// RbenvBinstubs - Make rbenv aware of bundler binstubs
// Homepage: https://github.com/Purple-Devs/rbenv-binstubs

import (
	"fmt"
	
	"os/exec"
)

func installRbenvBinstubs() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvbinstubs_tar_url := "https://github.com/Purple-Devs/rbenv-binstubs/archive/refs/tags/v1.5.tar.gz"
	rbenvbinstubs_cmd_tar := exec.Command("curl", "-L", rbenvbinstubs_tar_url, "-o", "package.tar.gz")
	err := rbenvbinstubs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvbinstubs_zip_url := "https://github.com/Purple-Devs/rbenv-binstubs/archive/refs/tags/v1.5.zip"
	rbenvbinstubs_cmd_zip := exec.Command("curl", "-L", rbenvbinstubs_zip_url, "-o", "package.zip")
	err = rbenvbinstubs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvbinstubs_bin_url := "https://github.com/Purple-Devs/rbenv-binstubs/archive/refs/tags/v1.5.bin"
	rbenvbinstubs_cmd_bin := exec.Command("curl", "-L", rbenvbinstubs_bin_url, "-o", "binary.bin")
	err = rbenvbinstubs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvbinstubs_src_url := "https://github.com/Purple-Devs/rbenv-binstubs/archive/refs/tags/v1.5.src.tar.gz"
	rbenvbinstubs_cmd_src := exec.Command("curl", "-L", rbenvbinstubs_src_url, "-o", "source.tar.gz")
	err = rbenvbinstubs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvbinstubs_cmd_direct := exec.Command("./binary")
	err = rbenvbinstubs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
exec.Command("latte", "install", "rbenv")
}
