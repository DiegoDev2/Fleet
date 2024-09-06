package main

// Ahoy - Creates self documenting CLI programs from commands in YAML files
// Homepage: https://ahoy-cli.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installAhoy() {
	// Método 1: Descargar y extraer .tar.gz
	ahoy_tar_url := "https://github.com/ahoy-cli/ahoy/archive/refs/tags/v2.2.0.tar.gz"
	ahoy_cmd_tar := exec.Command("curl", "-L", ahoy_tar_url, "-o", "package.tar.gz")
	err := ahoy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ahoy_zip_url := "https://github.com/ahoy-cli/ahoy/archive/refs/tags/v2.2.0.zip"
	ahoy_cmd_zip := exec.Command("curl", "-L", ahoy_zip_url, "-o", "package.zip")
	err = ahoy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ahoy_bin_url := "https://github.com/ahoy-cli/ahoy/archive/refs/tags/v2.2.0.bin"
	ahoy_cmd_bin := exec.Command("curl", "-L", ahoy_bin_url, "-o", "binary.bin")
	err = ahoy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ahoy_src_url := "https://github.com/ahoy-cli/ahoy/archive/refs/tags/v2.2.0.src.tar.gz"
	ahoy_cmd_src := exec.Command("curl", "-L", ahoy_src_url, "-o", "source.tar.gz")
	err = ahoy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ahoy_cmd_direct := exec.Command("./binary")
	err = ahoy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
