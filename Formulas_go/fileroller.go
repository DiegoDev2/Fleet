package main

// FileRoller - GNOME archive manager
// Homepage: https://wiki.gnome.org/Apps/FileRoller

import (
	"fmt"
	
	"os/exec"
)

func installFileRoller() {
	// Método 1: Descargar y extraer .tar.gz
	fileroller_tar_url := "https://download.gnome.org/sources/file-roller/44/file-roller-44.3.tar.xz"
	fileroller_cmd_tar := exec.Command("curl", "-L", fileroller_tar_url, "-o", "package.tar.gz")
	err := fileroller_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fileroller_zip_url := "https://download.gnome.org/sources/file-roller/44/file-roller-44.3.tar.xz"
	fileroller_cmd_zip := exec.Command("curl", "-L", fileroller_zip_url, "-o", "package.zip")
	err = fileroller_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fileroller_bin_url := "https://download.gnome.org/sources/file-roller/44/file-roller-44.3.tar.xz"
	fileroller_cmd_bin := exec.Command("curl", "-L", fileroller_bin_url, "-o", "binary.bin")
	err = fileroller_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fileroller_src_url := "https://download.gnome.org/sources/file-roller/44/file-roller-44.3.tar.xz"
	fileroller_cmd_src := exec.Command("curl", "-L", fileroller_src_url, "-o", "source.tar.gz")
	err = fileroller_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fileroller_cmd_direct := exec.Command("./binary")
	err = fileroller_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
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
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
exec.Command("latte", "install", "hicolor-icon-theme")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: libadwaita")
exec.Command("latte", "install", "libadwaita")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
