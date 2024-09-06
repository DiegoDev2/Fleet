package main

// AtSpi2Core - Protocol definitions and daemon for D-Bus at-spi
// Homepage: https://www.freedesktop.org/wiki/Accessibility/AT-SPI2

import (
	"fmt"
	
	"os/exec"
)

func installAtSpi2Core() {
	// Método 1: Descargar y extraer .tar.gz
	atspi2core_tar_url := "https://download.gnome.org/sources/at-spi2-core/2.52/at-spi2-core-2.52.0.tar.xz"
	atspi2core_cmd_tar := exec.Command("curl", "-L", atspi2core_tar_url, "-o", "package.tar.gz")
	err := atspi2core_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atspi2core_zip_url := "https://download.gnome.org/sources/at-spi2-core/2.52/at-spi2-core-2.52.0.tar.xz"
	atspi2core_cmd_zip := exec.Command("curl", "-L", atspi2core_zip_url, "-o", "package.zip")
	err = atspi2core_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atspi2core_bin_url := "https://download.gnome.org/sources/at-spi2-core/2.52/at-spi2-core-2.52.0.tar.xz"
	atspi2core_cmd_bin := exec.Command("curl", "-L", atspi2core_bin_url, "-o", "binary.bin")
	err = atspi2core_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atspi2core_src_url := "https://download.gnome.org/sources/at-spi2-core/2.52/at-spi2-core-2.52.0.tar.xz"
	atspi2core_cmd_src := exec.Command("curl", "-L", atspi2core_src_url, "-o", "source.tar.gz")
	err = atspi2core_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atspi2core_cmd_direct := exec.Command("./binary")
	err = atspi2core_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxtst")
exec.Command("latte", "install", "libxtst")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
