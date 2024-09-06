package main

// Libdivecomputer - Library for communication with various dive computers
// Homepage: https://www.libdivecomputer.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibdivecomputer() {
	// Método 1: Descargar y extraer .tar.gz
	libdivecomputer_tar_url := "https://www.libdivecomputer.org/releases/libdivecomputer-0.8.0.tar.gz"
	libdivecomputer_cmd_tar := exec.Command("curl", "-L", libdivecomputer_tar_url, "-o", "package.tar.gz")
	err := libdivecomputer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdivecomputer_zip_url := "https://www.libdivecomputer.org/releases/libdivecomputer-0.8.0.zip"
	libdivecomputer_cmd_zip := exec.Command("curl", "-L", libdivecomputer_zip_url, "-o", "package.zip")
	err = libdivecomputer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdivecomputer_bin_url := "https://www.libdivecomputer.org/releases/libdivecomputer-0.8.0.bin"
	libdivecomputer_cmd_bin := exec.Command("curl", "-L", libdivecomputer_bin_url, "-o", "binary.bin")
	err = libdivecomputer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdivecomputer_src_url := "https://www.libdivecomputer.org/releases/libdivecomputer-0.8.0.src.tar.gz"
	libdivecomputer_cmd_src := exec.Command("curl", "-L", libdivecomputer_src_url, "-o", "source.tar.gz")
	err = libdivecomputer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdivecomputer_cmd_direct := exec.Command("./binary")
	err = libdivecomputer_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
