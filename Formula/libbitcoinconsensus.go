package main

// LibbitcoinConsensus - Bitcoin Consensus Library (optional)
// Homepage: https://github.com/libbitcoin/libbitcoin-consensus

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinConsensus() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoinconsensus_tar_url := "https://github.com/libbitcoin/libbitcoin-consensus/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoinconsensus_cmd_tar := exec.Command("curl", "-L", libbitcoinconsensus_tar_url, "-o", "package.tar.gz")
	err := libbitcoinconsensus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoinconsensus_zip_url := "https://github.com/libbitcoin/libbitcoin-consensus/archive/refs/tags/v3.8.0.zip"
	libbitcoinconsensus_cmd_zip := exec.Command("curl", "-L", libbitcoinconsensus_zip_url, "-o", "package.zip")
	err = libbitcoinconsensus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoinconsensus_bin_url := "https://github.com/libbitcoin/libbitcoin-consensus/archive/refs/tags/v3.8.0.bin"
	libbitcoinconsensus_cmd_bin := exec.Command("curl", "-L", libbitcoinconsensus_bin_url, "-o", "binary.bin")
	err = libbitcoinconsensus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoinconsensus_src_url := "https://github.com/libbitcoin/libbitcoin-consensus/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoinconsensus_cmd_src := exec.Command("curl", "-L", libbitcoinconsensus_src_url, "-o", "source.tar.gz")
	err = libbitcoinconsensus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoinconsensus_cmd_direct := exec.Command("./binary")
	err = libbitcoinconsensus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
