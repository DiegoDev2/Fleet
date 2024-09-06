package main

// Libadwaita - Building blocks for modern adaptive GNOME applications
// Homepage: https://gnome.pages.gitlab.gnome.org/libadwaita/

import (
	"fmt"
	
	"os/exec"
)

func installLibadwaita() {
	// Método 1: Descargar y extraer .tar.gz
	libadwaita_tar_url := "https://download.gnome.org/sources/libadwaita/1.5/libadwaita-1.5.3.tar.xz"
	libadwaita_cmd_tar := exec.Command("curl", "-L", libadwaita_tar_url, "-o", "package.tar.gz")
	err := libadwaita_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libadwaita_zip_url := "https://download.gnome.org/sources/libadwaita/1.5/libadwaita-1.5.3.tar.xz"
	libadwaita_cmd_zip := exec.Command("curl", "-L", libadwaita_zip_url, "-o", "package.zip")
	err = libadwaita_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libadwaita_bin_url := "https://download.gnome.org/sources/libadwaita/1.5/libadwaita-1.5.3.tar.xz"
	libadwaita_cmd_bin := exec.Command("curl", "-L", libadwaita_bin_url, "-o", "binary.bin")
	err = libadwaita_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libadwaita_src_url := "https://download.gnome.org/sources/libadwaita/1.5/libadwaita-1.5.3.tar.xz"
	libadwaita_cmd_src := exec.Command("curl", "-L", libadwaita_src_url, "-o", "source.tar.gz")
	err = libadwaita_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libadwaita_cmd_direct := exec.Command("./binary")
	err = libadwaita_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: appstream")
exec.Command("latte", "install", "appstream")
	fmt.Println("Instalando dependencia: fribidi")
exec.Command("latte", "install", "fribidi")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: graphene")
exec.Command("latte", "install", "graphene")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
