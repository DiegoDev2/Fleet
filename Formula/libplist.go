package main

// Libplist - Library for Apple Binary- and XML-Property Lists
// Homepage: https://www.libimobiledevice.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibplist() {
	// Método 1: Descargar y extraer .tar.gz
	libplist_tar_url := "https://github.com/libimobiledevice/libplist/releases/download/2.6.0/libplist-2.6.0.tar.bz2"
	libplist_cmd_tar := exec.Command("curl", "-L", libplist_tar_url, "-o", "package.tar.gz")
	err := libplist_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libplist_zip_url := "https://github.com/libimobiledevice/libplist/releases/download/2.6.0/libplist-2.6.0.tar.bz2"
	libplist_cmd_zip := exec.Command("curl", "-L", libplist_zip_url, "-o", "package.zip")
	err = libplist_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libplist_bin_url := "https://github.com/libimobiledevice/libplist/releases/download/2.6.0/libplist-2.6.0.tar.bz2"
	libplist_cmd_bin := exec.Command("curl", "-L", libplist_bin_url, "-o", "binary.bin")
	err = libplist_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libplist_src_url := "https://github.com/libimobiledevice/libplist/releases/download/2.6.0/libplist-2.6.0.tar.bz2"
	libplist_cmd_src := exec.Command("curl", "-L", libplist_src_url, "-o", "source.tar.gz")
	err = libplist_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libplist_cmd_direct := exec.Command("./binary")
	err = libplist_cmd_direct.Run()
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
}
