package main

// Firefly - Create and manage the Hyperledger FireFly stack for blockchain interaction
// Homepage: https://hyperledger.github.io/firefly/latest/

import (
	"fmt"
	
	"os/exec"
)

func installFirefly() {
	// Método 1: Descargar y extraer .tar.gz
	firefly_tar_url := "https://github.com/hyperledger/firefly-cli/archive/refs/tags/v1.3.1.tar.gz"
	firefly_cmd_tar := exec.Command("curl", "-L", firefly_tar_url, "-o", "package.tar.gz")
	err := firefly_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	firefly_zip_url := "https://github.com/hyperledger/firefly-cli/archive/refs/tags/v1.3.1.zip"
	firefly_cmd_zip := exec.Command("curl", "-L", firefly_zip_url, "-o", "package.zip")
	err = firefly_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	firefly_bin_url := "https://github.com/hyperledger/firefly-cli/archive/refs/tags/v1.3.1.bin"
	firefly_cmd_bin := exec.Command("curl", "-L", firefly_bin_url, "-o", "binary.bin")
	err = firefly_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	firefly_src_url := "https://github.com/hyperledger/firefly-cli/archive/refs/tags/v1.3.1.src.tar.gz"
	firefly_cmd_src := exec.Command("curl", "-L", firefly_src_url, "-o", "source.tar.gz")
	err = firefly_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	firefly_cmd_direct := exec.Command("./binary")
	err = firefly_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
