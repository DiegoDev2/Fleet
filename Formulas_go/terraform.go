package main

// Terraform - Tool to build, change, and version infrastructure
// Homepage: https://www.terraform.io/

import (
	"fmt"
	
	"os/exec"
)

func installTerraform() {
	// Método 1: Descargar y extraer .tar.gz
	terraform_tar_url := "https://github.com/hashicorp/terraform/archive/refs/tags/v1.5.7.tar.gz"
	terraform_cmd_tar := exec.Command("curl", "-L", terraform_tar_url, "-o", "package.tar.gz")
	err := terraform_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraform_zip_url := "https://github.com/hashicorp/terraform/archive/refs/tags/v1.5.7.zip"
	terraform_cmd_zip := exec.Command("curl", "-L", terraform_zip_url, "-o", "package.zip")
	err = terraform_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraform_bin_url := "https://github.com/hashicorp/terraform/archive/refs/tags/v1.5.7.bin"
	terraform_cmd_bin := exec.Command("curl", "-L", terraform_bin_url, "-o", "binary.bin")
	err = terraform_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraform_src_url := "https://github.com/hashicorp/terraform/archive/refs/tags/v1.5.7.src.tar.gz"
	terraform_cmd_src := exec.Command("curl", "-L", terraform_src_url, "-o", "source.tar.gz")
	err = terraform_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraform_cmd_direct := exec.Command("./binary")
	err = terraform_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
