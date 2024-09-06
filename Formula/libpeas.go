package main

// Libpeas - GObject plugin library
// Homepage: https://wiki.gnome.org/Projects/Libpeas

import (
	"fmt"
	
	"os/exec"
)

func installLibpeas() {
	// Método 1: Descargar y extraer .tar.gz
	libpeas_tar_url := "https://download.gnome.org/sources/libpeas/2.0/libpeas-2.0.3.tar.xz"
	libpeas_cmd_tar := exec.Command("curl", "-L", libpeas_tar_url, "-o", "package.tar.gz")
	err := libpeas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpeas_zip_url := "https://download.gnome.org/sources/libpeas/2.0/libpeas-2.0.3.tar.xz"
	libpeas_cmd_zip := exec.Command("curl", "-L", libpeas_zip_url, "-o", "package.zip")
	err = libpeas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpeas_bin_url := "https://download.gnome.org/sources/libpeas/2.0/libpeas-2.0.3.tar.xz"
	libpeas_cmd_bin := exec.Command("curl", "-L", libpeas_bin_url, "-o", "binary.bin")
	err = libpeas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpeas_src_url := "https://download.gnome.org/sources/libpeas/2.0/libpeas-2.0.3.tar.xz"
	libpeas_cmd_src := exec.Command("curl", "-L", libpeas_src_url, "-o", "source.tar.gz")
	err = libpeas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpeas_cmd_direct := exec.Command("./binary")
	err = libpeas_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gjs")
	exec.Command("latte", "install", "gjs").Run()
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
	fmt.Println("Instalando dependencia: spidermonkey@115")
	exec.Command("latte", "install", "spidermonkey@115").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
