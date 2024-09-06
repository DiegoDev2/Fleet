package main

// Libabigail - ABI Generic Analysis and Instrumentation Library
// Homepage: https://sourceware.org/libabigail/

import (
	"fmt"
	
	"os/exec"
)

func installLibabigail() {
	// Método 1: Descargar y extraer .tar.gz
	libabigail_tar_url := "https://mirrors.kernel.org/sourceware/libabigail/libabigail-2.5.tar.xz"
	libabigail_cmd_tar := exec.Command("curl", "-L", libabigail_tar_url, "-o", "package.tar.gz")
	err := libabigail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libabigail_zip_url := "https://mirrors.kernel.org/sourceware/libabigail/libabigail-2.5.tar.xz"
	libabigail_cmd_zip := exec.Command("curl", "-L", libabigail_zip_url, "-o", "package.zip")
	err = libabigail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libabigail_bin_url := "https://mirrors.kernel.org/sourceware/libabigail/libabigail-2.5.tar.xz"
	libabigail_cmd_bin := exec.Command("curl", "-L", libabigail_bin_url, "-o", "binary.bin")
	err = libabigail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libabigail_src_url := "https://mirrors.kernel.org/sourceware/libabigail/libabigail-2.5.tar.xz"
	libabigail_cmd_src := exec.Command("curl", "-L", libabigail_src_url, "-o", "source.tar.gz")
	err = libabigail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libabigail_cmd_direct := exec.Command("./binary")
	err = libabigail_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
}
