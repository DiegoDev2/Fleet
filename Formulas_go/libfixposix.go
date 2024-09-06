package main

// Libfixposix - Thin wrapper over POSIX syscalls
// Homepage: https://github.com/sionescu/libfixposix

import (
	"fmt"
	
	"os/exec"
)

func installLibfixposix() {
	// Método 1: Descargar y extraer .tar.gz
	libfixposix_tar_url := "https://github.com/sionescu/libfixposix/archive/refs/tags/v0.5.1.tar.gz"
	libfixposix_cmd_tar := exec.Command("curl", "-L", libfixposix_tar_url, "-o", "package.tar.gz")
	err := libfixposix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libfixposix_zip_url := "https://github.com/sionescu/libfixposix/archive/refs/tags/v0.5.1.zip"
	libfixposix_cmd_zip := exec.Command("curl", "-L", libfixposix_zip_url, "-o", "package.zip")
	err = libfixposix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libfixposix_bin_url := "https://github.com/sionescu/libfixposix/archive/refs/tags/v0.5.1.bin"
	libfixposix_cmd_bin := exec.Command("curl", "-L", libfixposix_bin_url, "-o", "binary.bin")
	err = libfixposix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libfixposix_src_url := "https://github.com/sionescu/libfixposix/archive/refs/tags/v0.5.1.src.tar.gz"
	libfixposix_cmd_src := exec.Command("curl", "-L", libfixposix_src_url, "-o", "source.tar.gz")
	err = libfixposix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libfixposix_cmd_direct := exec.Command("./binary")
	err = libfixposix_cmd_direct.Run()
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
