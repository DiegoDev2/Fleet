package main

// LeanCli - Command-line tool to develop and manage LeanCloud apps
// Homepage: https://github.com/leancloud/lean-cli

import (
	"fmt"
	
	"os/exec"
)

func installLeanCli() {
	// Método 1: Descargar y extraer .tar.gz
	leancli_tar_url := "https://github.com/leancloud/lean-cli/archive/refs/tags/v1.2.4.tar.gz"
	leancli_cmd_tar := exec.Command("curl", "-L", leancli_tar_url, "-o", "package.tar.gz")
	err := leancli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leancli_zip_url := "https://github.com/leancloud/lean-cli/archive/refs/tags/v1.2.4.zip"
	leancli_cmd_zip := exec.Command("curl", "-L", leancli_zip_url, "-o", "package.zip")
	err = leancli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leancli_bin_url := "https://github.com/leancloud/lean-cli/archive/refs/tags/v1.2.4.bin"
	leancli_cmd_bin := exec.Command("curl", "-L", leancli_bin_url, "-o", "binary.bin")
	err = leancli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leancli_src_url := "https://github.com/leancloud/lean-cli/archive/refs/tags/v1.2.4.src.tar.gz"
	leancli_cmd_src := exec.Command("curl", "-L", leancli_src_url, "-o", "source.tar.gz")
	err = leancli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leancli_cmd_direct := exec.Command("./binary")
	err = leancli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
