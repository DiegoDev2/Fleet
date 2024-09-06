package main

// Erigon - Implementation of Ethereum (execution client), on the efficiency frontier
// Homepage: https://github.com/ledgerwatch/erigon

import (
	"fmt"
	
	"os/exec"
)

func installErigon() {
	// Método 1: Descargar y extraer .tar.gz
	erigon_tar_url := "https://github.com/ledgerwatch/erigon/archive/refs/tags/v2.55.1.tar.gz"
	erigon_cmd_tar := exec.Command("curl", "-L", erigon_tar_url, "-o", "package.tar.gz")
	err := erigon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	erigon_zip_url := "https://github.com/ledgerwatch/erigon/archive/refs/tags/v2.55.1.zip"
	erigon_cmd_zip := exec.Command("curl", "-L", erigon_zip_url, "-o", "package.zip")
	err = erigon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	erigon_bin_url := "https://github.com/ledgerwatch/erigon/archive/refs/tags/v2.55.1.bin"
	erigon_cmd_bin := exec.Command("curl", "-L", erigon_bin_url, "-o", "binary.bin")
	err = erigon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	erigon_src_url := "https://github.com/ledgerwatch/erigon/archive/refs/tags/v2.55.1.src.tar.gz"
	erigon_cmd_src := exec.Command("curl", "-L", erigon_src_url, "-o", "source.tar.gz")
	err = erigon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	erigon_cmd_direct := exec.Command("./binary")
	err = erigon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
