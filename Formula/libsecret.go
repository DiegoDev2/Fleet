package main

// Libsecret - Library for storing/retrieving passwords and other secrets
// Homepage: https://wiki.gnome.org/Projects/Libsecret

import (
	"fmt"
	
	"os/exec"
)

func installLibsecret() {
	// Método 1: Descargar y extraer .tar.gz
	libsecret_tar_url := "https://download.gnome.org/sources/libsecret/0.21/libsecret-0.21.4.tar.xz"
	libsecret_cmd_tar := exec.Command("curl", "-L", libsecret_tar_url, "-o", "package.tar.gz")
	err := libsecret_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsecret_zip_url := "https://download.gnome.org/sources/libsecret/0.21/libsecret-0.21.4.tar.xz"
	libsecret_cmd_zip := exec.Command("curl", "-L", libsecret_zip_url, "-o", "package.zip")
	err = libsecret_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsecret_bin_url := "https://download.gnome.org/sources/libsecret/0.21/libsecret-0.21.4.tar.xz"
	libsecret_cmd_bin := exec.Command("curl", "-L", libsecret_bin_url, "-o", "binary.bin")
	err = libsecret_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsecret_src_url := "https://download.gnome.org/sources/libsecret/0.21/libsecret-0.21.4.tar.xz"
	libsecret_cmd_src := exec.Command("curl", "-L", libsecret_src_url, "-o", "source.tar.gz")
	err = libsecret_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsecret_cmd_direct := exec.Command("./binary")
	err = libsecret_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
	exec.Command("latte", "install", "docbook-xsl").Run()
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
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
}
