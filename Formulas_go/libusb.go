package main

// Libusb - Library for USB device access
// Homepage: https://libusb.info/

import (
	"fmt"
	
	"os/exec"
)

func installLibusb() {
	// Método 1: Descargar y extraer .tar.gz
	libusb_tar_url := "https://github.com/libusb/libusb/releases/download/v1.0.27/libusb-1.0.27.tar.bz2"
	libusb_cmd_tar := exec.Command("curl", "-L", libusb_tar_url, "-o", "package.tar.gz")
	err := libusb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libusb_zip_url := "https://github.com/libusb/libusb/releases/download/v1.0.27/libusb-1.0.27.tar.bz2"
	libusb_cmd_zip := exec.Command("curl", "-L", libusb_zip_url, "-o", "package.zip")
	err = libusb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libusb_bin_url := "https://github.com/libusb/libusb/releases/download/v1.0.27/libusb-1.0.27.tar.bz2"
	libusb_cmd_bin := exec.Command("curl", "-L", libusb_bin_url, "-o", "binary.bin")
	err = libusb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libusb_src_url := "https://github.com/libusb/libusb/releases/download/v1.0.27/libusb-1.0.27.tar.bz2"
	libusb_cmd_src := exec.Command("curl", "-L", libusb_src_url, "-o", "source.tar.gz")
	err = libusb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libusb_cmd_direct := exec.Command("./binary")
	err = libusb_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
