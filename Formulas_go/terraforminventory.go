package main

// TerraformInventory - Go app which generates a dynamic Ansible inventory from a Terraform state file
// Homepage: https://github.com/adammck/terraform-inventory

import (
	"fmt"
	
	"os/exec"
)

func installTerraformInventory() {
	// Método 1: Descargar y extraer .tar.gz
	terraforminventory_tar_url := "https://github.com/adammck/terraform-inventory/archive/refs/tags/v0.10.tar.gz"
	terraforminventory_cmd_tar := exec.Command("curl", "-L", terraforminventory_tar_url, "-o", "package.tar.gz")
	err := terraforminventory_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraforminventory_zip_url := "https://github.com/adammck/terraform-inventory/archive/refs/tags/v0.10.zip"
	terraforminventory_cmd_zip := exec.Command("curl", "-L", terraforminventory_zip_url, "-o", "package.zip")
	err = terraforminventory_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraforminventory_bin_url := "https://github.com/adammck/terraform-inventory/archive/refs/tags/v0.10.bin"
	terraforminventory_cmd_bin := exec.Command("curl", "-L", terraforminventory_bin_url, "-o", "binary.bin")
	err = terraforminventory_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraforminventory_src_url := "https://github.com/adammck/terraform-inventory/archive/refs/tags/v0.10.src.tar.gz"
	terraforminventory_cmd_src := exec.Command("curl", "-L", terraforminventory_src_url, "-o", "source.tar.gz")
	err = terraforminventory_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraforminventory_cmd_direct := exec.Command("./binary")
	err = terraforminventory_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
