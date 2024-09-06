package main

// Avahi - Service Discovery for Linux using mDNS/DNS-SD
// Homepage: https://avahi.org

import (
	"fmt"
	
	"os/exec"
)

func installAvahi() {
	// Método 1: Descargar y extraer .tar.gz
	avahi_tar_url := "https://github.com/lathiat/avahi/archive/refs/tags/v0.8.tar.gz"
	avahi_cmd_tar := exec.Command("curl", "-L", avahi_tar_url, "-o", "package.tar.gz")
	err := avahi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	avahi_zip_url := "https://github.com/lathiat/avahi/archive/refs/tags/v0.8.zip"
	avahi_cmd_zip := exec.Command("curl", "-L", avahi_zip_url, "-o", "package.zip")
	err = avahi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	avahi_bin_url := "https://github.com/lathiat/avahi/archive/refs/tags/v0.8.bin"
	avahi_cmd_bin := exec.Command("curl", "-L", avahi_bin_url, "-o", "binary.bin")
	err = avahi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	avahi_src_url := "https://github.com/lathiat/avahi/archive/refs/tags/v0.8.src.tar.gz"
	avahi_cmd_src := exec.Command("curl", "-L", avahi_src_url, "-o", "source.tar.gz")
	err = avahi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	avahi_cmd_direct := exec.Command("./binary")
	err = avahi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: m4")
exec.Command("latte", "install", "m4")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xmltoman")
exec.Command("latte", "install", "xmltoman")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
	fmt.Println("Instalando dependencia: gdbm")
exec.Command("latte", "install", "gdbm")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libdaemon")
exec.Command("latte", "install", "libdaemon")
}
