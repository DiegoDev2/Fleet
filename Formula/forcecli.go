package main

// Forcecli - Command-line interface to Force.com
// Homepage: https://force-cli.herokuapp.com/

import (
	"fmt"
	
	"os/exec"
)

func installForcecli() {
	// Método 1: Descargar y extraer .tar.gz
	forcecli_tar_url := "https://github.com/ForceCLI/force/archive/refs/tags/v1.0.6.tar.gz"
	forcecli_cmd_tar := exec.Command("curl", "-L", forcecli_tar_url, "-o", "package.tar.gz")
	err := forcecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	forcecli_zip_url := "https://github.com/ForceCLI/force/archive/refs/tags/v1.0.6.zip"
	forcecli_cmd_zip := exec.Command("curl", "-L", forcecli_zip_url, "-o", "package.zip")
	err = forcecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	forcecli_bin_url := "https://github.com/ForceCLI/force/archive/refs/tags/v1.0.6.bin"
	forcecli_cmd_bin := exec.Command("curl", "-L", forcecli_bin_url, "-o", "binary.bin")
	err = forcecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	forcecli_src_url := "https://github.com/ForceCLI/force/archive/refs/tags/v1.0.6.src.tar.gz"
	forcecli_cmd_src := exec.Command("curl", "-L", forcecli_src_url, "-o", "source.tar.gz")
	err = forcecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	forcecli_cmd_direct := exec.Command("./binary")
	err = forcecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
