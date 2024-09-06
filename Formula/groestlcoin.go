package main

// Groestlcoin - Decentralized, peer to peer payment network
// Homepage: https://groestlcoin.org/groestlcoin-core-wallet/

import (
	"fmt"
	
	"os/exec"
)

func installGroestlcoin() {
	// Método 1: Descargar y extraer .tar.gz
	groestlcoin_tar_url := "https://github.com/Groestlcoin/groestlcoin/releases/download/v27.0/groestlcoin-27.0.tar.gz"
	groestlcoin_cmd_tar := exec.Command("curl", "-L", groestlcoin_tar_url, "-o", "package.tar.gz")
	err := groestlcoin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	groestlcoin_zip_url := "https://github.com/Groestlcoin/groestlcoin/releases/download/v27.0/groestlcoin-27.0.zip"
	groestlcoin_cmd_zip := exec.Command("curl", "-L", groestlcoin_zip_url, "-o", "package.zip")
	err = groestlcoin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	groestlcoin_bin_url := "https://github.com/Groestlcoin/groestlcoin/releases/download/v27.0/groestlcoin-27.0.bin"
	groestlcoin_cmd_bin := exec.Command("curl", "-L", groestlcoin_bin_url, "-o", "binary.bin")
	err = groestlcoin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	groestlcoin_src_url := "https://github.com/Groestlcoin/groestlcoin/releases/download/v27.0/groestlcoin-27.0.src.tar.gz"
	groestlcoin_cmd_src := exec.Command("curl", "-L", groestlcoin_src_url, "-o", "source.tar.gz")
	err = groestlcoin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	groestlcoin_cmd_direct := exec.Command("./binary")
	err = groestlcoin_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: berkeley-db@5")
	exec.Command("latte", "install", "berkeley-db@5").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: miniupnpc")
	exec.Command("latte", "install", "miniupnpc").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
