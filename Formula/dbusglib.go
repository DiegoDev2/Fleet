package main

// DbusGlib - GLib bindings for the D-Bus message bus system
// Homepage: https://wiki.freedesktop.org/www/Software/DBusBindings/

import (
	"fmt"
	
	"os/exec"
)

func installDbusGlib() {
	// Método 1: Descargar y extraer .tar.gz
	dbusglib_tar_url := "https://dbus.freedesktop.org/releases/dbus-glib/dbus-glib-0.112.tar.gz"
	dbusglib_cmd_tar := exec.Command("curl", "-L", dbusglib_tar_url, "-o", "package.tar.gz")
	err := dbusglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbusglib_zip_url := "https://dbus.freedesktop.org/releases/dbus-glib/dbus-glib-0.112.zip"
	dbusglib_cmd_zip := exec.Command("curl", "-L", dbusglib_zip_url, "-o", "package.zip")
	err = dbusglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbusglib_bin_url := "https://dbus.freedesktop.org/releases/dbus-glib/dbus-glib-0.112.bin"
	dbusglib_cmd_bin := exec.Command("curl", "-L", dbusglib_bin_url, "-o", "binary.bin")
	err = dbusglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbusglib_src_url := "https://dbus.freedesktop.org/releases/dbus-glib/dbus-glib-0.112.src.tar.gz"
	dbusglib_cmd_src := exec.Command("curl", "-L", dbusglib_src_url, "-o", "source.tar.gz")
	err = dbusglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbusglib_cmd_direct := exec.Command("./binary")
	err = dbusglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
