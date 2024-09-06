package main

// Cadence - Resource-oriented smart contract programming language
// Homepage: https://github.com/onflow/cadence

import (
	"fmt"
	
	"os/exec"
)

func installCadence() {
	// Método 1: Descargar y extraer .tar.gz
	cadence_tar_url := "https://github.com/onflow/cadence/archive/refs/tags/v0.42.12.tar.gz"
	cadence_cmd_tar := exec.Command("curl", "-L", cadence_tar_url, "-o", "package.tar.gz")
	err := cadence_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cadence_zip_url := "https://github.com/onflow/cadence/archive/refs/tags/v0.42.12.zip"
	cadence_cmd_zip := exec.Command("curl", "-L", cadence_zip_url, "-o", "package.zip")
	err = cadence_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cadence_bin_url := "https://github.com/onflow/cadence/archive/refs/tags/v0.42.12.bin"
	cadence_cmd_bin := exec.Command("curl", "-L", cadence_bin_url, "-o", "binary.bin")
	err = cadence_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cadence_src_url := "https://github.com/onflow/cadence/archive/refs/tags/v0.42.12.src.tar.gz"
	cadence_cmd_src := exec.Command("curl", "-L", cadence_src_url, "-o", "source.tar.gz")
	err = cadence_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cadence_cmd_direct := exec.Command("./binary")
	err = cadence_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
