package main

// Nghttp2 - HTTP/2 C Library
// Homepage: https://nghttp2.org/

import (
	"fmt"
	
	"os/exec"
)

func installNghttp2() {
	// Método 1: Descargar y extraer .tar.gz
	nghttp2_tar_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.tar.gz"
	nghttp2_cmd_tar := exec.Command("curl", "-L", nghttp2_tar_url, "-o", "package.tar.gz")
	err := nghttp2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nghttp2_zip_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.zip"
	nghttp2_cmd_zip := exec.Command("curl", "-L", nghttp2_zip_url, "-o", "package.zip")
	err = nghttp2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nghttp2_bin_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.bin"
	nghttp2_cmd_bin := exec.Command("curl", "-L", nghttp2_bin_url, "-o", "binary.bin")
	err = nghttp2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nghttp2_src_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.src.tar.gz"
	nghttp2_cmd_src := exec.Command("curl", "-L", nghttp2_src_url, "-o", "source.tar.gz")
	err = nghttp2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nghttp2_cmd_direct := exec.Command("./binary")
	err = nghttp2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: c-ares")
exec.Command("latte", "install", "c-ares")
	fmt.Println("Instalando dependencia: jemalloc")
exec.Command("latte", "install", "jemalloc")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
