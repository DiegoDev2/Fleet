package main

// Libexif - EXIF parsing library
// Homepage: https://libexif.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibexif() {
	// Método 1: Descargar y extraer .tar.gz
	libexif_tar_url := "https://github.com/libexif/libexif/releases/download/v0.6.24/libexif-0.6.24.tar.bz2"
	libexif_cmd_tar := exec.Command("curl", "-L", libexif_tar_url, "-o", "package.tar.gz")
	err := libexif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libexif_zip_url := "https://github.com/libexif/libexif/releases/download/v0.6.24/libexif-0.6.24.tar.bz2"
	libexif_cmd_zip := exec.Command("curl", "-L", libexif_zip_url, "-o", "package.zip")
	err = libexif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libexif_bin_url := "https://github.com/libexif/libexif/releases/download/v0.6.24/libexif-0.6.24.tar.bz2"
	libexif_cmd_bin := exec.Command("curl", "-L", libexif_bin_url, "-o", "binary.bin")
	err = libexif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libexif_src_url := "https://github.com/libexif/libexif/releases/download/v0.6.24/libexif-0.6.24.tar.bz2"
	libexif_cmd_src := exec.Command("curl", "-L", libexif_src_url, "-o", "source.tar.gz")
	err = libexif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libexif_cmd_direct := exec.Command("./binary")
	err = libexif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
