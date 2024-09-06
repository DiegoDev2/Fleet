package main

// TerraformLandscape - Improve Terraform's plan output
// Homepage: https://github.com/coinbase/terraform-landscape

import (
	"fmt"
	
	"os/exec"
)

func installTerraformLandscape() {
	// Método 1: Descargar y extraer .tar.gz
	terraformlandscape_tar_url := "https://github.com/coinbase/terraform-landscape/archive/refs/tags/v0.3.4.tar.gz"
	terraformlandscape_cmd_tar := exec.Command("curl", "-L", terraformlandscape_tar_url, "-o", "package.tar.gz")
	err := terraformlandscape_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformlandscape_zip_url := "https://github.com/coinbase/terraform-landscape/archive/refs/tags/v0.3.4.zip"
	terraformlandscape_cmd_zip := exec.Command("curl", "-L", terraformlandscape_zip_url, "-o", "package.zip")
	err = terraformlandscape_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformlandscape_bin_url := "https://github.com/coinbase/terraform-landscape/archive/refs/tags/v0.3.4.bin"
	terraformlandscape_cmd_bin := exec.Command("curl", "-L", terraformlandscape_bin_url, "-o", "binary.bin")
	err = terraformlandscape_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformlandscape_src_url := "https://github.com/coinbase/terraform-landscape/archive/refs/tags/v0.3.4.src.tar.gz"
	terraformlandscape_cmd_src := exec.Command("curl", "-L", terraformlandscape_src_url, "-o", "source.tar.gz")
	err = terraformlandscape_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformlandscape_cmd_direct := exec.Command("./binary")
	err = terraformlandscape_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
	exec.Command("latte", "install", "ruby").Run()
}
