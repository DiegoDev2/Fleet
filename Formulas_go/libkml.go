package main

// Libkml - Library to parse, generate and operate on KML
// Homepage: https://github.com/libkml/libkml

import (
	"fmt"
	
	"os/exec"
)

func installLibkml() {
	// Método 1: Descargar y extraer .tar.gz
	libkml_tar_url := "https://github.com/libkml/libkml/archive/refs/tags/1.3.0.tar.gz"
	libkml_cmd_tar := exec.Command("curl", "-L", libkml_tar_url, "-o", "package.tar.gz")
	err := libkml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libkml_zip_url := "https://github.com/libkml/libkml/archive/refs/tags/1.3.0.zip"
	libkml_cmd_zip := exec.Command("curl", "-L", libkml_zip_url, "-o", "package.zip")
	err = libkml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libkml_bin_url := "https://github.com/libkml/libkml/archive/refs/tags/1.3.0.bin"
	libkml_cmd_bin := exec.Command("curl", "-L", libkml_bin_url, "-o", "binary.bin")
	err = libkml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libkml_src_url := "https://github.com/libkml/libkml/archive/refs/tags/1.3.0.src.tar.gz"
	libkml_cmd_src := exec.Command("curl", "-L", libkml_src_url, "-o", "source.tar.gz")
	err = libkml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libkml_cmd_direct := exec.Command("./binary")
	err = libkml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
	fmt.Println("Instalando dependencia: uriparser")
exec.Command("latte", "install", "uriparser")
}
