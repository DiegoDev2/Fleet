package main

// LibgeditAmtk - Actions, Menus and Toolbars Kit for GTK applications
// Homepage: https://gedit-technology.net

import (
	"fmt"
	
	"os/exec"
)

func installLibgeditAmtk() {
	// Método 1: Descargar y extraer .tar.gz
	libgeditamtk_tar_url := "https://gitlab.gnome.org/World/gedit/libgedit-amtk/-/archive/5.9.0/libgedit-amtk-5.9.0.tar.bz2"
	libgeditamtk_cmd_tar := exec.Command("curl", "-L", libgeditamtk_tar_url, "-o", "package.tar.gz")
	err := libgeditamtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgeditamtk_zip_url := "https://gitlab.gnome.org/World/gedit/libgedit-amtk/-/archive/5.9.0/libgedit-amtk-5.9.0.tar.bz2"
	libgeditamtk_cmd_zip := exec.Command("curl", "-L", libgeditamtk_zip_url, "-o", "package.zip")
	err = libgeditamtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgeditamtk_bin_url := "https://gitlab.gnome.org/World/gedit/libgedit-amtk/-/archive/5.9.0/libgedit-amtk-5.9.0.tar.bz2"
	libgeditamtk_cmd_bin := exec.Command("curl", "-L", libgeditamtk_bin_url, "-o", "binary.bin")
	err = libgeditamtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgeditamtk_src_url := "https://gitlab.gnome.org/World/gedit/libgedit-amtk/-/archive/5.9.0/libgedit-amtk-5.9.0.tar.bz2"
	libgeditamtk_cmd_src := exec.Command("curl", "-L", libgeditamtk_src_url, "-o", "source.tar.gz")
	err = libgeditamtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgeditamtk_cmd_direct := exec.Command("./binary")
	err = libgeditamtk_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
