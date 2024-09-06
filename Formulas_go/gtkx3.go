package main

// Gtkx3 - Toolkit for creating graphical user interfaces
// Homepage: https://gtk.org/

import (
	"fmt"
	
	"os/exec"
)

func installGtkx3() {
	// Método 1: Descargar y extraer .tar.gz
	gtkx3_tar_url := "https://download.gnome.org/sources/gtk+/3.24/gtk+-3.24.43.tar.xz"
	gtkx3_cmd_tar := exec.Command("curl", "-L", gtkx3_tar_url, "-o", "package.tar.gz")
	err := gtkx3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkx3_zip_url := "https://download.gnome.org/sources/gtk+/3.24/gtk+-3.24.43.tar.xz"
	gtkx3_cmd_zip := exec.Command("curl", "-L", gtkx3_zip_url, "-o", "package.zip")
	err = gtkx3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkx3_bin_url := "https://download.gnome.org/sources/gtk+/3.24/gtk+-3.24.43.tar.xz"
	gtkx3_cmd_bin := exec.Command("curl", "-L", gtkx3_bin_url, "-o", "binary.bin")
	err = gtkx3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkx3_src_url := "https://download.gnome.org/sources/gtk+/3.24/gtk+-3.24.43.tar.xz"
	gtkx3_cmd_src := exec.Command("curl", "-L", gtkx3_src_url, "-o", "source.tar.gz")
	err = gtkx3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkx3_cmd_direct := exec.Command("./binary")
	err = gtkx3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
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
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: fribidi")
exec.Command("latte", "install", "fribidi")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gsettings-desktop-schemas")
exec.Command("latte", "install", "gsettings-desktop-schemas")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
exec.Command("latte", "install", "hicolor-icon-theme")
	fmt.Println("Instalando dependencia: libepoxy")
exec.Command("latte", "install", "libepoxy")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: iso-codes")
exec.Command("latte", "install", "iso-codes")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxdamage")
exec.Command("latte", "install", "libxdamage")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxkbcommon")
exec.Command("latte", "install", "libxkbcommon")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
	fmt.Println("Instalando dependencia: wayland-protocols")
exec.Command("latte", "install", "wayland-protocols")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
