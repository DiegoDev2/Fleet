package main

// LibgeditTepl - Gedit Technology - Text editor product line
// Homepage: https://gitlab.gnome.org/World/gedit/libgedit-tepl

import (
	"fmt"
	
	"os/exec"
)

func installLibgeditTepl() {
	// Método 1: Descargar y extraer .tar.gz
	libgedittepl_tar_url := "https://gitlab.gnome.org/World/gedit/libgedit-tepl/-/archive/6.10.0/libgedit-tepl-6.10.0.tar.bz2"
	libgedittepl_cmd_tar := exec.Command("curl", "-L", libgedittepl_tar_url, "-o", "package.tar.gz")
	err := libgedittepl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgedittepl_zip_url := "https://gitlab.gnome.org/World/gedit/libgedit-tepl/-/archive/6.10.0/libgedit-tepl-6.10.0.tar.bz2"
	libgedittepl_cmd_zip := exec.Command("curl", "-L", libgedittepl_zip_url, "-o", "package.zip")
	err = libgedittepl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgedittepl_bin_url := "https://gitlab.gnome.org/World/gedit/libgedit-tepl/-/archive/6.10.0/libgedit-tepl-6.10.0.tar.bz2"
	libgedittepl_cmd_bin := exec.Command("curl", "-L", libgedittepl_bin_url, "-o", "binary.bin")
	err = libgedittepl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgedittepl_src_url := "https://gitlab.gnome.org/World/gedit/libgedit-tepl/-/archive/6.10.0/libgedit-tepl-6.10.0.tar.bz2"
	libgedittepl_cmd_src := exec.Command("curl", "-L", libgedittepl_src_url, "-o", "source.tar.gz")
	err = libgedittepl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgedittepl_cmd_direct := exec.Command("./binary")
	err = libgedittepl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libgedit-amtk")
	exec.Command("latte", "install", "libgedit-amtk").Run()
	fmt.Println("Instalando dependencia: libgedit-gfls")
	exec.Command("latte", "install", "libgedit-gfls").Run()
	fmt.Println("Instalando dependencia: libgedit-gtksourceview")
	exec.Command("latte", "install", "libgedit-gtksourceview").Run()
	fmt.Println("Instalando dependencia: libhandy")
	exec.Command("latte", "install", "libhandy").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
