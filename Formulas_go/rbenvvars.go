package main

// RbenvVars - Safely sets global and per-project environment variables
// Homepage: https://github.com/rbenv/rbenv-vars

import (
	"fmt"
	
	"os/exec"
)

func installRbenvVars() {
	// Método 1: Descargar y extraer .tar.gz
	rbenvvars_tar_url := "https://github.com/rbenv/rbenv-vars/archive/refs/tags/v1.2.0.tar.gz"
	rbenvvars_cmd_tar := exec.Command("curl", "-L", rbenvvars_tar_url, "-o", "package.tar.gz")
	err := rbenvvars_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rbenvvars_zip_url := "https://github.com/rbenv/rbenv-vars/archive/refs/tags/v1.2.0.zip"
	rbenvvars_cmd_zip := exec.Command("curl", "-L", rbenvvars_zip_url, "-o", "package.zip")
	err = rbenvvars_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rbenvvars_bin_url := "https://github.com/rbenv/rbenv-vars/archive/refs/tags/v1.2.0.bin"
	rbenvvars_cmd_bin := exec.Command("curl", "-L", rbenvvars_bin_url, "-o", "binary.bin")
	err = rbenvvars_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rbenvvars_src_url := "https://github.com/rbenv/rbenv-vars/archive/refs/tags/v1.2.0.src.tar.gz"
	rbenvvars_cmd_src := exec.Command("curl", "-L", rbenvvars_src_url, "-o", "source.tar.gz")
	err = rbenvvars_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rbenvvars_cmd_direct := exec.Command("./binary")
	err = rbenvvars_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rbenv")
exec.Command("latte", "install", "rbenv")
}
