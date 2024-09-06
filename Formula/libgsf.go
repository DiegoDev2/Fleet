package main

// Libgsf - I/O abstraction library for dealing with structured file formats
// Homepage: https://gitlab.gnome.org/GNOME/libgsf

import (
	"fmt"
	
	"os/exec"
)

func installLibgsf() {
	// Método 1: Descargar y extraer .tar.gz
	libgsf_tar_url := "https://download.gnome.org/sources/libgsf/1.14/libgsf-1.14.52.tar.xz"
	libgsf_cmd_tar := exec.Command("curl", "-L", libgsf_tar_url, "-o", "package.tar.gz")
	err := libgsf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgsf_zip_url := "https://download.gnome.org/sources/libgsf/1.14/libgsf-1.14.52.tar.xz"
	libgsf_cmd_zip := exec.Command("curl", "-L", libgsf_zip_url, "-o", "package.zip")
	err = libgsf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgsf_bin_url := "https://download.gnome.org/sources/libgsf/1.14/libgsf-1.14.52.tar.xz"
	libgsf_cmd_bin := exec.Command("curl", "-L", libgsf_bin_url, "-o", "binary.bin")
	err = libgsf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgsf_src_url := "https://download.gnome.org/sources/libgsf/1.14/libgsf-1.14.52.tar.xz"
	libgsf_cmd_src := exec.Command("curl", "-L", libgsf_src_url, "-o", "source.tar.gz")
	err = libgsf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgsf_cmd_direct := exec.Command("./binary")
	err = libgsf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gtk-doc")
	exec.Command("latte", "install", "gtk-doc").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
