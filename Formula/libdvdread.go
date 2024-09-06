package main

// Libdvdread - C library for reading DVD-video images
// Homepage: https://www.videolan.org/developers/libdvdnav.html

import (
	"fmt"
	
	"os/exec"
)

func installLibdvdread() {
	// Método 1: Descargar y extraer .tar.gz
	libdvdread_tar_url := "https://download.videolan.org/pub/videolan/libdvdread/6.1.3/libdvdread-6.1.3.tar.bz2"
	libdvdread_cmd_tar := exec.Command("curl", "-L", libdvdread_tar_url, "-o", "package.tar.gz")
	err := libdvdread_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdvdread_zip_url := "https://download.videolan.org/pub/videolan/libdvdread/6.1.3/libdvdread-6.1.3.tar.bz2"
	libdvdread_cmd_zip := exec.Command("curl", "-L", libdvdread_zip_url, "-o", "package.zip")
	err = libdvdread_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdvdread_bin_url := "https://download.videolan.org/pub/videolan/libdvdread/6.1.3/libdvdread-6.1.3.tar.bz2"
	libdvdread_cmd_bin := exec.Command("curl", "-L", libdvdread_bin_url, "-o", "binary.bin")
	err = libdvdread_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdvdread_src_url := "https://download.videolan.org/pub/videolan/libdvdread/6.1.3/libdvdread-6.1.3.tar.bz2"
	libdvdread_cmd_src := exec.Command("curl", "-L", libdvdread_src_url, "-o", "source.tar.gz")
	err = libdvdread_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdvdread_cmd_direct := exec.Command("./binary")
	err = libdvdread_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libdvdcss")
	exec.Command("latte", "install", "libdvdcss").Run()
}
