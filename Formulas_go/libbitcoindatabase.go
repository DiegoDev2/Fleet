package main

// LibbitcoinDatabase - Bitcoin High Performance Blockchain Database
// Homepage: https://github.com/libbitcoin/libbitcoin-database

import (
	"fmt"
	
	"os/exec"
)

func installLibbitcoinDatabase() {
	// Método 1: Descargar y extraer .tar.gz
	libbitcoindatabase_tar_url := "https://github.com/libbitcoin/libbitcoin-database/archive/refs/tags/v3.8.0.tar.gz"
	libbitcoindatabase_cmd_tar := exec.Command("curl", "-L", libbitcoindatabase_tar_url, "-o", "package.tar.gz")
	err := libbitcoindatabase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbitcoindatabase_zip_url := "https://github.com/libbitcoin/libbitcoin-database/archive/refs/tags/v3.8.0.zip"
	libbitcoindatabase_cmd_zip := exec.Command("curl", "-L", libbitcoindatabase_zip_url, "-o", "package.zip")
	err = libbitcoindatabase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbitcoindatabase_bin_url := "https://github.com/libbitcoin/libbitcoin-database/archive/refs/tags/v3.8.0.bin"
	libbitcoindatabase_cmd_bin := exec.Command("curl", "-L", libbitcoindatabase_bin_url, "-o", "binary.bin")
	err = libbitcoindatabase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbitcoindatabase_src_url := "https://github.com/libbitcoin/libbitcoin-database/archive/refs/tags/v3.8.0.src.tar.gz"
	libbitcoindatabase_cmd_src := exec.Command("curl", "-L", libbitcoindatabase_src_url, "-o", "source.tar.gz")
	err = libbitcoindatabase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbitcoindatabase_cmd_direct := exec.Command("./binary")
	err = libbitcoindatabase_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libbitcoin-system")
exec.Command("latte", "install", "libbitcoin-system")
}
