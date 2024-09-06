package main

// Libssh2 - C library implementing the SSH2 protocol
// Homepage: https://libssh2.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibssh2() {
	// Método 1: Descargar y extraer .tar.gz
	libssh2_tar_url := "https://libssh2.org/download/libssh2-1.11.0.tar.gz"
	libssh2_cmd_tar := exec.Command("curl", "-L", libssh2_tar_url, "-o", "package.tar.gz")
	err := libssh2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libssh2_zip_url := "https://libssh2.org/download/libssh2-1.11.0.zip"
	libssh2_cmd_zip := exec.Command("curl", "-L", libssh2_zip_url, "-o", "package.zip")
	err = libssh2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libssh2_bin_url := "https://libssh2.org/download/libssh2-1.11.0.bin"
	libssh2_cmd_bin := exec.Command("curl", "-L", libssh2_bin_url, "-o", "binary.bin")
	err = libssh2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libssh2_src_url := "https://libssh2.org/download/libssh2-1.11.0.src.tar.gz"
	libssh2_cmd_src := exec.Command("curl", "-L", libssh2_src_url, "-o", "source.tar.gz")
	err = libssh2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libssh2_cmd_direct := exec.Command("./binary")
	err = libssh2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
