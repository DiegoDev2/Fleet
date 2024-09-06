package main

// GitterCli - Extremely simple Gitter client for terminals
// Homepage: https://github.com/RodrigoEspinosa/gitter-cli

import (
	"fmt"
	
	"os/exec"
)

func installGitterCli() {
	// Método 1: Descargar y extraer .tar.gz
	gittercli_tar_url := "https://github.com/RodrigoEspinosa/gitter-cli/archive/refs/tags/v0.8.5.tar.gz"
	gittercli_cmd_tar := exec.Command("curl", "-L", gittercli_tar_url, "-o", "package.tar.gz")
	err := gittercli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gittercli_zip_url := "https://github.com/RodrigoEspinosa/gitter-cli/archive/refs/tags/v0.8.5.zip"
	gittercli_cmd_zip := exec.Command("curl", "-L", gittercli_zip_url, "-o", "package.zip")
	err = gittercli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gittercli_bin_url := "https://github.com/RodrigoEspinosa/gitter-cli/archive/refs/tags/v0.8.5.bin"
	gittercli_cmd_bin := exec.Command("curl", "-L", gittercli_bin_url, "-o", "binary.bin")
	err = gittercli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gittercli_src_url := "https://github.com/RodrigoEspinosa/gitter-cli/archive/refs/tags/v0.8.5.src.tar.gz"
	gittercli_cmd_src := exec.Command("curl", "-L", gittercli_src_url, "-o", "source.tar.gz")
	err = gittercli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gittercli_cmd_direct := exec.Command("./binary")
	err = gittercli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
