package main

// Liblouis - Open-source braille translator and back-translator
// Homepage: https://liblouis.io

import (
	"fmt"
	
	"os/exec"
)

func installLiblouis() {
	// Método 1: Descargar y extraer .tar.gz
	liblouis_tar_url := "https://github.com/liblouis/liblouis/releases/download/v3.31.0/liblouis-3.31.0.tar.gz"
	liblouis_cmd_tar := exec.Command("curl", "-L", liblouis_tar_url, "-o", "package.tar.gz")
	err := liblouis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblouis_zip_url := "https://github.com/liblouis/liblouis/releases/download/v3.31.0/liblouis-3.31.0.zip"
	liblouis_cmd_zip := exec.Command("curl", "-L", liblouis_zip_url, "-o", "package.zip")
	err = liblouis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblouis_bin_url := "https://github.com/liblouis/liblouis/releases/download/v3.31.0/liblouis-3.31.0.bin"
	liblouis_cmd_bin := exec.Command("curl", "-L", liblouis_bin_url, "-o", "binary.bin")
	err = liblouis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblouis_src_url := "https://github.com/liblouis/liblouis/releases/download/v3.31.0/liblouis-3.31.0.src.tar.gz"
	liblouis_cmd_src := exec.Command("curl", "-L", liblouis_src_url, "-o", "source.tar.gz")
	err = liblouis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblouis_cmd_direct := exec.Command("./binary")
	err = liblouis_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
