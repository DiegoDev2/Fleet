package main

// Gtkspell3 - Gtk widget for highlighting and replacing misspelled words
// Homepage: https://gtkspell.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGtkspell3() {
	// Método 1: Descargar y extraer .tar.gz
	gtkspell3_tar_url := "https://downloads.sourceforge.net/project/gtkspell/3.0.10/gtkspell3-3.0.10.tar.xz"
	gtkspell3_cmd_tar := exec.Command("curl", "-L", gtkspell3_tar_url, "-o", "package.tar.gz")
	err := gtkspell3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkspell3_zip_url := "https://downloads.sourceforge.net/project/gtkspell/3.0.10/gtkspell3-3.0.10.tar.xz"
	gtkspell3_cmd_zip := exec.Command("curl", "-L", gtkspell3_zip_url, "-o", "package.zip")
	err = gtkspell3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkspell3_bin_url := "https://downloads.sourceforge.net/project/gtkspell/3.0.10/gtkspell3-3.0.10.tar.xz"
	gtkspell3_cmd_bin := exec.Command("curl", "-L", gtkspell3_bin_url, "-o", "binary.bin")
	err = gtkspell3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkspell3_src_url := "https://downloads.sourceforge.net/project/gtkspell/3.0.10/gtkspell3-3.0.10.tar.xz"
	gtkspell3_cmd_src := exec.Command("curl", "-L", gtkspell3_src_url, "-o", "source.tar.gz")
	err = gtkspell3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkspell3_cmd_direct := exec.Command("./binary")
	err = gtkspell3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: enchant")
exec.Command("latte", "install", "enchant")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
