package main

// Libdca - Library for decoding DTS Coherent Acoustics streams
// Homepage: https://www.videolan.org/developers/libdca.html

import (
	"fmt"
	
	"os/exec"
)

func installLibdca() {
	// Método 1: Descargar y extraer .tar.gz
	libdca_tar_url := "https://download.videolan.org/pub/videolan/libdca/0.0.7/libdca-0.0.7.tar.bz2"
	libdca_cmd_tar := exec.Command("curl", "-L", libdca_tar_url, "-o", "package.tar.gz")
	err := libdca_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdca_zip_url := "https://download.videolan.org/pub/videolan/libdca/0.0.7/libdca-0.0.7.tar.bz2"
	libdca_cmd_zip := exec.Command("curl", "-L", libdca_zip_url, "-o", "package.zip")
	err = libdca_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdca_bin_url := "https://download.videolan.org/pub/videolan/libdca/0.0.7/libdca-0.0.7.tar.bz2"
	libdca_cmd_bin := exec.Command("curl", "-L", libdca_bin_url, "-o", "binary.bin")
	err = libdca_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdca_src_url := "https://download.videolan.org/pub/videolan/libdca/0.0.7/libdca-0.0.7.tar.bz2"
	libdca_cmd_src := exec.Command("curl", "-L", libdca_src_url, "-o", "source.tar.gz")
	err = libdca_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdca_cmd_direct := exec.Command("./binary")
	err = libdca_cmd_direct.Run()
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
