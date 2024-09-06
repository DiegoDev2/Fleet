package main

// Libnghttp2 - HTTP/2 C Library
// Homepage: https://nghttp2.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibnghttp2() {
	// Método 1: Descargar y extraer .tar.gz
	libnghttp2_tar_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.tar.gz"
	libnghttp2_cmd_tar := exec.Command("curl", "-L", libnghttp2_tar_url, "-o", "package.tar.gz")
	err := libnghttp2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnghttp2_zip_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.zip"
	libnghttp2_cmd_zip := exec.Command("curl", "-L", libnghttp2_zip_url, "-o", "package.zip")
	err = libnghttp2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnghttp2_bin_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.bin"
	libnghttp2_cmd_bin := exec.Command("curl", "-L", libnghttp2_bin_url, "-o", "binary.bin")
	err = libnghttp2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnghttp2_src_url := "https://github.com/nghttp2/nghttp2/releases/download/v1.63.0/nghttp2-1.63.0.src.tar.gz"
	libnghttp2_cmd_src := exec.Command("curl", "-L", libnghttp2_src_url, "-o", "source.tar.gz")
	err = libnghttp2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnghttp2_cmd_direct := exec.Command("./binary")
	err = libnghttp2_cmd_direct.Run()
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
}
