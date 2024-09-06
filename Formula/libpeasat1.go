package main

// LibpeasAT1 - GObject plugin library
// Homepage: https://wiki.gnome.org/Projects/Libpeas

import (
	"fmt"
	
	"os/exec"
)

func installLibpeasAT1() {
	// Método 1: Descargar y extraer .tar.gz
	libpeasat1_tar_url := "https://download.gnome.org/sources/libpeas/1.36/libpeas-1.36.0.tar.xz"
	libpeasat1_cmd_tar := exec.Command("curl", "-L", libpeasat1_tar_url, "-o", "package.tar.gz")
	err := libpeasat1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpeasat1_zip_url := "https://download.gnome.org/sources/libpeas/1.36/libpeas-1.36.0.tar.xz"
	libpeasat1_cmd_zip := exec.Command("curl", "-L", libpeasat1_zip_url, "-o", "package.zip")
	err = libpeasat1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpeasat1_bin_url := "https://download.gnome.org/sources/libpeas/1.36/libpeas-1.36.0.tar.xz"
	libpeasat1_cmd_bin := exec.Command("curl", "-L", libpeasat1_bin_url, "-o", "binary.bin")
	err = libpeasat1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpeasat1_src_url := "https://download.gnome.org/sources/libpeas/1.36/libpeas-1.36.0.tar.xz"
	libpeasat1_cmd_src := exec.Command("curl", "-L", libpeasat1_src_url, "-o", "source.tar.gz")
	err = libpeasat1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpeasat1_cmd_direct := exec.Command("./binary")
	err = libpeasat1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: pygobject3")
	exec.Command("latte", "install", "pygobject3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
