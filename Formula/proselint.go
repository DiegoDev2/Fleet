package main

// Proselint - Linter for prose
// Homepage: https://github.com/amperser/proselint

import (
	"fmt"
	
	"os/exec"
)

func installProselint() {
	// Método 1: Descargar y extraer .tar.gz
	proselint_tar_url := "https://files.pythonhosted.org/packages/58/66/bc509b61df9a317689f6a87679f2f9f625f6f02dfb9d0e220bd41f121f07/proselint-0.14.0.tar.gz"
	proselint_cmd_tar := exec.Command("curl", "-L", proselint_tar_url, "-o", "package.tar.gz")
	err := proselint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proselint_zip_url := "https://files.pythonhosted.org/packages/58/66/bc509b61df9a317689f6a87679f2f9f625f6f02dfb9d0e220bd41f121f07/proselint-0.14.0.zip"
	proselint_cmd_zip := exec.Command("curl", "-L", proselint_zip_url, "-o", "package.zip")
	err = proselint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proselint_bin_url := "https://files.pythonhosted.org/packages/58/66/bc509b61df9a317689f6a87679f2f9f625f6f02dfb9d0e220bd41f121f07/proselint-0.14.0.bin"
	proselint_cmd_bin := exec.Command("curl", "-L", proselint_bin_url, "-o", "binary.bin")
	err = proselint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proselint_src_url := "https://files.pythonhosted.org/packages/58/66/bc509b61df9a317689f6a87679f2f9f625f6f02dfb9d0e220bd41f121f07/proselint-0.14.0.src.tar.gz"
	proselint_cmd_src := exec.Command("curl", "-L", proselint_src_url, "-o", "source.tar.gz")
	err = proselint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proselint_cmd_direct := exec.Command("./binary")
	err = proselint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
