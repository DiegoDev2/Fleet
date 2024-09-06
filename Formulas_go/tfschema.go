package main

// Tfschema - Schema inspector for Terraform/OpenTofu providers
// Homepage: https://github.com/minamijoyo/tfschema

import (
	"fmt"
	
	"os/exec"
)

func installTfschema() {
	// Método 1: Descargar y extraer .tar.gz
	tfschema_tar_url := "https://github.com/minamijoyo/tfschema/archive/refs/tags/v0.7.9.tar.gz"
	tfschema_cmd_tar := exec.Command("curl", "-L", tfschema_tar_url, "-o", "package.tar.gz")
	err := tfschema_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfschema_zip_url := "https://github.com/minamijoyo/tfschema/archive/refs/tags/v0.7.9.zip"
	tfschema_cmd_zip := exec.Command("curl", "-L", tfschema_zip_url, "-o", "package.zip")
	err = tfschema_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfschema_bin_url := "https://github.com/minamijoyo/tfschema/archive/refs/tags/v0.7.9.bin"
	tfschema_cmd_bin := exec.Command("curl", "-L", tfschema_bin_url, "-o", "binary.bin")
	err = tfschema_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfschema_src_url := "https://github.com/minamijoyo/tfschema/archive/refs/tags/v0.7.9.src.tar.gz"
	tfschema_cmd_src := exec.Command("curl", "-L", tfschema_src_url, "-o", "source.tar.gz")
	err = tfschema_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfschema_cmd_direct := exec.Command("./binary")
	err = tfschema_cmd_direct.Run()
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
