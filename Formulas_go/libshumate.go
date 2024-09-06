package main

// Libshumate - Shumate is a GTK toolkit providing widgets for embedded maps
// Homepage: https://gitlab.gnome.org/GNOME/libshumate

import (
	"fmt"
	
	"os/exec"
)

func installLibshumate() {
	// Método 1: Descargar y extraer .tar.gz
	libshumate_tar_url := "https://download.gnome.org/sources/libshumate/1.2/libshumate-1.2.3.tar.xz"
	libshumate_cmd_tar := exec.Command("curl", "-L", libshumate_tar_url, "-o", "package.tar.gz")
	err := libshumate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libshumate_zip_url := "https://download.gnome.org/sources/libshumate/1.2/libshumate-1.2.3.tar.xz"
	libshumate_cmd_zip := exec.Command("curl", "-L", libshumate_zip_url, "-o", "package.zip")
	err = libshumate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libshumate_bin_url := "https://download.gnome.org/sources/libshumate/1.2/libshumate-1.2.3.tar.xz"
	libshumate_cmd_bin := exec.Command("curl", "-L", libshumate_bin_url, "-o", "binary.bin")
	err = libshumate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libshumate_src_url := "https://download.gnome.org/sources/libshumate/1.2/libshumate-1.2.3.tar.xz"
	libshumate_cmd_src := exec.Command("curl", "-L", libshumate_src_url, "-o", "source.tar.gz")
	err = libshumate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libshumate_cmd_direct := exec.Command("./binary")
	err = libshumate_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: graphene")
exec.Command("latte", "install", "graphene")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: protobuf-c")
exec.Command("latte", "install", "protobuf-c")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
