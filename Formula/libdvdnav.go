package main

// Libdvdnav - DVD navigation library
// Homepage: https://www.videolan.org/developers/libdvdnav.html

import (
	"fmt"
	
	"os/exec"
)

func installLibdvdnav() {
	// Método 1: Descargar y extraer .tar.gz
	libdvdnav_tar_url := "https://download.videolan.org/pub/videolan/libdvdnav/6.1.1/libdvdnav-6.1.1.tar.bz2"
	libdvdnav_cmd_tar := exec.Command("curl", "-L", libdvdnav_tar_url, "-o", "package.tar.gz")
	err := libdvdnav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdvdnav_zip_url := "https://download.videolan.org/pub/videolan/libdvdnav/6.1.1/libdvdnav-6.1.1.tar.bz2"
	libdvdnav_cmd_zip := exec.Command("curl", "-L", libdvdnav_zip_url, "-o", "package.zip")
	err = libdvdnav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdvdnav_bin_url := "https://download.videolan.org/pub/videolan/libdvdnav/6.1.1/libdvdnav-6.1.1.tar.bz2"
	libdvdnav_cmd_bin := exec.Command("curl", "-L", libdvdnav_bin_url, "-o", "binary.bin")
	err = libdvdnav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdvdnav_src_url := "https://download.videolan.org/pub/videolan/libdvdnav/6.1.1/libdvdnav-6.1.1.tar.bz2"
	libdvdnav_cmd_src := exec.Command("curl", "-L", libdvdnav_src_url, "-o", "source.tar.gz")
	err = libdvdnav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdvdnav_cmd_direct := exec.Command("./binary")
	err = libdvdnav_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libdvdread")
	exec.Command("latte", "install", "libdvdread").Run()
}
