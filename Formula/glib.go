package main

// Glib - Core application library for C
// Homepage: https://docs.gtk.org/glib/

import (
	"fmt"
	
	"os/exec"
)

func installGlib() {
	// Método 1: Descargar y extraer .tar.gz
	glib_tar_url := "https://download.gnome.org/sources/glib/2.80/glib-2.80.4.tar.xz"
	glib_cmd_tar := exec.Command("curl", "-L", glib_tar_url, "-o", "package.tar.gz")
	err := glib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glib_zip_url := "https://download.gnome.org/sources/glib/2.80/glib-2.80.4.tar.xz"
	glib_cmd_zip := exec.Command("curl", "-L", glib_zip_url, "-o", "package.zip")
	err = glib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glib_bin_url := "https://download.gnome.org/sources/glib/2.80/glib-2.80.4.tar.xz"
	glib_cmd_bin := exec.Command("curl", "-L", glib_bin_url, "-o", "binary.bin")
	err = glib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glib_src_url := "https://download.gnome.org/sources/glib/2.80/glib-2.80.4.tar.xz"
	glib_cmd_src := exec.Command("curl", "-L", glib_src_url, "-o", "source.tar.gz")
	err = glib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glib_cmd_direct := exec.Command("./binary")
	err = glib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
