package main

// LibbitcoinProtocol - Bitcoin Blockchain Query Protocol
// Homepage: https://github.com/libbitcoin/libbitcoin-protocol

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinProtocol() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinprotocol_tar_url := "https://github.com/libbitcoin/libbitcoin-protocol/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinprotocol_cmd_tar := exec.Command("curl", "-L", libbitcoinprotocol_tar_url, "-o", "package.tar.gz")
	err := libbitcoinprotocol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinprotocol_zip_url := "https://github.com/libbitcoin/libbitcoin-protocol/archive/refs/tags/v3.8.0.zip"
	libbitcoinprotocol_cmd_zip := exec.Command("curl", "-L", libbitcoinprotocol_zip_url, "-o", "package.zip")
	err = libbitcoinprotocol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinprotocol_bin_url := "https://github.com/libbitcoin/libbitcoin-protocol/archive/refs/tags/v3.8.0.bin"
	libbitcoinprotocol_cmd_bin := exec.Command("curl", "-L", libbitcoinprotocol_bin_url, "-o", "binary.bin")
	err = libbitcoinprotocol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinprotocol_src_url := "https://github.com/libbitcoin/libbitcoin-protocol/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinprotocol_cmd_src := exec.Command("curl", "-L", libbitcoinprotocol_src_url, "-o", "source.tar.gz")
	err = libbitcoinprotocol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinprotocol_cmd_direct := exec.Command("./binary")
	err = libbitcoinprotocol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost@1.76")
	exec.Command("latte", "install", "boost@1.76").Run()
	fmt.Println("Instalando dependencia: libbitcoin-system")
	exec.Command("latte", "install", "libbitcoin-system").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}
