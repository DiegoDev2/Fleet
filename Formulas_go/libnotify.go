package main

// Libnotify - Library that sends desktop notifications to a notification daemon
// Homepage: https://gitlab.gnome.org/GNOME/libnotify

import (
	"fmt"
	
	"os/exec"
)

func installLibnotify() {
	// Método 1: Descargar y extraer .tar.gz
	libnotify_tar_url := "https://download.gnome.org/sources/libnotify/0.8/libnotify-0.8.3.tar.xz"
	libnotify_cmd_tar := exec.Command("curl", "-L", libnotify_tar_url, "-o", "package.tar.gz")
	err := libnotify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnotify_zip_url := "https://download.gnome.org/sources/libnotify/0.8/libnotify-0.8.3.tar.xz"
	libnotify_cmd_zip := exec.Command("curl", "-L", libnotify_zip_url, "-o", "package.zip")
	err = libnotify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnotify_bin_url := "https://download.gnome.org/sources/libnotify/0.8/libnotify-0.8.3.tar.xz"
	libnotify_cmd_bin := exec.Command("curl", "-L", libnotify_bin_url, "-o", "binary.bin")
	err = libnotify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnotify_src_url := "https://download.gnome.org/sources/libnotify/0.8/libnotify-0.8.3.tar.xz"
	libnotify_cmd_src := exec.Command("curl", "-L", libnotify_src_url, "-o", "source.tar.gz")
	err = libnotify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnotify_cmd_direct := exec.Command("./binary")
	err = libnotify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
