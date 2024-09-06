package main

// Libzip - C library for reading, creating, and modifying zip archives
// Homepage: https://libzip.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibzip() {
	// Método 1: Descargar y extraer .tar.gz
	libzip_tar_url := "https://libzip.org/download/libzip-1.10.1.tar.xz"
	libzip_cmd_tar := exec.Command("curl", "-L", libzip_tar_url, "-o", "package.tar.gz")
	err := libzip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libzip_zip_url := "https://libzip.org/download/libzip-1.10.1.tar.xz"
	libzip_cmd_zip := exec.Command("curl", "-L", libzip_zip_url, "-o", "package.zip")
	err = libzip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libzip_bin_url := "https://libzip.org/download/libzip-1.10.1.tar.xz"
	libzip_cmd_bin := exec.Command("curl", "-L", libzip_bin_url, "-o", "binary.bin")
	err = libzip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libzip_src_url := "https://libzip.org/download/libzip-1.10.1.tar.xz"
	libzip_cmd_src := exec.Command("curl", "-L", libzip_src_url, "-o", "source.tar.gz")
	err = libzip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libzip_cmd_direct := exec.Command("./binary")
	err = libzip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
