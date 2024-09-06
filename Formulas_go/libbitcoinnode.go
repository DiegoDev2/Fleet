package main

// LibbitcoinNode - Bitcoin Full Node
// Homepage: https://github.com/libbitcoin/libbitcoin-node

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinNode() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinnode_tar_url := "https://github.com/libbitcoin/libbitcoin-node/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinnode_cmd_tar := exec.Command("curl", "-L", libbitcoinnode_tar_url, "-o", "package.tar.gz")
	err := libbitcoinnode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinnode_zip_url := "https://github.com/libbitcoin/libbitcoin-node/archive/refs/tags/v3.8.0.zip"
	libbitcoinnode_cmd_zip := exec.Command("curl", "-L", libbitcoinnode_zip_url, "-o", "package.zip")
	err = libbitcoinnode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinnode_bin_url := "https://github.com/libbitcoin/libbitcoin-node/archive/refs/tags/v3.8.0.bin"
	libbitcoinnode_cmd_bin := exec.Command("curl", "-L", libbitcoinnode_bin_url, "-o", "binary.bin")
	err = libbitcoinnode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinnode_src_url := "https://github.com/libbitcoin/libbitcoin-node/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinnode_cmd_src := exec.Command("curl", "-L", libbitcoinnode_src_url, "-o", "source.tar.gz")
	err = libbitcoinnode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinnode_cmd_direct := exec.Command("./binary")
	err = libbitcoinnode_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost@1.76")
exec.Command("latte", "install", "boost@1.76")
	fmt.Println("Instalando dependencia: libbitcoin-blockchain")
exec.Command("latte", "install", "libbitcoin-blockchain")
	fmt.Println("Instalando dependencia: libbitcoin-network")
exec.Command("latte", "install", "libbitcoin-network")
}
