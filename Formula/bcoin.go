package main

// Bcoin - Javascript bitcoin library for node.js and browsers
// Homepage: https://bcoin.io

import (
	"fmt"
	
	"os/exec"
)

func installBcoin() {
	// Método 1: Descargar y extraer .tar.gz
	bcoin_tar_url := "https://github.com/bcoin-org/bcoin/archive/refs/tags/v2.2.0.tar.gz"
	bcoin_cmd_tar := exec.Command("curl", "-L", bcoin_tar_url, "-o", "package.tar.gz")
	err := bcoin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bcoin_zip_url := "https://github.com/bcoin-org/bcoin/archive/refs/tags/v2.2.0.zip"
	bcoin_cmd_zip := exec.Command("curl", "-L", bcoin_zip_url, "-o", "package.zip")
	err = bcoin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bcoin_bin_url := "https://github.com/bcoin-org/bcoin/archive/refs/tags/v2.2.0.bin"
	bcoin_cmd_bin := exec.Command("curl", "-L", bcoin_bin_url, "-o", "binary.bin")
	err = bcoin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bcoin_src_url := "https://github.com/bcoin-org/bcoin/archive/refs/tags/v2.2.0.src.tar.gz"
	bcoin_cmd_src := exec.Command("curl", "-L", bcoin_src_url, "-o", "source.tar.gz")
	err = bcoin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bcoin_cmd_direct := exec.Command("./binary")
	err = bcoin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
