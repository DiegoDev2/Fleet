package main

// SnykCli - Scans and monitors projects for security vulnerabilities
// Homepage: https://snyk.io

import (
	"fmt"
	
	"os/exec"
)

func installSnykCli() {
	// Método 1: Descargar y extraer .tar.gz
	snykcli_tar_url := "https://registry.npmjs.org/snyk/-/snyk-1.1293.0.tgz"
	snykcli_cmd_tar := exec.Command("curl", "-L", snykcli_tar_url, "-o", "package.tar.gz")
	err := snykcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snykcli_zip_url := "https://registry.npmjs.org/snyk/-/snyk-1.1293.0.tgz"
	snykcli_cmd_zip := exec.Command("curl", "-L", snykcli_zip_url, "-o", "package.zip")
	err = snykcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snykcli_bin_url := "https://registry.npmjs.org/snyk/-/snyk-1.1293.0.tgz"
	snykcli_cmd_bin := exec.Command("curl", "-L", snykcli_bin_url, "-o", "binary.bin")
	err = snykcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snykcli_src_url := "https://registry.npmjs.org/snyk/-/snyk-1.1293.0.tgz"
	snykcli_cmd_src := exec.Command("curl", "-L", snykcli_src_url, "-o", "source.tar.gz")
	err = snykcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snykcli_cmd_direct := exec.Command("./binary")
	err = snykcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
