package main

// ChainloopCli - CLI for interacting with Chainloop
// Homepage: https://docs.chainloop.dev

import (
	"fmt"
	
	"os/exec"
)

func installChainloopCli() {
	// Método 1: Descargar y extraer .tar.gz
	chainloopcli_tar_url := "https://github.com/chainloop-dev/chainloop/archive/refs/tags/v0.96.5.tar.gz"
	chainloopcli_cmd_tar := exec.Command("curl", "-L", chainloopcli_tar_url, "-o", "package.tar.gz")
	err := chainloopcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chainloopcli_zip_url := "https://github.com/chainloop-dev/chainloop/archive/refs/tags/v0.96.5.zip"
	chainloopcli_cmd_zip := exec.Command("curl", "-L", chainloopcli_zip_url, "-o", "package.zip")
	err = chainloopcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chainloopcli_bin_url := "https://github.com/chainloop-dev/chainloop/archive/refs/tags/v0.96.5.bin"
	chainloopcli_cmd_bin := exec.Command("curl", "-L", chainloopcli_bin_url, "-o", "binary.bin")
	err = chainloopcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chainloopcli_src_url := "https://github.com/chainloop-dev/chainloop/archive/refs/tags/v0.96.5.src.tar.gz"
	chainloopcli_cmd_src := exec.Command("curl", "-L", chainloopcli_src_url, "-o", "source.tar.gz")
	err = chainloopcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chainloopcli_cmd_direct := exec.Command("./binary")
	err = chainloopcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
