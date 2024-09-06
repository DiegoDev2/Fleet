package main

// Libfixbuf - Implements the IPFIX Protocol as a C library
// Homepage: https://tools.netsa.cert.org/fixbuf/

import (
	"fmt"
	
	"os/exec"
)

func installLibfixbuf() {
	// Método 1: Descargar y extraer .tar.gz
	libfixbuf_tar_url := "https://tools.netsa.cert.org/releases/libfixbuf-2.5.0.tar.gz"
	libfixbuf_cmd_tar := exec.Command("curl", "-L", libfixbuf_tar_url, "-o", "package.tar.gz")
	err := libfixbuf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfixbuf_zip_url := "https://tools.netsa.cert.org/releases/libfixbuf-2.5.0.zip"
	libfixbuf_cmd_zip := exec.Command("curl", "-L", libfixbuf_zip_url, "-o", "package.zip")
	err = libfixbuf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfixbuf_bin_url := "https://tools.netsa.cert.org/releases/libfixbuf-2.5.0.bin"
	libfixbuf_cmd_bin := exec.Command("curl", "-L", libfixbuf_bin_url, "-o", "binary.bin")
	err = libfixbuf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfixbuf_src_url := "https://tools.netsa.cert.org/releases/libfixbuf-2.5.0.src.tar.gz"
	libfixbuf_cmd_src := exec.Command("curl", "-L", libfixbuf_src_url, "-o", "source.tar.gz")
	err = libfixbuf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfixbuf_cmd_direct := exec.Command("./binary")
	err = libfixbuf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
