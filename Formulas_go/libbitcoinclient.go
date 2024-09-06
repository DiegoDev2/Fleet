package main

// LibbitcoinClient - Bitcoin Client Query Library
// Homepage: https://github.com/libbitcoin/libbitcoin-client

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinClient() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinclient_tar_url := "https://github.com/libbitcoin/libbitcoin-client/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinclient_cmd_tar := exec.Command("curl", "-L", libbitcoinclient_tar_url, "-o", "package.tar.gz")
	err := libbitcoinclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinclient_zip_url := "https://github.com/libbitcoin/libbitcoin-client/archive/refs/tags/v3.8.0.zip"
	libbitcoinclient_cmd_zip := exec.Command("curl", "-L", libbitcoinclient_zip_url, "-o", "package.zip")
	err = libbitcoinclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinclient_bin_url := "https://github.com/libbitcoin/libbitcoin-client/archive/refs/tags/v3.8.0.bin"
	libbitcoinclient_cmd_bin := exec.Command("curl", "-L", libbitcoinclient_bin_url, "-o", "binary.bin")
	err = libbitcoinclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinclient_src_url := "https://github.com/libbitcoin/libbitcoin-client/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinclient_cmd_src := exec.Command("curl", "-L", libbitcoinclient_src_url, "-o", "source.tar.gz")
	err = libbitcoinclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinclient_cmd_direct := exec.Command("./binary")
	err = libbitcoinclient_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libbitcoin-protocol")
exec.Command("latte", "install", "libbitcoin-protocol")
	fmt.Println("Instalando dependencia: libsodium")
exec.Command("latte", "install", "libsodium")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
}
