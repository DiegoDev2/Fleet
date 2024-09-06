package main

// Libdvdcss - Access DVDs as block devices without the decryption
// Homepage: https://www.videolan.org/developers/libdvdcss.html

import (
	"fmt"
	
	"os/exec"
)

func installLibdvdcss() {
	// Método 1: Descargar y extraer .tar.gz
	libdvdcss_tar_url := "https://download.videolan.org/pub/videolan/libdvdcss/1.4.3/libdvdcss-1.4.3.tar.bz2"
	libdvdcss_cmd_tar := exec.Command("curl", "-L", libdvdcss_tar_url, "-o", "package.tar.gz")
	err := libdvdcss_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdvdcss_zip_url := "https://download.videolan.org/pub/videolan/libdvdcss/1.4.3/libdvdcss-1.4.3.tar.bz2"
	libdvdcss_cmd_zip := exec.Command("curl", "-L", libdvdcss_zip_url, "-o", "package.zip")
	err = libdvdcss_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdvdcss_bin_url := "https://download.videolan.org/pub/videolan/libdvdcss/1.4.3/libdvdcss-1.4.3.tar.bz2"
	libdvdcss_cmd_bin := exec.Command("curl", "-L", libdvdcss_bin_url, "-o", "binary.bin")
	err = libdvdcss_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdvdcss_src_url := "https://download.videolan.org/pub/videolan/libdvdcss/1.4.3/libdvdcss-1.4.3.tar.bz2"
	libdvdcss_cmd_src := exec.Command("curl", "-L", libdvdcss_src_url, "-o", "source.tar.gz")
	err = libdvdcss_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdvdcss_cmd_direct := exec.Command("./binary")
	err = libdvdcss_cmd_direct.Run()
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
}
