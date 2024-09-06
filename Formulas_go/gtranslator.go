package main

// Gtranslator - GNOME gettext PO file editor
// Homepage: https://wiki.gnome.org/Design/Apps/Translator

import (
	"fmt"
	
	"os/exec"
)

func installGtranslator() {
	// Método 1: Descargar y extraer .tar.gz
	gtranslator_tar_url := "https://download.gnome.org/sources/gtranslator/46/gtranslator-46.1.tar.xz"
	gtranslator_cmd_tar := exec.Command("curl", "-L", gtranslator_tar_url, "-o", "package.tar.gz")
	err := gtranslator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtranslator_zip_url := "https://download.gnome.org/sources/gtranslator/46/gtranslator-46.1.tar.xz"
	gtranslator_cmd_zip := exec.Command("curl", "-L", gtranslator_zip_url, "-o", "package.zip")
	err = gtranslator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtranslator_bin_url := "https://download.gnome.org/sources/gtranslator/46/gtranslator-46.1.tar.xz"
	gtranslator_cmd_bin := exec.Command("curl", "-L", gtranslator_bin_url, "-o", "binary.bin")
	err = gtranslator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtranslator_src_url := "https://download.gnome.org/sources/gtranslator/46/gtranslator-46.1.tar.xz"
	gtranslator_cmd_src := exec.Command("curl", "-L", gtranslator_src_url, "-o", "source.tar.gz")
	err = gtranslator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtranslator_cmd_direct := exec.Command("./binary")
	err = gtranslator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: desktop-file-utils")
exec.Command("latte", "install", "desktop-file-utils")
	fmt.Println("Instalando dependencia: itstool")
exec.Command("latte", "install", "itstool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gspell")
exec.Command("latte", "install", "gspell")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: gtksourceview5")
exec.Command("latte", "install", "gtksourceview5")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: libadwaita")
exec.Command("latte", "install", "libadwaita")
	fmt.Println("Instalando dependencia: libgda")
exec.Command("latte", "install", "libgda")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: libspelling")
exec.Command("latte", "install", "libspelling")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
