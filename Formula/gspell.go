package main

// Gspell - Flexible API to implement spellchecking in GTK+ applications
// Homepage: https://gitlab.gnome.org/GNOME/gspell

import (
	"fmt"
	
	"os/exec"
)

func installGspell() {
	// Método 1: Descargar y extraer .tar.gz
	gspell_tar_url := "https://download.gnome.org/sources/gspell/1.12/gspell-1.12.2.tar.xz"
	gspell_cmd_tar := exec.Command("curl", "-L", gspell_tar_url, "-o", "package.tar.gz")
	err := gspell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gspell_zip_url := "https://download.gnome.org/sources/gspell/1.12/gspell-1.12.2.tar.xz"
	gspell_cmd_zip := exec.Command("curl", "-L", gspell_zip_url, "-o", "package.zip")
	err = gspell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gspell_bin_url := "https://download.gnome.org/sources/gspell/1.12/gspell-1.12.2.tar.xz"
	gspell_cmd_bin := exec.Command("curl", "-L", gspell_bin_url, "-o", "binary.bin")
	err = gspell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gspell_src_url := "https://download.gnome.org/sources/gspell/1.12/gspell-1.12.2.tar.xz"
	gspell_cmd_src := exec.Command("curl", "-L", gspell_src_url, "-o", "source.tar.gz")
	err = gspell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gspell_cmd_direct := exec.Command("./binary")
	err = gspell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: enchant")
	exec.Command("latte", "install", "enchant").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
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
	fmt.Println("Instalando dependencia: gtk-mac-integration")
	exec.Command("latte", "install", "gtk-mac-integration").Run()
}
