package main

// Tflint - Linter for Terraform files
// Homepage: https://github.com/terraform-linters/tflint

import (
	"fmt"
	
	"os/exec"
)

func installTflint() {
	// Método 1: Descargar y extraer .tar.gz
	tflint_tar_url := "https://github.com/terraform-linters/tflint/archive/refs/tags/v0.53.0.tar.gz"
	tflint_cmd_tar := exec.Command("curl", "-L", tflint_tar_url, "-o", "package.tar.gz")
	err := tflint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tflint_zip_url := "https://github.com/terraform-linters/tflint/archive/refs/tags/v0.53.0.zip"
	tflint_cmd_zip := exec.Command("curl", "-L", tflint_zip_url, "-o", "package.zip")
	err = tflint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tflint_bin_url := "https://github.com/terraform-linters/tflint/archive/refs/tags/v0.53.0.bin"
	tflint_cmd_bin := exec.Command("curl", "-L", tflint_bin_url, "-o", "binary.bin")
	err = tflint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tflint_src_url := "https://github.com/terraform-linters/tflint/archive/refs/tags/v0.53.0.src.tar.gz"
	tflint_cmd_src := exec.Command("curl", "-L", tflint_src_url, "-o", "source.tar.gz")
	err = tflint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tflint_cmd_direct := exec.Command("./binary")
	err = tflint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
