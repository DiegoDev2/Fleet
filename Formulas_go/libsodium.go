package main

// Libsodium - NaCl networking and cryptography library
// Homepage: https://libsodium.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibsodium() {
	// Método 1: Descargar y extraer .tar.gz
	libsodium_tar_url := "https://download.libsodium.org/libsodium/releases/libsodium-1.0.20.tar.gz"
	libsodium_cmd_tar := exec.Command("curl", "-L", libsodium_tar_url, "-o", "package.tar.gz")
	err := libsodium_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsodium_zip_url := "https://download.libsodium.org/libsodium/releases/libsodium-1.0.20.zip"
	libsodium_cmd_zip := exec.Command("curl", "-L", libsodium_zip_url, "-o", "package.zip")
	err = libsodium_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsodium_bin_url := "https://download.libsodium.org/libsodium/releases/libsodium-1.0.20.bin"
	libsodium_cmd_bin := exec.Command("curl", "-L", libsodium_bin_url, "-o", "binary.bin")
	err = libsodium_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsodium_src_url := "https://download.libsodium.org/libsodium/releases/libsodium-1.0.20.src.tar.gz"
	libsodium_cmd_src := exec.Command("curl", "-L", libsodium_src_url, "-o", "source.tar.gz")
	err = libsodium_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsodium_cmd_direct := exec.Command("./binary")
	err = libsodium_cmd_direct.Run()
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
