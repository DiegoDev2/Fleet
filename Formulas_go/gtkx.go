package main

// Gtkx - GUI toolkit
// Homepage: https://gtk.org/

import (
	"fmt"
	
	"os/exec"
)

func installGtkx() {
	// Método 1: Descargar y extraer .tar.gz
	gtkx_tar_url := "https://download.gnome.org/sources/gtk+/2.24/gtk+-2.24.33.tar.xz"
	gtkx_cmd_tar := exec.Command("curl", "-L", gtkx_tar_url, "-o", "package.tar.gz")
	err := gtkx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkx_zip_url := "https://download.gnome.org/sources/gtk+/2.24/gtk+-2.24.33.tar.xz"
	gtkx_cmd_zip := exec.Command("curl", "-L", gtkx_zip_url, "-o", "package.zip")
	err = gtkx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkx_bin_url := "https://download.gnome.org/sources/gtk+/2.24/gtk+-2.24.33.tar.xz"
	gtkx_cmd_bin := exec.Command("curl", "-L", gtkx_bin_url, "-o", "binary.bin")
	err = gtkx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkx_src_url := "https://download.gnome.org/sources/gtk+/2.24/gtk+-2.24.33.tar.xz"
	gtkx_cmd_src := exec.Command("curl", "-L", gtkx_src_url, "-o", "source.tar.gz")
	err = gtkx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkx_cmd_direct := exec.Command("./binary")
	err = gtkx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
exec.Command("latte", "install", "hicolor-icon-theme")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcomposite")
exec.Command("latte", "install", "libxcomposite")
	fmt.Println("Instalando dependencia: libxcursor")
exec.Command("latte", "install", "libxcursor")
	fmt.Println("Instalando dependencia: libxdamage")
exec.Command("latte", "install", "libxdamage")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxrender")
exec.Command("latte", "install", "libxrender")
}
