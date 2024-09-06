package main

// Stylelint - Modern CSS linter
// Homepage: https://stylelint.io/

import (
	"fmt"
	
	"os/exec"
)

func installStylelint() {
	// Método 1: Descargar y extraer .tar.gz
	stylelint_tar_url := "https://registry.npmjs.org/stylelint/-/stylelint-16.9.0.tgz"
	stylelint_cmd_tar := exec.Command("curl", "-L", stylelint_tar_url, "-o", "package.tar.gz")
	err := stylelint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stylelint_zip_url := "https://registry.npmjs.org/stylelint/-/stylelint-16.9.0.tgz"
	stylelint_cmd_zip := exec.Command("curl", "-L", stylelint_zip_url, "-o", "package.zip")
	err = stylelint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stylelint_bin_url := "https://registry.npmjs.org/stylelint/-/stylelint-16.9.0.tgz"
	stylelint_cmd_bin := exec.Command("curl", "-L", stylelint_bin_url, "-o", "binary.bin")
	err = stylelint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stylelint_src_url := "https://registry.npmjs.org/stylelint/-/stylelint-16.9.0.tgz"
	stylelint_cmd_src := exec.Command("curl", "-L", stylelint_src_url, "-o", "source.tar.gz")
	err = stylelint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stylelint_cmd_direct := exec.Command("./binary")
	err = stylelint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
