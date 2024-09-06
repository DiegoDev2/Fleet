package main

// Goocanvas - Canvas widget for GTK+ using the Cairo 2D library for drawing
// Homepage: https://wiki.gnome.org/Projects/GooCanvas

import (
	"fmt"
	
	"os/exec"
)

func installGoocanvas() {
	// Método 1: Descargar y extraer .tar.gz
	goocanvas_tar_url := "https://download.gnome.org/sources/goocanvas/3.0/goocanvas-3.0.0.tar.xz"
	goocanvas_cmd_tar := exec.Command("curl", "-L", goocanvas_tar_url, "-o", "package.tar.gz")
	err := goocanvas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goocanvas_zip_url := "https://download.gnome.org/sources/goocanvas/3.0/goocanvas-3.0.0.tar.xz"
	goocanvas_cmd_zip := exec.Command("curl", "-L", goocanvas_zip_url, "-o", "package.zip")
	err = goocanvas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goocanvas_bin_url := "https://download.gnome.org/sources/goocanvas/3.0/goocanvas-3.0.0.tar.xz"
	goocanvas_cmd_bin := exec.Command("curl", "-L", goocanvas_bin_url, "-o", "binary.bin")
	err = goocanvas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goocanvas_src_url := "https://download.gnome.org/sources/goocanvas/3.0/goocanvas-3.0.0.tar.xz"
	goocanvas_cmd_src := exec.Command("curl", "-L", goocanvas_src_url, "-o", "source.tar.gz")
	err = goocanvas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goocanvas_cmd_direct := exec.Command("./binary")
	err = goocanvas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
}
