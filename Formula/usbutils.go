package main

// Usbutils - List detailed info about USB devices
// Homepage: http://www.linux-usb.org/

import (
	"fmt"
	
	"os/exec"
)

func installUsbutils() {
	// Método 1: Descargar y extraer .tar.gz
	usbutils_tar_url := "https://mirrors.edge.kernel.org/pub/linux/utils/usb/usbutils/usbutils-017.tar.gz"
	usbutils_cmd_tar := exec.Command("curl", "-L", usbutils_tar_url, "-o", "package.tar.gz")
	err := usbutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	usbutils_zip_url := "https://mirrors.edge.kernel.org/pub/linux/utils/usb/usbutils/usbutils-017.zip"
	usbutils_cmd_zip := exec.Command("curl", "-L", usbutils_zip_url, "-o", "package.zip")
	err = usbutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	usbutils_bin_url := "https://mirrors.edge.kernel.org/pub/linux/utils/usb/usbutils/usbutils-017.bin"
	usbutils_cmd_bin := exec.Command("curl", "-L", usbutils_bin_url, "-o", "binary.bin")
	err = usbutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	usbutils_src_url := "https://mirrors.edge.kernel.org/pub/linux/utils/usb/usbutils/usbutils-017.src.tar.gz"
	usbutils_cmd_src := exec.Command("curl", "-L", usbutils_src_url, "-o", "source.tar.gz")
	err = usbutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	usbutils_cmd_direct := exec.Command("./binary")
	err = usbutils_cmd_direct.Run()
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
