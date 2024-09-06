package main

// Medusa - Solidity smart contract fuzzer powered by go-ethereum
// Homepage: https://github.com/crytic/medusa

import (
	"fmt"
	
	"os/exec"
)

func installMedusa() {
	// Método 1: Descargar y extraer .tar.gz
	medusa_tar_url := "https://github.com/crytic/medusa/archive/refs/tags/v0.1.6.tar.gz"
	medusa_cmd_tar := exec.Command("curl", "-L", medusa_tar_url, "-o", "package.tar.gz")
	err := medusa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	medusa_zip_url := "https://github.com/crytic/medusa/archive/refs/tags/v0.1.6.zip"
	medusa_cmd_zip := exec.Command("curl", "-L", medusa_zip_url, "-o", "package.zip")
	err = medusa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	medusa_bin_url := "https://github.com/crytic/medusa/archive/refs/tags/v0.1.6.bin"
	medusa_cmd_bin := exec.Command("curl", "-L", medusa_bin_url, "-o", "binary.bin")
	err = medusa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	medusa_src_url := "https://github.com/crytic/medusa/archive/refs/tags/v0.1.6.src.tar.gz"
	medusa_cmd_src := exec.Command("curl", "-L", medusa_src_url, "-o", "source.tar.gz")
	err = medusa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	medusa_cmd_direct := exec.Command("./binary")
	err = medusa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: truffle")
exec.Command("latte", "install", "truffle")
	fmt.Println("Instalando dependencia: crytic-compile")
exec.Command("latte", "install", "crytic-compile")
}
