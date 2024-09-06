package main

// Gtkglext - OpenGL extension to GTK+
// Homepage: https://gitlab.gnome.org/Archive/gtkglext

import (
	"fmt"
	
	"os/exec"
)

func installGtkglext() {
	// Método 1: Descargar y extraer .tar.gz
	gtkglext_tar_url := "https://download.gnome.org/sources/gtkglext/1.2/gtkglext-1.2.0.tar.gz"
	gtkglext_cmd_tar := exec.Command("curl", "-L", gtkglext_tar_url, "-o", "package.tar.gz")
	err := gtkglext_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkglext_zip_url := "https://download.gnome.org/sources/gtkglext/1.2/gtkglext-1.2.0.zip"
	gtkglext_cmd_zip := exec.Command("curl", "-L", gtkglext_zip_url, "-o", "package.zip")
	err = gtkglext_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkglext_bin_url := "https://download.gnome.org/sources/gtkglext/1.2/gtkglext-1.2.0.bin"
	gtkglext_cmd_bin := exec.Command("curl", "-L", gtkglext_bin_url, "-o", "binary.bin")
	err = gtkglext_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkglext_src_url := "https://download.gnome.org/sources/gtkglext/1.2/gtkglext-1.2.0.src.tar.gz"
	gtkglext_cmd_src := exec.Command("curl", "-L", gtkglext_src_url, "-o", "source.tar.gz")
	err = gtkglext_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkglext_cmd_direct := exec.Command("./binary")
	err = gtkglext_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
