package main

// TemplateGlib - GNOME templating library for GLib
// Homepage: https://gitlab.gnome.org/GNOME/template-glib

import (
	"fmt"
	
	"os/exec"
)

func installTemplateGlib() {
	// Método 1: Descargar y extraer .tar.gz
	templateglib_tar_url := "https://download.gnome.org/sources/template-glib/3.36/template-glib-3.36.2.tar.xz"
	templateglib_cmd_tar := exec.Command("curl", "-L", templateglib_tar_url, "-o", "package.tar.gz")
	err := templateglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	templateglib_zip_url := "https://download.gnome.org/sources/template-glib/3.36/template-glib-3.36.2.tar.xz"
	templateglib_cmd_zip := exec.Command("curl", "-L", templateglib_zip_url, "-o", "package.zip")
	err = templateglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	templateglib_bin_url := "https://download.gnome.org/sources/template-glib/3.36/template-glib-3.36.2.tar.xz"
	templateglib_cmd_bin := exec.Command("curl", "-L", templateglib_bin_url, "-o", "binary.bin")
	err = templateglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	templateglib_src_url := "https://download.gnome.org/sources/template-glib/3.36/template-glib-3.36.2.tar.xz"
	templateglib_cmd_src := exec.Command("curl", "-L", templateglib_src_url, "-o", "source.tar.gz")
	err = templateglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	templateglib_cmd_direct := exec.Command("./binary")
	err = templateglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
