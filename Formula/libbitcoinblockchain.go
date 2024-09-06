package main

// LibbitcoinBlockchain - Bitcoin Blockchain Library
// Homepage: https://github.com/libbitcoin/libbitcoin-blockchain

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinBlockchain() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinblockchain_tar_url := "https://github.com/libbitcoin/libbitcoin-blockchain/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinblockchain_cmd_tar := exec.Command("curl", "-L", libbitcoinblockchain_tar_url, "-o", "package.tar.gz")
	err := libbitcoinblockchain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinblockchain_zip_url := "https://github.com/libbitcoin/libbitcoin-blockchain/archive/refs/tags/v3.8.0.zip"
	libbitcoinblockchain_cmd_zip := exec.Command("curl", "-L", libbitcoinblockchain_zip_url, "-o", "package.zip")
	err = libbitcoinblockchain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinblockchain_bin_url := "https://github.com/libbitcoin/libbitcoin-blockchain/archive/refs/tags/v3.8.0.bin"
	libbitcoinblockchain_cmd_bin := exec.Command("curl", "-L", libbitcoinblockchain_bin_url, "-o", "binary.bin")
	err = libbitcoinblockchain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinblockchain_src_url := "https://github.com/libbitcoin/libbitcoin-blockchain/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinblockchain_cmd_src := exec.Command("curl", "-L", libbitcoinblockchain_src_url, "-o", "source.tar.gz")
	err = libbitcoinblockchain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinblockchain_cmd_direct := exec.Command("./binary")
	err = libbitcoinblockchain_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libbitcoin-consensus")
	exec.Command("latte", "install", "libbitcoin-consensus").Run()
	fmt.Println("Instalando dependencia: libbitcoin-database")
	exec.Command("latte", "install", "libbitcoin-database").Run()
}
