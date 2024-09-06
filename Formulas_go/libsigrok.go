package main

// Libsigrok - Drivers for logic analyzers and other supported devices
// Homepage: https://sigrok.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibsigrok() {
	// Método 1: Descargar y extraer .tar.gz
	libsigrok_tar_url := "https://sigrok.org/download/source/libsigrok/libsigrok-0.5.2.tar.gz"
	libsigrok_cmd_tar := exec.Command("curl", "-L", libsigrok_tar_url, "-o", "package.tar.gz")
	err := libsigrok_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsigrok_zip_url := "https://sigrok.org/download/source/libsigrok/libsigrok-0.5.2.zip"
	libsigrok_cmd_zip := exec.Command("curl", "-L", libsigrok_zip_url, "-o", "package.zip")
	err = libsigrok_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsigrok_bin_url := "https://sigrok.org/download/source/libsigrok/libsigrok-0.5.2.bin"
	libsigrok_cmd_bin := exec.Command("curl", "-L", libsigrok_bin_url, "-o", "binary.bin")
	err = libsigrok_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsigrok_src_url := "https://sigrok.org/download/source/libsigrok/libsigrok-0.5.2.src.tar.gz"
	libsigrok_cmd_src := exec.Command("curl", "-L", libsigrok_src_url, "-o", "source.tar.gz")
	err = libsigrok_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsigrok_cmd_direct := exec.Command("./binary")
	err = libsigrok_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: sdcc")
exec.Command("latte", "install", "sdcc")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: glibmm@2.66")
exec.Command("latte", "install", "glibmm@2.66")
	fmt.Println("Instalando dependencia: hidapi")
exec.Command("latte", "install", "hidapi")
	fmt.Println("Instalando dependencia: libftdi")
exec.Command("latte", "install", "libftdi")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: nettle")
exec.Command("latte", "install", "nettle")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pygobject3")
exec.Command("latte", "install", "pygobject3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libsigc++@2")
exec.Command("latte", "install", "libsigc++@2")
}
