package main

// Sui - Next-generation smart contract platform powered by the Move programming language
// Homepage: https://sui.io

import (
	"fmt"
	
	"os/exec"
)

func installSui() {
	// Método 1: Descargar y extraer .tar.gz
	sui_tar_url := "https://github.com/MystenLabs/sui/archive/refs/tags/testnet-v1.32.0.tar.gz"
	sui_cmd_tar := exec.Command("curl", "-L", sui_tar_url, "-o", "package.tar.gz")
	err := sui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sui_zip_url := "https://github.com/MystenLabs/sui/archive/refs/tags/testnet-v1.32.0.zip"
	sui_cmd_zip := exec.Command("curl", "-L", sui_zip_url, "-o", "package.zip")
	err = sui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sui_bin_url := "https://github.com/MystenLabs/sui/archive/refs/tags/testnet-v1.32.0.bin"
	sui_cmd_bin := exec.Command("curl", "-L", sui_bin_url, "-o", "binary.bin")
	err = sui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sui_src_url := "https://github.com/MystenLabs/sui/archive/refs/tags/testnet-v1.32.0.src.tar.gz"
	sui_cmd_src := exec.Command("curl", "-L", sui_src_url, "-o", "source.tar.gz")
	err = sui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sui_cmd_direct := exec.Command("./binary")
	err = sui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
