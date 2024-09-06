package main

// Libgdata - GLib-based library for accessing online service APIs
// Homepage: https://wiki.gnome.org/Projects/libgdata

import (
	"fmt"
	
	"os/exec"
)

func installLibgdata() {
	// Método 1: Descargar y extraer .tar.gz
	libgdata_tar_url := "https://download.gnome.org/sources/libgdata/0.18/libgdata-0.18.1.tar.xz"
	libgdata_cmd_tar := exec.Command("curl", "-L", libgdata_tar_url, "-o", "package.tar.gz")
	err := libgdata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgdata_zip_url := "https://download.gnome.org/sources/libgdata/0.18/libgdata-0.18.1.tar.xz"
	libgdata_cmd_zip := exec.Command("curl", "-L", libgdata_zip_url, "-o", "package.zip")
	err = libgdata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgdata_bin_url := "https://download.gnome.org/sources/libgdata/0.18/libgdata-0.18.1.tar.xz"
	libgdata_cmd_bin := exec.Command("curl", "-L", libgdata_bin_url, "-o", "binary.bin")
	err = libgdata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgdata_src_url := "https://download.gnome.org/sources/libgdata/0.18/libgdata-0.18.1.tar.xz"
	libgdata_cmd_src := exec.Command("curl", "-L", libgdata_src_url, "-o", "source.tar.gz")
	err = libgdata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgdata_cmd_direct := exec.Command("./binary")
	err = libgdata_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: liboauth")
exec.Command("latte", "install", "liboauth")
	fmt.Println("Instalando dependencia: libsoup@2")
exec.Command("latte", "install", "libsoup@2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
