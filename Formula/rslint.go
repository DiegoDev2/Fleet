package main

// Rslint - Extremely fast JavaScript and TypeScript linter
// Homepage: https://rslint.org/

import (
	"fmt"
	
	"os/exec"
)

func installRslint() {
	// Método 1: Descargar y extraer .tar.gz
	rslint_tar_url := "https://github.com/rslint/rslint/archive/refs/tags/v0.3.2.tar.gz"
	rslint_cmd_tar := exec.Command("curl", "-L", rslint_tar_url, "-o", "package.tar.gz")
	err := rslint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rslint_zip_url := "https://github.com/rslint/rslint/archive/refs/tags/v0.3.2.zip"
	rslint_cmd_zip := exec.Command("curl", "-L", rslint_zip_url, "-o", "package.zip")
	err = rslint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rslint_bin_url := "https://github.com/rslint/rslint/archive/refs/tags/v0.3.2.bin"
	rslint_cmd_bin := exec.Command("curl", "-L", rslint_bin_url, "-o", "binary.bin")
	err = rslint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rslint_src_url := "https://github.com/rslint/rslint/archive/refs/tags/v0.3.2.src.tar.gz"
	rslint_cmd_src := exec.Command("curl", "-L", rslint_src_url, "-o", "source.tar.gz")
	err = rslint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rslint_cmd_direct := exec.Command("./binary")
	err = rslint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
