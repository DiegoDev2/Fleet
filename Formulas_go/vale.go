package main

// Vale - Syntax-aware linter for prose
// Homepage: https://vale.sh/

import (
	"fmt"
	
	"os/exec"
)

func installVale() {
	// Método 1: Descargar y extraer .tar.gz
	vale_tar_url := "https://github.com/errata-ai/vale/archive/refs/tags/v3.7.1.tar.gz"
	vale_cmd_tar := exec.Command("curl", "-L", vale_tar_url, "-o", "package.tar.gz")
	err := vale_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vale_zip_url := "https://github.com/errata-ai/vale/archive/refs/tags/v3.7.1.zip"
	vale_cmd_zip := exec.Command("curl", "-L", vale_zip_url, "-o", "package.zip")
	err = vale_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vale_bin_url := "https://github.com/errata-ai/vale/archive/refs/tags/v3.7.1.bin"
	vale_cmd_bin := exec.Command("curl", "-L", vale_bin_url, "-o", "binary.bin")
	err = vale_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vale_src_url := "https://github.com/errata-ai/vale/archive/refs/tags/v3.7.1.src.tar.gz"
	vale_cmd_src := exec.Command("curl", "-L", vale_src_url, "-o", "source.tar.gz")
	err = vale_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vale_cmd_direct := exec.Command("./binary")
	err = vale_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
