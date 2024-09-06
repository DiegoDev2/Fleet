package main

// FaasCli - CLI for templating and/or deploying FaaS functions
// Homepage: https://www.openfaas.com/

import (
	"fmt"
	
	"os/exec"
)

func installFaasCli() {
	// Método 1: Descargar y extraer .tar.gz
	faascli_tar_url := "https://github.com/openfaas/faas-cli/archive/refs/tags/0.16.34.tar.gz"
	faascli_cmd_tar := exec.Command("curl", "-L", faascli_tar_url, "-o", "package.tar.gz")
	err := faascli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faascli_zip_url := "https://github.com/openfaas/faas-cli/archive/refs/tags/0.16.34.zip"
	faascli_cmd_zip := exec.Command("curl", "-L", faascli_zip_url, "-o", "package.zip")
	err = faascli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faascli_bin_url := "https://github.com/openfaas/faas-cli/archive/refs/tags/0.16.34.bin"
	faascli_cmd_bin := exec.Command("curl", "-L", faascli_bin_url, "-o", "binary.bin")
	err = faascli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faascli_src_url := "https://github.com/openfaas/faas-cli/archive/refs/tags/0.16.34.src.tar.gz"
	faascli_cmd_src := exec.Command("curl", "-L", faascli_src_url, "-o", "source.tar.gz")
	err = faascli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faascli_cmd_direct := exec.Command("./binary")
	err = faascli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
