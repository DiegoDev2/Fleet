package main

// Dbus - Message bus system, providing inter-application communication
// Homepage: https://wiki.freedesktop.org/www/Software/dbus

import (
	"fmt"
	
	"os/exec"
)

func installDbus() {
	// Método 1: Descargar y extraer .tar.gz
	dbus_tar_url := "https://dbus.freedesktop.org/releases/dbus/dbus-1.14.10.tar.xz"
	dbus_cmd_tar := exec.Command("curl", "-L", dbus_tar_url, "-o", "package.tar.gz")
	err := dbus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbus_zip_url := "https://dbus.freedesktop.org/releases/dbus/dbus-1.14.10.tar.xz"
	dbus_cmd_zip := exec.Command("curl", "-L", dbus_zip_url, "-o", "package.zip")
	err = dbus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbus_bin_url := "https://dbus.freedesktop.org/releases/dbus/dbus-1.14.10.tar.xz"
	dbus_cmd_bin := exec.Command("curl", "-L", dbus_bin_url, "-o", "binary.bin")
	err = dbus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbus_src_url := "https://dbus.freedesktop.org/releases/dbus/dbus-1.14.10.tar.xz"
	dbus_cmd_src := exec.Command("curl", "-L", dbus_src_url, "-o", "source.tar.gz")
	err = dbus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbus_cmd_direct := exec.Command("./binary")
	err = dbus_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xmlto")
exec.Command("latte", "install", "xmlto")
}
