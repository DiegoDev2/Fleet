package main

// Tfmigrate - Terraform/OpenTofu state migration tool for GitOps
// Homepage: https://github.com/minamijoyo/tfmigrate

import (
	"fmt"
	
	"os/exec"
)

func installTfmigrate() {
	// Método 1: Descargar y extraer .tar.gz
	tfmigrate_tar_url := "https://github.com/minamijoyo/tfmigrate/archive/refs/tags/v0.3.24.tar.gz"
	tfmigrate_cmd_tar := exec.Command("curl", "-L", tfmigrate_tar_url, "-o", "package.tar.gz")
	err := tfmigrate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfmigrate_zip_url := "https://github.com/minamijoyo/tfmigrate/archive/refs/tags/v0.3.24.zip"
	tfmigrate_cmd_zip := exec.Command("curl", "-L", tfmigrate_zip_url, "-o", "package.zip")
	err = tfmigrate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfmigrate_bin_url := "https://github.com/minamijoyo/tfmigrate/archive/refs/tags/v0.3.24.bin"
	tfmigrate_cmd_bin := exec.Command("curl", "-L", tfmigrate_bin_url, "-o", "binary.bin")
	err = tfmigrate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfmigrate_src_url := "https://github.com/minamijoyo/tfmigrate/archive/refs/tags/v0.3.24.src.tar.gz"
	tfmigrate_cmd_src := exec.Command("curl", "-L", tfmigrate_src_url, "-o", "source.tar.gz")
	err = tfmigrate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfmigrate_cmd_direct := exec.Command("./binary")
	err = tfmigrate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: opentofu")
exec.Command("latte", "install", "opentofu")
}
