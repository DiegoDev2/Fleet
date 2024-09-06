package main

// Libantlr3c - ANTLRv3 parsing library for C
// Homepage: https://www.antlr3.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibantlr3c() {
	// Método 1: Descargar y extraer .tar.gz
	libantlr3c_tar_url := "https://github.com/antlr/antlr3/archive/refs/tags/3.5.3.tar.gz"
	libantlr3c_cmd_tar := exec.Command("curl", "-L", libantlr3c_tar_url, "-o", "package.tar.gz")
	err := libantlr3c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libantlr3c_zip_url := "https://github.com/antlr/antlr3/archive/refs/tags/3.5.3.zip"
	libantlr3c_cmd_zip := exec.Command("curl", "-L", libantlr3c_zip_url, "-o", "package.zip")
	err = libantlr3c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libantlr3c_bin_url := "https://github.com/antlr/antlr3/archive/refs/tags/3.5.3.bin"
	libantlr3c_cmd_bin := exec.Command("curl", "-L", libantlr3c_bin_url, "-o", "binary.bin")
	err = libantlr3c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libantlr3c_src_url := "https://github.com/antlr/antlr3/archive/refs/tags/3.5.3.src.tar.gz"
	libantlr3c_cmd_src := exec.Command("curl", "-L", libantlr3c_src_url, "-o", "source.tar.gz")
	err = libantlr3c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libantlr3c_cmd_direct := exec.Command("./binary")
	err = libantlr3c_cmd_direct.Run()
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
