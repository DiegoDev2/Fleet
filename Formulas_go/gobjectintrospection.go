package main

// GobjectIntrospection - Generate introspection data for GObject libraries
// Homepage: https://gi.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installGobjectIntrospection() {
	// Método 1: Descargar y extraer .tar.gz
	gobjectintrospection_tar_url := "https://download.gnome.org/sources/gobject-introspection/1.80/gobject-introspection-1.80.1.tar.xz"
	gobjectintrospection_cmd_tar := exec.Command("curl", "-L", gobjectintrospection_tar_url, "-o", "package.tar.gz")
	err := gobjectintrospection_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gobjectintrospection_zip_url := "https://download.gnome.org/sources/gobject-introspection/1.80/gobject-introspection-1.80.1.tar.xz"
	gobjectintrospection_cmd_zip := exec.Command("curl", "-L", gobjectintrospection_zip_url, "-o", "package.zip")
	err = gobjectintrospection_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gobjectintrospection_bin_url := "https://download.gnome.org/sources/gobject-introspection/1.80/gobject-introspection-1.80.1.tar.xz"
	gobjectintrospection_cmd_bin := exec.Command("curl", "-L", gobjectintrospection_bin_url, "-o", "binary.bin")
	err = gobjectintrospection_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gobjectintrospection_src_url := "https://download.gnome.org/sources/gobject-introspection/1.80/gobject-introspection-1.80.1.tar.xz"
	gobjectintrospection_cmd_src := exec.Command("curl", "-L", gobjectintrospection_src_url, "-o", "source.tar.gz")
	err = gobjectintrospection_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gobjectintrospection_cmd_direct := exec.Command("./binary")
	err = gobjectintrospection_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
