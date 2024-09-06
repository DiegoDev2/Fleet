package main

// Libmicrohttpd - Light HTTP/1.1 server library
// Homepage: https://www.gnu.org/software/libmicrohttpd/

import (
	"fmt"
	
	"os/exec"
)

func installLibmicrohttpd() {
	// Método 1: Descargar y extraer .tar.gz
	libmicrohttpd_tar_url := "https://ftp.gnu.org/gnu/libmicrohttpd/libmicrohttpd-1.0.1.tar.gz"
	libmicrohttpd_cmd_tar := exec.Command("curl", "-L", libmicrohttpd_tar_url, "-o", "package.tar.gz")
	err := libmicrohttpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmicrohttpd_zip_url := "https://ftp.gnu.org/gnu/libmicrohttpd/libmicrohttpd-1.0.1.zip"
	libmicrohttpd_cmd_zip := exec.Command("curl", "-L", libmicrohttpd_zip_url, "-o", "package.zip")
	err = libmicrohttpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmicrohttpd_bin_url := "https://ftp.gnu.org/gnu/libmicrohttpd/libmicrohttpd-1.0.1.bin"
	libmicrohttpd_cmd_bin := exec.Command("curl", "-L", libmicrohttpd_bin_url, "-o", "binary.bin")
	err = libmicrohttpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmicrohttpd_src_url := "https://ftp.gnu.org/gnu/libmicrohttpd/libmicrohttpd-1.0.1.src.tar.gz"
	libmicrohttpd_cmd_src := exec.Command("curl", "-L", libmicrohttpd_src_url, "-o", "source.tar.gz")
	err = libmicrohttpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmicrohttpd_cmd_direct := exec.Command("./binary")
	err = libmicrohttpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
}
