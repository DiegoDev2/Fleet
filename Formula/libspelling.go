package main

// Libspelling - Spellcheck library for GTK 4
// Homepage: https://gitlab.gnome.org/GNOME/libspelling

import (
	"fmt"
	
	"os/exec"
)

func installLibspelling() {
	// Método 1: Descargar y extraer .tar.gz
	libspelling_tar_url := "https://gitlab.gnome.org/GNOME/libspelling/-/archive/0.2.1/libspelling-0.2.1.tar.bz2"
	libspelling_cmd_tar := exec.Command("curl", "-L", libspelling_tar_url, "-o", "package.tar.gz")
	err := libspelling_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspelling_zip_url := "https://gitlab.gnome.org/GNOME/libspelling/-/archive/0.2.1/libspelling-0.2.1.tar.bz2"
	libspelling_cmd_zip := exec.Command("curl", "-L", libspelling_zip_url, "-o", "package.zip")
	err = libspelling_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspelling_bin_url := "https://gitlab.gnome.org/GNOME/libspelling/-/archive/0.2.1/libspelling-0.2.1.tar.bz2"
	libspelling_cmd_bin := exec.Command("curl", "-L", libspelling_bin_url, "-o", "binary.bin")
	err = libspelling_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspelling_src_url := "https://gitlab.gnome.org/GNOME/libspelling/-/archive/0.2.1/libspelling-0.2.1.tar.bz2"
	libspelling_cmd_src := exec.Command("curl", "-L", libspelling_src_url, "-o", "source.tar.gz")
	err = libspelling_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspelling_cmd_direct := exec.Command("./binary")
	err = libspelling_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: enchant")
	exec.Command("latte", "install", "enchant").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: gtksourceview5")
	exec.Command("latte", "install", "gtksourceview5").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: graphene")
	exec.Command("latte", "install", "graphene").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
}
