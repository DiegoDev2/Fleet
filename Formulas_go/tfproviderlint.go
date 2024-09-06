package main

// Tfproviderlint - Terraform Provider Lint Tool
// Homepage: https://github.com/bflad/tfproviderlint

import (
	"fmt"
	
	"os/exec"
)

func installTfproviderlint() {
	// Método 1: Descargar y extraer .tar.gz
	tfproviderlint_tar_url := "https://github.com/bflad/tfproviderlint/archive/refs/tags/v0.30.0.tar.gz"
	tfproviderlint_cmd_tar := exec.Command("curl", "-L", tfproviderlint_tar_url, "-o", "package.tar.gz")
	err := tfproviderlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfproviderlint_zip_url := "https://github.com/bflad/tfproviderlint/archive/refs/tags/v0.30.0.zip"
	tfproviderlint_cmd_zip := exec.Command("curl", "-L", tfproviderlint_zip_url, "-o", "package.zip")
	err = tfproviderlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfproviderlint_bin_url := "https://github.com/bflad/tfproviderlint/archive/refs/tags/v0.30.0.bin"
	tfproviderlint_cmd_bin := exec.Command("curl", "-L", tfproviderlint_bin_url, "-o", "binary.bin")
	err = tfproviderlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfproviderlint_src_url := "https://github.com/bflad/tfproviderlint/archive/refs/tags/v0.30.0.src.tar.gz"
	tfproviderlint_cmd_src := exec.Command("curl", "-L", tfproviderlint_src_url, "-o", "source.tar.gz")
	err = tfproviderlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfproviderlint_cmd_direct := exec.Command("./binary")
	err = tfproviderlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
