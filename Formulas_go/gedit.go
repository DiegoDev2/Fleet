package main

// Gedit - GNOME text editor
// Homepage: https://gedit-technology.github.io/apps/gedit/

import (
	"fmt"
	
	"os/exec"
)

func installGedit() {
	// Método 1: Descargar y extraer .tar.gz
	gedit_tar_url := "https://download.gnome.org/sources/gedit/47/gedit-47.0.tar.xz"
	gedit_cmd_tar := exec.Command("curl", "-L", gedit_tar_url, "-o", "package.tar.gz")
	err := gedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gedit_zip_url := "https://download.gnome.org/sources/gedit/47/gedit-47.0.tar.xz"
	gedit_cmd_zip := exec.Command("curl", "-L", gedit_zip_url, "-o", "package.zip")
	err = gedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gedit_bin_url := "https://download.gnome.org/sources/gedit/47/gedit-47.0.tar.xz"
	gedit_cmd_bin := exec.Command("curl", "-L", gedit_bin_url, "-o", "binary.bin")
	err = gedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gedit_src_url := "https://download.gnome.org/sources/gedit/47/gedit-47.0.tar.xz"
	gedit_cmd_src := exec.Command("curl", "-L", gedit_src_url, "-o", "source.tar.gz")
	err = gedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gedit_cmd_direct := exec.Command("./binary")
	err = gedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: desktop-file-utils")
exec.Command("latte", "install", "desktop-file-utils")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
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
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gsettings-desktop-schemas")
exec.Command("latte", "install", "gsettings-desktop-schemas")
	fmt.Println("Instalando dependencia: gspell")
exec.Command("latte", "install", "gspell")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libgedit-amtk")
exec.Command("latte", "install", "libgedit-amtk")
	fmt.Println("Instalando dependencia: libgedit-gfls")
exec.Command("latte", "install", "libgedit-gfls")
	fmt.Println("Instalando dependencia: libgedit-gtksourceview")
exec.Command("latte", "install", "libgedit-gtksourceview")
	fmt.Println("Instalando dependencia: libgedit-tepl")
exec.Command("latte", "install", "libgedit-tepl")
	fmt.Println("Instalando dependencia: libpeas@1")
exec.Command("latte", "install", "libpeas@1")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gtk-mac-integration")
exec.Command("latte", "install", "gtk-mac-integration")
}
