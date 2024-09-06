package main

// Cogl - Low level OpenGL abstraction library developed for Clutter
// Homepage: https://gitlab.gnome.org/GNOME/cogl

import (
	"fmt"
	
	"os/exec"
)

func installCogl() {
	// Método 1: Descargar y extraer .tar.gz
	cogl_tar_url := "https://download.gnome.org/sources/cogl/1.22/cogl-1.22.8.tar.xz"
	cogl_cmd_tar := exec.Command("curl", "-L", cogl_tar_url, "-o", "package.tar.gz")
	err := cogl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cogl_zip_url := "https://download.gnome.org/sources/cogl/1.22/cogl-1.22.8.tar.xz"
	cogl_cmd_zip := exec.Command("curl", "-L", cogl_zip_url, "-o", "package.zip")
	err = cogl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cogl_bin_url := "https://download.gnome.org/sources/cogl/1.22/cogl-1.22.8.tar.xz"
	cogl_cmd_bin := exec.Command("curl", "-L", cogl_bin_url, "-o", "binary.bin")
	err = cogl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cogl_src_url := "https://download.gnome.org/sources/cogl/1.22/cogl-1.22.8.tar.xz"
	cogl_cmd_src := exec.Command("curl", "-L", cogl_src_url, "-o", "source.tar.gz")
	err = cogl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cogl_cmd_direct := exec.Command("./binary")
	err = cogl_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: libxcomposite")
exec.Command("latte", "install", "libxcomposite")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
