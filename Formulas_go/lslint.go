package main

// LsLint - Extremely fast file and directory name linter
// Homepage: https://ls-lint.org/

import (
	"fmt"
	
	"os/exec"
)

func installLsLint() {
	// Método 1: Descargar y extraer .tar.gz
	lslint_tar_url := "https://github.com/loeffel-io/ls-lint/archive/refs/tags/v2.2.3.tar.gz"
	lslint_cmd_tar := exec.Command("curl", "-L", lslint_tar_url, "-o", "package.tar.gz")
	err := lslint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lslint_zip_url := "https://github.com/loeffel-io/ls-lint/archive/refs/tags/v2.2.3.zip"
	lslint_cmd_zip := exec.Command("curl", "-L", lslint_zip_url, "-o", "package.zip")
	err = lslint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lslint_bin_url := "https://github.com/loeffel-io/ls-lint/archive/refs/tags/v2.2.3.bin"
	lslint_cmd_bin := exec.Command("curl", "-L", lslint_bin_url, "-o", "binary.bin")
	err = lslint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lslint_src_url := "https://github.com/loeffel-io/ls-lint/archive/refs/tags/v2.2.3.src.tar.gz"
	lslint_cmd_src := exec.Command("curl", "-L", lslint_src_url, "-o", "source.tar.gz")
	err = lslint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lslint_cmd_direct := exec.Command("./binary")
	err = lslint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
