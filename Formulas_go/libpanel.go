package main

// Libpanel - Dock/panel library for GTK 4
// Homepage: https://gitlab.gnome.org/GNOME/libpanel

import (
	"fmt"
	
	"os/exec"
)

func installLibpanel() {
	// Método 1: Descargar y extraer .tar.gz
	libpanel_tar_url := "https://download.gnome.org/sources/libpanel/1.6/libpanel-1.6.0.tar.xz"
	libpanel_cmd_tar := exec.Command("curl", "-L", libpanel_tar_url, "-o", "package.tar.gz")
	err := libpanel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpanel_zip_url := "https://download.gnome.org/sources/libpanel/1.6/libpanel-1.6.0.tar.xz"
	libpanel_cmd_zip := exec.Command("curl", "-L", libpanel_zip_url, "-o", "package.zip")
	err = libpanel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpanel_bin_url := "https://download.gnome.org/sources/libpanel/1.6/libpanel-1.6.0.tar.xz"
	libpanel_cmd_bin := exec.Command("curl", "-L", libpanel_bin_url, "-o", "binary.bin")
	err = libpanel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpanel_src_url := "https://download.gnome.org/sources/libpanel/1.6/libpanel-1.6.0.tar.xz"
	libpanel_cmd_src := exec.Command("curl", "-L", libpanel_src_url, "-o", "source.tar.gz")
	err = libpanel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpanel_cmd_direct := exec.Command("./binary")
	err = libpanel_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gi-docgen")
exec.Command("latte", "install", "gi-docgen")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: graphene")
exec.Command("latte", "install", "graphene")
	fmt.Println("Instalando dependencia: gtk4")
exec.Command("latte", "install", "gtk4")
	fmt.Println("Instalando dependencia: libadwaita")
exec.Command("latte", "install", "libadwaita")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
