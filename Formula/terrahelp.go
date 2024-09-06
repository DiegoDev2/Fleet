package main

// Terrahelp - Tool providing extra functionality for Terraform
// Homepage: https://github.com/opencredo/terrahelp

import (
	"fmt"
	
	"os/exec"
)

func installTerrahelp() {
	// Método 1: Descargar y extraer .tar.gz
	terrahelp_tar_url := "https://github.com/opencredo/terrahelp/archive/refs/tags/v0.7.5.tar.gz"
	terrahelp_cmd_tar := exec.Command("curl", "-L", terrahelp_tar_url, "-o", "package.tar.gz")
	err := terrahelp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terrahelp_zip_url := "https://github.com/opencredo/terrahelp/archive/refs/tags/v0.7.5.zip"
	terrahelp_cmd_zip := exec.Command("curl", "-L", terrahelp_zip_url, "-o", "package.zip")
	err = terrahelp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terrahelp_bin_url := "https://github.com/opencredo/terrahelp/archive/refs/tags/v0.7.5.bin"
	terrahelp_cmd_bin := exec.Command("curl", "-L", terrahelp_bin_url, "-o", "binary.bin")
	err = terrahelp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terrahelp_src_url := "https://github.com/opencredo/terrahelp/archive/refs/tags/v0.7.5.src.tar.gz"
	terrahelp_cmd_src := exec.Command("curl", "-L", terrahelp_src_url, "-o", "source.tar.gz")
	err = terrahelp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terrahelp_cmd_direct := exec.Command("./binary")
	err = terrahelp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
