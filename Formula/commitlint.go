package main

// Commitlint - Lint commit messages according to a commit convention
// Homepage: https://commitlint.js.org/#/

import (
	"fmt"
	
	"os/exec"
)

func installCommitlint() {
	// Método 1: Descargar y extraer .tar.gz
	commitlint_tar_url := "https://registry.npmjs.org/commitlint/-/commitlint-19.4.1.tgz"
	commitlint_cmd_tar := exec.Command("curl", "-L", commitlint_tar_url, "-o", "package.tar.gz")
	err := commitlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	commitlint_zip_url := "https://registry.npmjs.org/commitlint/-/commitlint-19.4.1.tgz"
	commitlint_cmd_zip := exec.Command("curl", "-L", commitlint_zip_url, "-o", "package.zip")
	err = commitlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	commitlint_bin_url := "https://registry.npmjs.org/commitlint/-/commitlint-19.4.1.tgz"
	commitlint_cmd_bin := exec.Command("curl", "-L", commitlint_bin_url, "-o", "binary.bin")
	err = commitlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	commitlint_src_url := "https://registry.npmjs.org/commitlint/-/commitlint-19.4.1.tgz"
	commitlint_cmd_src := exec.Command("curl", "-L", commitlint_src_url, "-o", "source.tar.gz")
	err = commitlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	commitlint_cmd_direct := exec.Command("./binary")
	err = commitlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
