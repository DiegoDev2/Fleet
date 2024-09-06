package main

// Gom - GObject wrapper around SQLite
// Homepage: https://wiki.gnome.org/Projects/Gom

import (
	"fmt"
	
	"os/exec"
)

func installGom() {
	// Método 1: Descargar y extraer .tar.gz
	gom_tar_url := "https://download.gnome.org/sources/gom/0.4/gom-0.4.tar.xz"
	gom_cmd_tar := exec.Command("curl", "-L", gom_tar_url, "-o", "package.tar.gz")
	err := gom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gom_zip_url := "https://download.gnome.org/sources/gom/0.4/gom-0.4.tar.xz"
	gom_cmd_zip := exec.Command("curl", "-L", gom_zip_url, "-o", "package.zip")
	err = gom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gom_bin_url := "https://download.gnome.org/sources/gom/0.4/gom-0.4.tar.xz"
	gom_cmd_bin := exec.Command("curl", "-L", gom_bin_url, "-o", "binary.bin")
	err = gom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gom_src_url := "https://download.gnome.org/sources/gom/0.4/gom-0.4.tar.xz"
	gom_cmd_src := exec.Command("curl", "-L", gom_src_url, "-o", "source.tar.gz")
	err = gom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gom_cmd_direct := exec.Command("./binary")
	err = gom_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
