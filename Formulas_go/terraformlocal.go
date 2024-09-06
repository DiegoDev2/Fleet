package main

// TerraformLocal - CLI wrapper to deploy your Terraform applications directly to LocalStack
// Homepage: https://localstack.cloud/

import (
	"fmt"
	
	"os/exec"
)

func installTerraformLocal() {
	// Método 1: Descargar y extraer .tar.gz
	terraformlocal_tar_url := "https://files.pythonhosted.org/packages/88/80/099853d1614341cbde71938ddd86cf9ded44abf518a384e05d8f0407002c/terraform_local-0.19.0.tar.gz"
	terraformlocal_cmd_tar := exec.Command("curl", "-L", terraformlocal_tar_url, "-o", "package.tar.gz")
	err := terraformlocal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformlocal_zip_url := "https://files.pythonhosted.org/packages/88/80/099853d1614341cbde71938ddd86cf9ded44abf518a384e05d8f0407002c/terraform_local-0.19.0.zip"
	terraformlocal_cmd_zip := exec.Command("curl", "-L", terraformlocal_zip_url, "-o", "package.zip")
	err = terraformlocal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformlocal_bin_url := "https://files.pythonhosted.org/packages/88/80/099853d1614341cbde71938ddd86cf9ded44abf518a384e05d8f0407002c/terraform_local-0.19.0.bin"
	terraformlocal_cmd_bin := exec.Command("curl", "-L", terraformlocal_bin_url, "-o", "binary.bin")
	err = terraformlocal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformlocal_src_url := "https://files.pythonhosted.org/packages/88/80/099853d1614341cbde71938ddd86cf9ded44abf518a384e05d8f0407002c/terraform_local-0.19.0.src.tar.gz"
	terraformlocal_cmd_src := exec.Command("curl", "-L", terraformlocal_src_url, "-o", "source.tar.gz")
	err = terraformlocal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformlocal_cmd_direct := exec.Command("./binary")
	err = terraformlocal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: localstack")
exec.Command("latte", "install", "localstack")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
