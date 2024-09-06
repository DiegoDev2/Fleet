package main

// Pint - Prometheus rule linter/validator
// Homepage: https://cloudflare.github.io/pint/

import (
	"fmt"
	
	"os/exec"
)

func installPint() {
	// Método 1: Descargar y extraer .tar.gz
	pint_tar_url := "https://github.com/cloudflare/pint/archive/refs/tags/v0.65.1.tar.gz"
	pint_cmd_tar := exec.Command("curl", "-L", pint_tar_url, "-o", "package.tar.gz")
	err := pint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pint_zip_url := "https://github.com/cloudflare/pint/archive/refs/tags/v0.65.1.zip"
	pint_cmd_zip := exec.Command("curl", "-L", pint_zip_url, "-o", "package.zip")
	err = pint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pint_bin_url := "https://github.com/cloudflare/pint/archive/refs/tags/v0.65.1.bin"
	pint_cmd_bin := exec.Command("curl", "-L", pint_bin_url, "-o", "binary.bin")
	err = pint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pint_src_url := "https://github.com/cloudflare/pint/archive/refs/tags/v0.65.1.src.tar.gz"
	pint_cmd_src := exec.Command("curl", "-L", pint_src_url, "-o", "source.tar.gz")
	err = pint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pint_cmd_direct := exec.Command("./binary")
	err = pint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
