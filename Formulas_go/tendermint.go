package main

// Tendermint - BFT state machine replication for applications in any programming languages
// Homepage: https://tendermint.com/

import (
	"fmt"
	
	"os/exec"
)

func installTendermint() {
	// Método 1: Descargar y extraer .tar.gz
	tendermint_tar_url := "https://github.com/tendermint/tendermint/archive/refs/tags/v0.35.9.tar.gz"
	tendermint_cmd_tar := exec.Command("curl", "-L", tendermint_tar_url, "-o", "package.tar.gz")
	err := tendermint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tendermint_zip_url := "https://github.com/tendermint/tendermint/archive/refs/tags/v0.35.9.zip"
	tendermint_cmd_zip := exec.Command("curl", "-L", tendermint_zip_url, "-o", "package.zip")
	err = tendermint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tendermint_bin_url := "https://github.com/tendermint/tendermint/archive/refs/tags/v0.35.9.bin"
	tendermint_cmd_bin := exec.Command("curl", "-L", tendermint_bin_url, "-o", "binary.bin")
	err = tendermint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tendermint_src_url := "https://github.com/tendermint/tendermint/archive/refs/tags/v0.35.9.src.tar.gz"
	tendermint_cmd_src := exec.Command("curl", "-L", tendermint_src_url, "-o", "source.tar.gz")
	err = tendermint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tendermint_cmd_direct := exec.Command("./binary")
	err = tendermint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
