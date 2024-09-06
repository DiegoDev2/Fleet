package main

// Libgusb - GObject wrappers for libusb1
// Homepage: https://github.com/hughsie/libgusb

import (
	"fmt"
	
	"os/exec"
)

func installLibgusb() {
	// Método 1: Descargar y extraer .tar.gz
	libgusb_tar_url := "https://github.com/hughsie/libgusb/archive/refs/tags/0.4.9.tar.gz"
	libgusb_cmd_tar := exec.Command("curl", "-L", libgusb_tar_url, "-o", "package.tar.gz")
	err := libgusb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgusb_zip_url := "https://github.com/hughsie/libgusb/archive/refs/tags/0.4.9.zip"
	libgusb_cmd_zip := exec.Command("curl", "-L", libgusb_zip_url, "-o", "package.zip")
	err = libgusb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgusb_bin_url := "https://github.com/hughsie/libgusb/archive/refs/tags/0.4.9.bin"
	libgusb_cmd_bin := exec.Command("curl", "-L", libgusb_bin_url, "-o", "binary.bin")
	err = libgusb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgusb_src_url := "https://github.com/hughsie/libgusb/archive/refs/tags/0.4.9.src.tar.gz"
	libgusb_cmd_src := exec.Command("curl", "-L", libgusb_src_url, "-o", "source.tar.gz")
	err = libgusb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgusb_cmd_direct := exec.Command("./binary")
	err = libgusb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: usb.ids")
	exec.Command("latte", "install", "usb.ids").Run()
}
