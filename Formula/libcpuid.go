package main

// Libcpuid - Small C library for x86 CPU detection and feature extraction
// Homepage: https://github.com/anrieff/libcpuid

import (
	"fmt"
	
	"os/exec"
)

func installLibcpuid() {
	// Método 1: Descargar y extraer .tar.gz
	libcpuid_tar_url := "https://github.com/anrieff/libcpuid/archive/refs/tags/v0.7.0.tar.gz"
	libcpuid_cmd_tar := exec.Command("curl", "-L", libcpuid_tar_url, "-o", "package.tar.gz")
	err := libcpuid_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcpuid_zip_url := "https://github.com/anrieff/libcpuid/archive/refs/tags/v0.7.0.zip"
	libcpuid_cmd_zip := exec.Command("curl", "-L", libcpuid_zip_url, "-o", "package.zip")
	err = libcpuid_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcpuid_bin_url := "https://github.com/anrieff/libcpuid/archive/refs/tags/v0.7.0.bin"
	libcpuid_cmd_bin := exec.Command("curl", "-L", libcpuid_bin_url, "-o", "binary.bin")
	err = libcpuid_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcpuid_src_url := "https://github.com/anrieff/libcpuid/archive/refs/tags/v0.7.0.src.tar.gz"
	libcpuid_cmd_src := exec.Command("curl", "-L", libcpuid_src_url, "-o", "source.tar.gz")
	err = libcpuid_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcpuid_cmd_direct := exec.Command("./binary")
	err = libcpuid_cmd_direct.Run()
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
}
