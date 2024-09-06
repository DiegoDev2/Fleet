package main

// Terragrunt - Thin wrapper for Terraform e.g. for locking state
// Homepage: https://terragrunt.gruntwork.io/

import (
	"fmt"
	
	"os/exec"
)

func installTerragrunt() {
	// Método 1: Descargar y extraer .tar.gz
	terragrunt_tar_url := "https://github.com/gruntwork-io/terragrunt/archive/refs/tags/v0.67.3.tar.gz"
	terragrunt_cmd_tar := exec.Command("curl", "-L", terragrunt_tar_url, "-o", "package.tar.gz")
	err := terragrunt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terragrunt_zip_url := "https://github.com/gruntwork-io/terragrunt/archive/refs/tags/v0.67.3.zip"
	terragrunt_cmd_zip := exec.Command("curl", "-L", terragrunt_zip_url, "-o", "package.zip")
	err = terragrunt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terragrunt_bin_url := "https://github.com/gruntwork-io/terragrunt/archive/refs/tags/v0.67.3.bin"
	terragrunt_cmd_bin := exec.Command("curl", "-L", terragrunt_bin_url, "-o", "binary.bin")
	err = terragrunt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terragrunt_src_url := "https://github.com/gruntwork-io/terragrunt/archive/refs/tags/v0.67.3.src.tar.gz"
	terragrunt_cmd_src := exec.Command("curl", "-L", terragrunt_src_url, "-o", "source.tar.gz")
	err = terragrunt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terragrunt_cmd_direct := exec.Command("./binary")
	err = terragrunt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
