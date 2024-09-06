package main

// TerraformDocs - Tool to generate documentation from Terraform modules
// Homepage: https://github.com/terraform-docs/terraform-docs

import (
	"fmt"
	
	"os/exec"
)

func installTerraformDocs() {
	// Método 1: Descargar y extraer .tar.gz
	terraformdocs_tar_url := "https://github.com/terraform-docs/terraform-docs/archive/refs/tags/v0.18.0.tar.gz"
	terraformdocs_cmd_tar := exec.Command("curl", "-L", terraformdocs_tar_url, "-o", "package.tar.gz")
	err := terraformdocs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformdocs_zip_url := "https://github.com/terraform-docs/terraform-docs/archive/refs/tags/v0.18.0.zip"
	terraformdocs_cmd_zip := exec.Command("curl", "-L", terraformdocs_zip_url, "-o", "package.zip")
	err = terraformdocs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformdocs_bin_url := "https://github.com/terraform-docs/terraform-docs/archive/refs/tags/v0.18.0.bin"
	terraformdocs_cmd_bin := exec.Command("curl", "-L", terraformdocs_bin_url, "-o", "binary.bin")
	err = terraformdocs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformdocs_src_url := "https://github.com/terraform-docs/terraform-docs/archive/refs/tags/v0.18.0.src.tar.gz"
	terraformdocs_cmd_src := exec.Command("curl", "-L", terraformdocs_src_url, "-o", "source.tar.gz")
	err = terraformdocs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformdocs_cmd_direct := exec.Command("./binary")
	err = terraformdocs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
