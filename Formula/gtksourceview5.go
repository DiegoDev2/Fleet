package main

// Gtksourceview5 - Text view with syntax, undo/redo, and text marks
// Homepage: https://projects.gnome.org/gtksourceview/

import (
	"fmt"
	
	"os/exec"
)

func installGtksourceview5() {
	// Método 1: Descargar y extraer .tar.gz
	gtksourceview5_tar_url := "https://download.gnome.org/sources/gtksourceview/5.12/gtksourceview-5.12.1.tar.xz"
	gtksourceview5_cmd_tar := exec.Command("curl", "-L", gtksourceview5_tar_url, "-o", "package.tar.gz")
	err := gtksourceview5_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtksourceview5_zip_url := "https://download.gnome.org/sources/gtksourceview/5.12/gtksourceview-5.12.1.tar.xz"
	gtksourceview5_cmd_zip := exec.Command("curl", "-L", gtksourceview5_zip_url, "-o", "package.zip")
	err = gtksourceview5_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtksourceview5_bin_url := "https://download.gnome.org/sources/gtksourceview/5.12/gtksourceview-5.12.1.tar.xz"
	gtksourceview5_cmd_bin := exec.Command("curl", "-L", gtksourceview5_bin_url, "-o", "binary.bin")
	err = gtksourceview5_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtksourceview5_src_url := "https://download.gnome.org/sources/gtksourceview/5.12/gtksourceview-5.12.1.tar.xz"
	gtksourceview5_cmd_src := exec.Command("curl", "-L", gtksourceview5_src_url, "-o", "source.tar.gz")
	err = gtksourceview5_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtksourceview5_cmd_direct := exec.Command("./binary")
	err = gtksourceview5_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: fribidi")
	exec.Command("latte", "install", "fribidi").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
