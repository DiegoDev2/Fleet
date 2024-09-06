package main

// Gjs - JavaScript Bindings for GNOME
// Homepage: https://gitlab.gnome.org/GNOME/gjs/wikis/Home

import (
	"fmt"
	
	"os/exec"
)

func installGjs() {
	// Método 1: Descargar y extraer .tar.gz
	gjs_tar_url := "https://download.gnome.org/sources/gjs/1.80/gjs-1.80.2.tar.xz"
	gjs_cmd_tar := exec.Command("curl", "-L", gjs_tar_url, "-o", "package.tar.gz")
	err := gjs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gjs_zip_url := "https://download.gnome.org/sources/gjs/1.80/gjs-1.80.2.tar.xz"
	gjs_cmd_zip := exec.Command("curl", "-L", gjs_zip_url, "-o", "package.zip")
	err = gjs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gjs_bin_url := "https://download.gnome.org/sources/gjs/1.80/gjs-1.80.2.tar.xz"
	gjs_cmd_bin := exec.Command("curl", "-L", gjs_bin_url, "-o", "binary.bin")
	err = gjs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gjs_src_url := "https://download.gnome.org/sources/gjs/1.80/gjs-1.80.2.tar.xz"
	gjs_cmd_src := exec.Command("curl", "-L", gjs_src_url, "-o", "source.tar.gz")
	err = gjs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gjs_cmd_direct := exec.Command("./binary")
	err = gjs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: spidermonkey@115")
exec.Command("latte", "install", "spidermonkey@115")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
