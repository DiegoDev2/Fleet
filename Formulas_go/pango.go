package main

// Pango - Framework for layout and rendering of i18n text
// Homepage: https://pango.gnome.org

import (
	"fmt"
	
	"os/exec"
)

func installPango() {
	// Método 1: Descargar y extraer .tar.gz
	pango_tar_url := "https://download.gnome.org/sources/pango/1.54/pango-1.54.0.tar.xz"
	pango_cmd_tar := exec.Command("curl", "-L", pango_tar_url, "-o", "package.tar.gz")
	err := pango_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pango_zip_url := "https://download.gnome.org/sources/pango/1.54/pango-1.54.0.tar.xz"
	pango_cmd_zip := exec.Command("curl", "-L", pango_zip_url, "-o", "package.zip")
	err = pango_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pango_bin_url := "https://download.gnome.org/sources/pango/1.54/pango-1.54.0.tar.xz"
	pango_cmd_bin := exec.Command("curl", "-L", pango_bin_url, "-o", "binary.bin")
	err = pango_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pango_src_url := "https://download.gnome.org/sources/pango/1.54/pango-1.54.0.tar.xz"
	pango_cmd_src := exec.Command("curl", "-L", pango_src_url, "-o", "source.tar.gz")
	err = pango_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pango_cmd_direct := exec.Command("./binary")
	err = pango_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: fribidi")
exec.Command("latte", "install", "fribidi")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
