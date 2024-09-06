package main

// JfrogCli - Command-line interface for JFrog products
// Homepage: https://www.jfrog.com/confluence/display/CLI/JFrog+CLI

import (
	"fmt"
	
	"os/exec"
)

func installJfrogCli() {
	// Método 1: Descargar y extraer .tar.gz
	jfrogcli_tar_url := "https://github.com/jfrog/jfrog-cli/archive/refs/tags/v2.67.0.tar.gz"
	jfrogcli_cmd_tar := exec.Command("curl", "-L", jfrogcli_tar_url, "-o", "package.tar.gz")
	err := jfrogcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jfrogcli_zip_url := "https://github.com/jfrog/jfrog-cli/archive/refs/tags/v2.67.0.zip"
	jfrogcli_cmd_zip := exec.Command("curl", "-L", jfrogcli_zip_url, "-o", "package.zip")
	err = jfrogcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jfrogcli_bin_url := "https://github.com/jfrog/jfrog-cli/archive/refs/tags/v2.67.0.bin"
	jfrogcli_cmd_bin := exec.Command("curl", "-L", jfrogcli_bin_url, "-o", "binary.bin")
	err = jfrogcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jfrogcli_src_url := "https://github.com/jfrog/jfrog-cli/archive/refs/tags/v2.67.0.src.tar.gz"
	jfrogcli_cmd_src := exec.Command("curl", "-L", jfrogcli_src_url, "-o", "source.tar.gz")
	err = jfrogcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jfrogcli_cmd_direct := exec.Command("./binary")
	err = jfrogcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
