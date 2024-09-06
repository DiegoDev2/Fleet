package main

// JsonGlib - Library for JSON, based on GLib
// Homepage: https://wiki.gnome.org/Projects/JsonGlib

import (
	"fmt"
	
	"os/exec"
)

func installJsonGlib() {
	// Método 1: Descargar y extraer .tar.gz
	jsonglib_tar_url := "https://download.gnome.org/sources/json-glib/1.8/json-glib-1.8.0.tar.xz"
	jsonglib_cmd_tar := exec.Command("curl", "-L", jsonglib_tar_url, "-o", "package.tar.gz")
	err := jsonglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jsonglib_zip_url := "https://download.gnome.org/sources/json-glib/1.8/json-glib-1.8.0.tar.xz"
	jsonglib_cmd_zip := exec.Command("curl", "-L", jsonglib_zip_url, "-o", "package.zip")
	err = jsonglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jsonglib_bin_url := "https://download.gnome.org/sources/json-glib/1.8/json-glib-1.8.0.tar.xz"
	jsonglib_cmd_bin := exec.Command("curl", "-L", jsonglib_bin_url, "-o", "binary.bin")
	err = jsonglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jsonglib_src_url := "https://download.gnome.org/sources/json-glib/1.8/json-glib-1.8.0.tar.xz"
	jsonglib_cmd_src := exec.Command("curl", "-L", jsonglib_src_url, "-o", "source.tar.gz")
	err = jsonglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jsonglib_cmd_direct := exec.Command("./binary")
	err = jsonglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
