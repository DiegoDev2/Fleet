package main

// Truffle - Development environment, testing framework and asset pipeline for Ethereum
// Homepage: https://trufflesuite.com

import (
	"fmt"
	
	"os/exec"
)

func installTruffle() {
	// Método 1: Descargar y extraer .tar.gz
	truffle_tar_url := "https://registry.npmjs.org/truffle/-/truffle-5.11.5.tgz"
	truffle_cmd_tar := exec.Command("curl", "-L", truffle_tar_url, "-o", "package.tar.gz")
	err := truffle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	truffle_zip_url := "https://registry.npmjs.org/truffle/-/truffle-5.11.5.tgz"
	truffle_cmd_zip := exec.Command("curl", "-L", truffle_zip_url, "-o", "package.zip")
	err = truffle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	truffle_bin_url := "https://registry.npmjs.org/truffle/-/truffle-5.11.5.tgz"
	truffle_cmd_bin := exec.Command("curl", "-L", truffle_bin_url, "-o", "binary.bin")
	err = truffle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	truffle_src_url := "https://registry.npmjs.org/truffle/-/truffle-5.11.5.tgz"
	truffle_cmd_src := exec.Command("curl", "-L", truffle_src_url, "-o", "source.tar.gz")
	err = truffle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	truffle_cmd_direct := exec.Command("./binary")
	err = truffle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
