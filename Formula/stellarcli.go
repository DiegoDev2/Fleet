package main

// StellarCli - Stellar command-line tool for interacting with the Stellar network
// Homepage: https://developers.stellar.org

import (
	"fmt"
	
	"os/exec"
)

func installStellarCli() {
	// Método 1: Descargar y extraer .tar.gz
	stellarcli_tar_url := "https://github.com/stellar/stellar-cli/archive/refs/tags/v21.4.1.tar.gz"
	stellarcli_cmd_tar := exec.Command("curl", "-L", stellarcli_tar_url, "-o", "package.tar.gz")
	err := stellarcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stellarcli_zip_url := "https://github.com/stellar/stellar-cli/archive/refs/tags/v21.4.1.zip"
	stellarcli_cmd_zip := exec.Command("curl", "-L", stellarcli_zip_url, "-o", "package.zip")
	err = stellarcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stellarcli_bin_url := "https://github.com/stellar/stellar-cli/archive/refs/tags/v21.4.1.bin"
	stellarcli_cmd_bin := exec.Command("curl", "-L", stellarcli_bin_url, "-o", "binary.bin")
	err = stellarcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stellarcli_src_url := "https://github.com/stellar/stellar-cli/archive/refs/tags/v21.4.1.src.tar.gz"
	stellarcli_cmd_src := exec.Command("curl", "-L", stellarcli_src_url, "-o", "source.tar.gz")
	err = stellarcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stellarcli_cmd_direct := exec.Command("./binary")
	err = stellarcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
