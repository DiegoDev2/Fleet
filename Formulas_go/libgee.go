package main

// Libgee - Collection library providing GObject-based interfaces
// Homepage: https://wiki.gnome.org/Projects/Libgee

import (
	"fmt"
	
	"os/exec"
)

func installLibgee() {
	// Método 1: Descargar y extraer .tar.gz
	libgee_tar_url := "https://download.gnome.org/sources/libgee/0.20/libgee-0.20.6.tar.xz"
	libgee_cmd_tar := exec.Command("curl", "-L", libgee_tar_url, "-o", "package.tar.gz")
	err := libgee_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgee_zip_url := "https://download.gnome.org/sources/libgee/0.20/libgee-0.20.6.tar.xz"
	libgee_cmd_zip := exec.Command("curl", "-L", libgee_zip_url, "-o", "package.zip")
	err = libgee_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgee_bin_url := "https://download.gnome.org/sources/libgee/0.20/libgee-0.20.6.tar.xz"
	libgee_cmd_bin := exec.Command("curl", "-L", libgee_bin_url, "-o", "binary.bin")
	err = libgee_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgee_src_url := "https://download.gnome.org/sources/libgee/0.20/libgee-0.20.6.tar.xz"
	libgee_cmd_src := exec.Command("curl", "-L", libgee_src_url, "-o", "source.tar.gz")
	err = libgee_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgee_cmd_direct := exec.Command("./binary")
	err = libgee_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
