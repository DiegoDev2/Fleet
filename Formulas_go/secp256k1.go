package main

// Secp256k1 - Optimized C library for EC operations on curve secp256k1
// Homepage: https://github.com/bitcoin-core/secp256k1

import (
	"fmt"
	
	"os/exec"
)

func installSecp256k1() {
	// Método 1: Descargar y extraer .tar.gz
	secp256k1_tar_url := "https://github.com/bitcoin-core/secp256k1/archive/refs/tags/v0.5.1.tar.gz"
	secp256k1_cmd_tar := exec.Command("curl", "-L", secp256k1_tar_url, "-o", "package.tar.gz")
	err := secp256k1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	secp256k1_zip_url := "https://github.com/bitcoin-core/secp256k1/archive/refs/tags/v0.5.1.zip"
	secp256k1_cmd_zip := exec.Command("curl", "-L", secp256k1_zip_url, "-o", "package.zip")
	err = secp256k1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	secp256k1_bin_url := "https://github.com/bitcoin-core/secp256k1/archive/refs/tags/v0.5.1.bin"
	secp256k1_cmd_bin := exec.Command("curl", "-L", secp256k1_bin_url, "-o", "binary.bin")
	err = secp256k1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	secp256k1_src_url := "https://github.com/bitcoin-core/secp256k1/archive/refs/tags/v0.5.1.src.tar.gz"
	secp256k1_cmd_src := exec.Command("curl", "-L", secp256k1_src_url, "-o", "source.tar.gz")
	err = secp256k1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	secp256k1_cmd_direct := exec.Command("./binary")
	err = secp256k1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
