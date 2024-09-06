package main

// Clutter - Generic high-level canvas library
// Homepage: https://wiki.gnome.org/Projects/Clutter

import (
	"fmt"
	
	"os/exec"
)

func installClutter() {
	// Método 1: Descargar y extraer .tar.gz
	clutter_tar_url := "https://download.gnome.org/sources/clutter/1.26/clutter-1.26.4.tar.xz"
	clutter_cmd_tar := exec.Command("curl", "-L", clutter_tar_url, "-o", "package.tar.gz")
	err := clutter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clutter_zip_url := "https://download.gnome.org/sources/clutter/1.26/clutter-1.26.4.tar.xz"
	clutter_cmd_zip := exec.Command("curl", "-L", clutter_zip_url, "-o", "package.zip")
	err = clutter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clutter_bin_url := "https://download.gnome.org/sources/clutter/1.26/clutter-1.26.4.tar.xz"
	clutter_cmd_bin := exec.Command("curl", "-L", clutter_bin_url, "-o", "binary.bin")
	err = clutter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clutter_src_url := "https://download.gnome.org/sources/clutter/1.26/clutter-1.26.4.tar.xz"
	clutter_cmd_src := exec.Command("curl", "-L", clutter_src_url, "-o", "source.tar.gz")
	err = clutter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clutter_cmd_direct := exec.Command("./binary")
	err = clutter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: cogl")
exec.Command("latte", "install", "cogl")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
