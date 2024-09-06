package main

// Librsvg - Library to render SVG files using Cairo
// Homepage: https://wiki.gnome.org/Projects/LibRsvg

import (
	"fmt"
	
	"os/exec"
)

func installLibrsvg() {
	// Método 1: Descargar y extraer .tar.gz
	librsvg_tar_url := "https://download.gnome.org/sources/librsvg/2.58/librsvg-2.58.3.tar.xz"
	librsvg_cmd_tar := exec.Command("curl", "-L", librsvg_tar_url, "-o", "package.tar.gz")
	err := librsvg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librsvg_zip_url := "https://download.gnome.org/sources/librsvg/2.58/librsvg-2.58.3.tar.xz"
	librsvg_cmd_zip := exec.Command("curl", "-L", librsvg_zip_url, "-o", "package.zip")
	err = librsvg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librsvg_bin_url := "https://download.gnome.org/sources/librsvg/2.58/librsvg-2.58.3.tar.xz"
	librsvg_cmd_bin := exec.Command("curl", "-L", librsvg_bin_url, "-o", "binary.bin")
	err = librsvg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librsvg_src_url := "https://download.gnome.org/sources/librsvg/2.58/librsvg-2.58.3.tar.xz"
	librsvg_cmd_src := exec.Command("curl", "-L", librsvg_src_url, "-o", "source.tar.gz")
	err = librsvg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librsvg_cmd_direct := exec.Command("./binary")
	err = librsvg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
