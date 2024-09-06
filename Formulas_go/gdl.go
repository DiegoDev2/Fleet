package main

// Gdl - GNOME Docking Library provides docking features for GTK+ 3
// Homepage: https://gitlab.gnome.org/GNOME/gdl

import (
	"fmt"
	
	"os/exec"
)

func installGdl() {
	// Método 1: Descargar y extraer .tar.gz
	gdl_tar_url := "https://download.gnome.org/sources/gdl/3.40/gdl-3.40.0.tar.xz"
	gdl_cmd_tar := exec.Command("curl", "-L", gdl_tar_url, "-o", "package.tar.gz")
	err := gdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdl_zip_url := "https://download.gnome.org/sources/gdl/3.40/gdl-3.40.0.tar.xz"
	gdl_cmd_zip := exec.Command("curl", "-L", gdl_zip_url, "-o", "package.zip")
	err = gdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdl_bin_url := "https://download.gnome.org/sources/gdl/3.40/gdl-3.40.0.tar.xz"
	gdl_cmd_bin := exec.Command("curl", "-L", gdl_bin_url, "-o", "binary.bin")
	err = gdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdl_src_url := "https://download.gnome.org/sources/gdl/3.40/gdl-3.40.0.tar.xz"
	gdl_cmd_src := exec.Command("curl", "-L", gdl_src_url, "-o", "source.tar.gz")
	err = gdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdl_cmd_direct := exec.Command("./binary")
	err = gdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
