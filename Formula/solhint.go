package main

// Solhint - Linter for Solidity code
// Homepage: https://protofire.github.io/solhint/

import (
	"fmt"
	
	"os/exec"
)

func installSolhint() {
	// Método 1: Descargar y extraer .tar.gz
	solhint_tar_url := "https://registry.npmjs.org/solhint/-/solhint-5.0.3.tgz"
	solhint_cmd_tar := exec.Command("curl", "-L", solhint_tar_url, "-o", "package.tar.gz")
	err := solhint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solhint_zip_url := "https://registry.npmjs.org/solhint/-/solhint-5.0.3.tgz"
	solhint_cmd_zip := exec.Command("curl", "-L", solhint_zip_url, "-o", "package.zip")
	err = solhint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solhint_bin_url := "https://registry.npmjs.org/solhint/-/solhint-5.0.3.tgz"
	solhint_cmd_bin := exec.Command("curl", "-L", solhint_bin_url, "-o", "binary.bin")
	err = solhint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solhint_src_url := "https://registry.npmjs.org/solhint/-/solhint-5.0.3.tgz"
	solhint_cmd_src := exec.Command("curl", "-L", solhint_src_url, "-o", "source.tar.gz")
	err = solhint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solhint_cmd_direct := exec.Command("./binary")
	err = solhint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
