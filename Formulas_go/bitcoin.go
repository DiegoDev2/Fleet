package main

// Bitcoin - Decentralized, peer to peer payment network
// Homepage: https://bitcoincore.org/

import (
	"fmt"
	
	"os/exec"
)

func installBitcoin() {
	// Método 1: Descargar y extraer .tar.gz
	bitcoin_tar_url := "https://bitcoincore.org/bin/bitcoin-core-27.1/bitcoin-27.1.tar.gz"
	bitcoin_cmd_tar := exec.Command("curl", "-L", bitcoin_tar_url, "-o", "package.tar.gz")
	err := bitcoin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bitcoin_zip_url := "https://bitcoincore.org/bin/bitcoin-core-27.1/bitcoin-27.1.zip"
	bitcoin_cmd_zip := exec.Command("curl", "-L", bitcoin_zip_url, "-o", "package.zip")
	err = bitcoin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bitcoin_bin_url := "https://bitcoincore.org/bin/bitcoin-core-27.1/bitcoin-27.1.bin"
	bitcoin_cmd_bin := exec.Command("curl", "-L", bitcoin_bin_url, "-o", "binary.bin")
	err = bitcoin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bitcoin_src_url := "https://bitcoincore.org/bin/bitcoin-core-27.1/bitcoin-27.1.src.tar.gz"
	bitcoin_cmd_src := exec.Command("curl", "-L", bitcoin_src_url, "-o", "source.tar.gz")
	err = bitcoin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bitcoin_cmd_direct := exec.Command("./binary")
	err = bitcoin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: miniupnpc")
exec.Command("latte", "install", "miniupnpc")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
