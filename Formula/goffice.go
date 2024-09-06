package main

// Goffice - Gnumeric spreadsheet program
// Homepage: https://gitlab.gnome.org/GNOME/goffice

import (
	"fmt"
	
	"os/exec"
)

func installGoffice() {
	// Método 1: Descargar y extraer .tar.gz
	goffice_tar_url := "https://download.gnome.org/sources/goffice/0.10/goffice-0.10.57.tar.xz"
	goffice_cmd_tar := exec.Command("curl", "-L", goffice_tar_url, "-o", "package.tar.gz")
	err := goffice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	goffice_zip_url := "https://download.gnome.org/sources/goffice/0.10/goffice-0.10.57.tar.xz"
	goffice_cmd_zip := exec.Command("curl", "-L", goffice_zip_url, "-o", "package.zip")
	err = goffice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	goffice_bin_url := "https://download.gnome.org/sources/goffice/0.10/goffice-0.10.57.tar.xz"
	goffice_cmd_bin := exec.Command("curl", "-L", goffice_bin_url, "-o", "binary.bin")
	err = goffice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	goffice_src_url := "https://download.gnome.org/sources/goffice/0.10/goffice-0.10.57.tar.xz"
	goffice_cmd_src := exec.Command("curl", "-L", goffice_src_url, "-o", "source.tar.gz")
	err = goffice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	goffice_cmd_direct := exec.Command("./binary")
	err = goffice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gtk-doc")
	exec.Command("latte", "install", "gtk-doc").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: libgsf")
	exec.Command("latte", "install", "libgsf").Run()
	fmt.Println("Instalando dependencia: librsvg")
	exec.Command("latte", "install", "librsvg").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: perl-xml-parser")
	exec.Command("latte", "install", "perl-xml-parser").Run()
}
