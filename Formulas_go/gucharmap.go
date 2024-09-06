package main

// Gucharmap - GNOME Character Map, based on the Unicode Character Database
// Homepage: https://wiki.gnome.org/Apps/Gucharmap

import (
	"fmt"
	
	"os/exec"
)

func installGucharmap() {
	// Método 1: Descargar y extraer .tar.gz
	gucharmap_tar_url := "https://gitlab.gnome.org/GNOME/gucharmap/-/archive/15.1.5/gucharmap-15.1.5.tar.bz2"
	gucharmap_cmd_tar := exec.Command("curl", "-L", gucharmap_tar_url, "-o", "package.tar.gz")
	err := gucharmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gucharmap_zip_url := "https://gitlab.gnome.org/GNOME/gucharmap/-/archive/15.1.5/gucharmap-15.1.5.tar.bz2"
	gucharmap_cmd_zip := exec.Command("curl", "-L", gucharmap_zip_url, "-o", "package.zip")
	err = gucharmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gucharmap_bin_url := "https://gitlab.gnome.org/GNOME/gucharmap/-/archive/15.1.5/gucharmap-15.1.5.tar.bz2"
	gucharmap_cmd_bin := exec.Command("curl", "-L", gucharmap_bin_url, "-o", "binary.bin")
	err = gucharmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gucharmap_src_url := "https://gitlab.gnome.org/GNOME/gucharmap/-/archive/15.1.5/gucharmap-15.1.5.tar.bz2"
	gucharmap_cmd_src := exec.Command("curl", "-L", gucharmap_src_url, "-o", "source.tar.gz")
	err = gucharmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gucharmap_cmd_direct := exec.Command("./binary")
	err = gucharmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: desktop-file-utils")
exec.Command("latte", "install", "desktop-file-utils")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: itstool")
exec.Command("latte", "install", "itstool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
