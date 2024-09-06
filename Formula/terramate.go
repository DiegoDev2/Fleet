package main

// Terramate - Managing Terraform stacks with change detections and code generations
// Homepage: https://terramate.io/docs/cli/

import (
	"fmt"
	
	"os/exec"
)

func installTerramate() {
	// Método 1: Descargar y extraer .tar.gz
	terramate_tar_url := "https://github.com/terramate-io/terramate/archive/refs/tags/v0.10.2.tar.gz"
	terramate_cmd_tar := exec.Command("curl", "-L", terramate_tar_url, "-o", "package.tar.gz")
	err := terramate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terramate_zip_url := "https://github.com/terramate-io/terramate/archive/refs/tags/v0.10.2.zip"
	terramate_cmd_zip := exec.Command("curl", "-L", terramate_zip_url, "-o", "package.zip")
	err = terramate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terramate_bin_url := "https://github.com/terramate-io/terramate/archive/refs/tags/v0.10.2.bin"
	terramate_cmd_bin := exec.Command("curl", "-L", terramate_bin_url, "-o", "binary.bin")
	err = terramate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terramate_src_url := "https://github.com/terramate-io/terramate/archive/refs/tags/v0.10.2.src.tar.gz"
	terramate_cmd_src := exec.Command("curl", "-L", terramate_src_url, "-o", "source.tar.gz")
	err = terramate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terramate_cmd_direct := exec.Command("./binary")
	err = terramate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
