package main

// ClutterGtk - GTK+ integration library for Clutter
// Homepage: https://wiki.gnome.org/Projects/Clutter

import (
	"fmt"
	
	"os/exec"
)

func installClutterGtk() {
	// Método 1: Descargar y extraer .tar.gz
	cluttergtk_tar_url := "https://download.gnome.org/sources/clutter-gtk/1.8/clutter-gtk-1.8.4.tar.xz"
	cluttergtk_cmd_tar := exec.Command("curl", "-L", cluttergtk_tar_url, "-o", "package.tar.gz")
	err := cluttergtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cluttergtk_zip_url := "https://download.gnome.org/sources/clutter-gtk/1.8/clutter-gtk-1.8.4.tar.xz"
	cluttergtk_cmd_zip := exec.Command("curl", "-L", cluttergtk_zip_url, "-o", "package.zip")
	err = cluttergtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cluttergtk_bin_url := "https://download.gnome.org/sources/clutter-gtk/1.8/clutter-gtk-1.8.4.tar.xz"
	cluttergtk_cmd_bin := exec.Command("curl", "-L", cluttergtk_bin_url, "-o", "binary.bin")
	err = cluttergtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cluttergtk_src_url := "https://download.gnome.org/sources/clutter-gtk/1.8/clutter-gtk-1.8.4.tar.xz"
	cluttergtk_cmd_src := exec.Command("curl", "-L", cluttergtk_src_url, "-o", "source.tar.gz")
	err = cluttergtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cluttergtk_cmd_direct := exec.Command("./binary")
	err = cluttergtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: clutter")
	exec.Command("latte", "install", "clutter").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
}
