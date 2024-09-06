package main

// Ethereum - Official Go implementation of the Ethereum protocol
// Homepage: https://geth.ethereum.org/

import (
	"fmt"
	
	"os/exec"
)

func installEthereum() {
	// Método 1: Descargar y extraer .tar.gz
	ethereum_tar_url := "https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.14.8.tar.gz"
	ethereum_cmd_tar := exec.Command("curl", "-L", ethereum_tar_url, "-o", "package.tar.gz")
	err := ethereum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ethereum_zip_url := "https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.14.8.zip"
	ethereum_cmd_zip := exec.Command("curl", "-L", ethereum_zip_url, "-o", "package.zip")
	err = ethereum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ethereum_bin_url := "https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.14.8.bin"
	ethereum_cmd_bin := exec.Command("curl", "-L", ethereum_bin_url, "-o", "binary.bin")
	err = ethereum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ethereum_src_url := "https://github.com/ethereum/go-ethereum/archive/refs/tags/v1.14.8.src.tar.gz"
	ethereum_cmd_src := exec.Command("curl", "-L", ethereum_src_url, "-o", "source.tar.gz")
	err = ethereum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ethereum_cmd_direct := exec.Command("./binary")
	err = ethereum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
