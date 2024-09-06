package main

// Libressl - Version of the SSL/TLS protocol forked from OpenSSL
// Homepage: https://www.libressl.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibressl() {
	// Método 1: Descargar y extraer .tar.gz
	libressl_tar_url := "https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-3.9.2.tar.gz"
	libressl_cmd_tar := exec.Command("curl", "-L", libressl_tar_url, "-o", "package.tar.gz")
	err := libressl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libressl_zip_url := "https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-3.9.2.zip"
	libressl_cmd_zip := exec.Command("curl", "-L", libressl_zip_url, "-o", "package.zip")
	err = libressl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libressl_bin_url := "https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-3.9.2.bin"
	libressl_cmd_bin := exec.Command("curl", "-L", libressl_bin_url, "-o", "binary.bin")
	err = libressl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libressl_src_url := "https://ftp.openbsd.org/pub/OpenBSD/LibreSSL/libressl-3.9.2.src.tar.gz"
	libressl_cmd_src := exec.Command("curl", "-L", libressl_src_url, "-o", "source.tar.gz")
	err = libressl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libressl_cmd_direct := exec.Command("./binary")
	err = libressl_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: ca-certificates")
	exec.Command("latte", "install", "ca-certificates").Run()
}
