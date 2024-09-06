package main

// Eslint - AST-based pattern checker for JavaScript
// Homepage: https://eslint.org

import (
	"fmt"
	
	"os/exec"
)

func installEslint() {
	// Método 1: Descargar y extraer .tar.gz
	eslint_tar_url := "https://registry.npmjs.org/eslint/-/eslint-9.9.1.tgz"
	eslint_cmd_tar := exec.Command("curl", "-L", eslint_tar_url, "-o", "package.tar.gz")
	err := eslint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	eslint_zip_url := "https://registry.npmjs.org/eslint/-/eslint-9.9.1.tgz"
	eslint_cmd_zip := exec.Command("curl", "-L", eslint_zip_url, "-o", "package.zip")
	err = eslint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	eslint_bin_url := "https://registry.npmjs.org/eslint/-/eslint-9.9.1.tgz"
	eslint_cmd_bin := exec.Command("curl", "-L", eslint_bin_url, "-o", "binary.bin")
	err = eslint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	eslint_src_url := "https://registry.npmjs.org/eslint/-/eslint-9.9.1.tgz"
	eslint_cmd_src := exec.Command("curl", "-L", eslint_src_url, "-o", "source.tar.gz")
	err = eslint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	eslint_cmd_direct := exec.Command("./binary")
	err = eslint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
