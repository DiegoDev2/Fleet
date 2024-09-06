package main

// Akamai - CLI toolkit for working with Akamai's APIs
// Homepage: https://github.com/akamai/cli

import (
	"fmt"
	
	"os/exec"
)

func installAkamai() {
	// Método 1: Descargar y extraer .tar.gz
	akamai_tar_url := "https://github.com/akamai/cli/archive/refs/tags/v1.6.0.tar.gz"
	akamai_cmd_tar := exec.Command("curl", "-L", akamai_tar_url, "-o", "package.tar.gz")
	err := akamai_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	akamai_zip_url := "https://github.com/akamai/cli/archive/refs/tags/v1.6.0.zip"
	akamai_cmd_zip := exec.Command("curl", "-L", akamai_zip_url, "-o", "package.zip")
	err = akamai_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	akamai_bin_url := "https://github.com/akamai/cli/archive/refs/tags/v1.6.0.bin"
	akamai_cmd_bin := exec.Command("curl", "-L", akamai_bin_url, "-o", "binary.bin")
	err = akamai_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	akamai_src_url := "https://github.com/akamai/cli/archive/refs/tags/v1.6.0.src.tar.gz"
	akamai_cmd_src := exec.Command("curl", "-L", akamai_src_url, "-o", "source.tar.gz")
	err = akamai_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	akamai_cmd_direct := exec.Command("./binary")
	err = akamai_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
