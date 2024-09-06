package main

// Libgphoto2 - Gphoto2 digital camera library
// Homepage: http://www.gphoto.org/proj/libgphoto2/

import (
	"fmt"
	
	"os/exec"
)

func installLibgphoto2() {
	// Método 1: Descargar y extraer .tar.gz
	libgphoto2_tar_url := "https://downloads.sourceforge.net/project/gphoto/libgphoto/2.5.31/libgphoto2-2.5.31.tar.bz2"
	libgphoto2_cmd_tar := exec.Command("curl", "-L", libgphoto2_tar_url, "-o", "package.tar.gz")
	err := libgphoto2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgphoto2_zip_url := "https://downloads.sourceforge.net/project/gphoto/libgphoto/2.5.31/libgphoto2-2.5.31.tar.bz2"
	libgphoto2_cmd_zip := exec.Command("curl", "-L", libgphoto2_zip_url, "-o", "package.zip")
	err = libgphoto2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgphoto2_bin_url := "https://downloads.sourceforge.net/project/gphoto/libgphoto/2.5.31/libgphoto2-2.5.31.tar.bz2"
	libgphoto2_cmd_bin := exec.Command("curl", "-L", libgphoto2_bin_url, "-o", "binary.bin")
	err = libgphoto2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgphoto2_src_url := "https://downloads.sourceforge.net/project/gphoto/libgphoto/2.5.31/libgphoto2-2.5.31.tar.bz2"
	libgphoto2_cmd_src := exec.Command("curl", "-L", libgphoto2_src_url, "-o", "source.tar.gz")
	err = libgphoto2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgphoto2_cmd_direct := exec.Command("./binary")
	err = libgphoto2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libexif")
	exec.Command("latte", "install", "libexif").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: libusb-compat")
	exec.Command("latte", "install", "libusb-compat").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
