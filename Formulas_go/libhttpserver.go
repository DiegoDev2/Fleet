package main

// Libhttpserver - C++ library of embedded Rest HTTP server
// Homepage: https://github.com/etr/libhttpserver

import (
	"fmt"
	
	"os/exec"
)

func installLibhttpserver() {
	// Método 1: Descargar y extraer .tar.gz
	libhttpserver_tar_url := "https://github.com/etr/libhttpserver/archive/refs/tags/0.19.0.tar.gz"
	libhttpserver_cmd_tar := exec.Command("curl", "-L", libhttpserver_tar_url, "-o", "package.tar.gz")
	err := libhttpserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libhttpserver_zip_url := "https://github.com/etr/libhttpserver/archive/refs/tags/0.19.0.zip"
	libhttpserver_cmd_zip := exec.Command("curl", "-L", libhttpserver_zip_url, "-o", "package.zip")
	err = libhttpserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libhttpserver_bin_url := "https://github.com/etr/libhttpserver/archive/refs/tags/0.19.0.bin"
	libhttpserver_cmd_bin := exec.Command("curl", "-L", libhttpserver_bin_url, "-o", "binary.bin")
	err = libhttpserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libhttpserver_src_url := "https://github.com/etr/libhttpserver/archive/refs/tags/0.19.0.src.tar.gz"
	libhttpserver_cmd_src := exec.Command("curl", "-L", libhttpserver_src_url, "-o", "source.tar.gz")
	err = libhttpserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libhttpserver_cmd_direct := exec.Command("./binary")
	err = libhttpserver_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libmicrohttpd")
exec.Command("latte", "install", "libmicrohttpd")
}
