package main

// Libslax - Implementation of the SLAX language (an XSLT alternative)
// Homepage: http://www.libslax.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibslax() {
	// Método 1: Descargar y extraer .tar.gz
	libslax_tar_url := "https://github.com/Juniper/libslax/releases/download/0.22.1/libslax-0.22.1.tar.gz"
	libslax_cmd_tar := exec.Command("curl", "-L", libslax_tar_url, "-o", "package.tar.gz")
	err := libslax_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libslax_zip_url := "https://github.com/Juniper/libslax/releases/download/0.22.1/libslax-0.22.1.zip"
	libslax_cmd_zip := exec.Command("curl", "-L", libslax_zip_url, "-o", "package.zip")
	err = libslax_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libslax_bin_url := "https://github.com/Juniper/libslax/releases/download/0.22.1/libslax-0.22.1.bin"
	libslax_cmd_bin := exec.Command("curl", "-L", libslax_bin_url, "-o", "binary.bin")
	err = libslax_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libslax_src_url := "https://github.com/Juniper/libslax/releases/download/0.22.1/libslax-0.22.1.src.tar.gz"
	libslax_cmd_src := exec.Command("curl", "-L", libslax_src_url, "-o", "source.tar.gz")
	err = libslax_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libslax_cmd_direct := exec.Command("./binary")
	err = libslax_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
