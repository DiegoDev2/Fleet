package main

// Hledger - Easy plain text accounting with command-line, terminal and web UIs
// Homepage: https://hledger.org/

import (
	"fmt"
	
	"os/exec"
)

func installHledger() {
	// Método 1: Descargar y extraer .tar.gz
	hledger_tar_url := "https://github.com/simonmichael/hledger/archive/refs/tags/1.34.tar.gz"
	hledger_cmd_tar := exec.Command("curl", "-L", hledger_tar_url, "-o", "package.tar.gz")
	err := hledger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hledger_zip_url := "https://github.com/simonmichael/hledger/archive/refs/tags/1.34.zip"
	hledger_cmd_zip := exec.Command("curl", "-L", hledger_zip_url, "-o", "package.zip")
	err = hledger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hledger_bin_url := "https://github.com/simonmichael/hledger/archive/refs/tags/1.34.bin"
	hledger_cmd_bin := exec.Command("curl", "-L", hledger_bin_url, "-o", "binary.bin")
	err = hledger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hledger_src_url := "https://github.com/simonmichael/hledger/archive/refs/tags/1.34.src.tar.gz"
	hledger_cmd_src := exec.Command("curl", "-L", hledger_src_url, "-o", "source.tar.gz")
	err = hledger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hledger_cmd_direct := exec.Command("./binary")
	err = hledger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ghc@9.8")
	exec.Command("latte", "install", "ghc@9.8").Run()
	fmt.Println("Instalando dependencia: haskell-stack")
	exec.Command("latte", "install", "haskell-stack").Run()
}
