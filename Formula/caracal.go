package main

// Caracal - Static analyzer for Starknet smart contracts
// Homepage: https://github.com/crytic/caracal

import (
	"fmt"
	
	"os/exec"
)

func installCaracal() {
	// Método 1: Descargar y extraer .tar.gz
	caracal_tar_url := "https://github.com/crytic/caracal/archive/refs/tags/v0.2.3.tar.gz"
	caracal_cmd_tar := exec.Command("curl", "-L", caracal_tar_url, "-o", "package.tar.gz")
	err := caracal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	caracal_zip_url := "https://github.com/crytic/caracal/archive/refs/tags/v0.2.3.zip"
	caracal_cmd_zip := exec.Command("curl", "-L", caracal_zip_url, "-o", "package.zip")
	err = caracal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	caracal_bin_url := "https://github.com/crytic/caracal/archive/refs/tags/v0.2.3.bin"
	caracal_cmd_bin := exec.Command("curl", "-L", caracal_bin_url, "-o", "binary.bin")
	err = caracal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	caracal_src_url := "https://github.com/crytic/caracal/archive/refs/tags/v0.2.3.src.tar.gz"
	caracal_cmd_src := exec.Command("curl", "-L", caracal_src_url, "-o", "source.tar.gz")
	err = caracal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	caracal_cmd_direct := exec.Command("./binary")
	err = caracal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
