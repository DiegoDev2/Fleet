package main

// GtkDoc - GTK+ documentation tool
// Homepage: https://gitlab.gnome.org/GNOME/gtk-doc

import (
	"fmt"
	
	"os/exec"
)

func installGtkDoc() {
	// Método 1: Descargar y extraer .tar.gz
	gtkdoc_tar_url := "https://download.gnome.org/sources/gtk-doc/1.34/gtk-doc-1.34.0.tar.xz"
	gtkdoc_cmd_tar := exec.Command("curl", "-L", gtkdoc_tar_url, "-o", "package.tar.gz")
	err := gtkdoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkdoc_zip_url := "https://download.gnome.org/sources/gtk-doc/1.34/gtk-doc-1.34.0.tar.xz"
	gtkdoc_cmd_zip := exec.Command("curl", "-L", gtkdoc_zip_url, "-o", "package.zip")
	err = gtkdoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkdoc_bin_url := "https://download.gnome.org/sources/gtk-doc/1.34/gtk-doc-1.34.0.tar.xz"
	gtkdoc_cmd_bin := exec.Command("curl", "-L", gtkdoc_bin_url, "-o", "binary.bin")
	err = gtkdoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkdoc_src_url := "https://download.gnome.org/sources/gtk-doc/1.34/gtk-doc-1.34.0.tar.xz"
	gtkdoc_cmd_src := exec.Command("curl", "-L", gtkdoc_src_url, "-o", "source.tar.gz")
	err = gtkdoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkdoc_cmd_direct := exec.Command("./binary")
	err = gtkdoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: docbook")
	exec.Command("latte", "install", "docbook").Run()
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
