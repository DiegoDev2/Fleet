package main

// TerraformLs - Terraform Language Server
// Homepage: https://github.com/hashicorp/terraform-ls

import (
	"fmt"
	
	"os/exec"
)

func installTerraformLs() {
	// Método 1: Descargar y extraer .tar.gz
	terraformls_tar_url := "https://github.com/hashicorp/terraform-ls/archive/refs/tags/v0.34.3.tar.gz"
	terraformls_cmd_tar := exec.Command("curl", "-L", terraformls_tar_url, "-o", "package.tar.gz")
	err := terraformls_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformls_zip_url := "https://github.com/hashicorp/terraform-ls/archive/refs/tags/v0.34.3.zip"
	terraformls_cmd_zip := exec.Command("curl", "-L", terraformls_zip_url, "-o", "package.zip")
	err = terraformls_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformls_bin_url := "https://github.com/hashicorp/terraform-ls/archive/refs/tags/v0.34.3.bin"
	terraformls_cmd_bin := exec.Command("curl", "-L", terraformls_bin_url, "-o", "binary.bin")
	err = terraformls_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformls_src_url := "https://github.com/hashicorp/terraform-ls/archive/refs/tags/v0.34.3.src.tar.gz"
	terraformls_cmd_src := exec.Command("curl", "-L", terraformls_src_url, "-o", "source.tar.gz")
	err = terraformls_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformls_cmd_direct := exec.Command("./binary")
	err = terraformls_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
