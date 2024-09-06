package main

// Librest - Library to access RESTful web services
// Homepage: https://wiki.gnome.org/Projects/Librest

import (
	"fmt"
	
	"os/exec"
)

func installLibrest() {
	// Método 1: Descargar y extraer .tar.gz
	librest_tar_url := "https://download.gnome.org/sources/rest/0.8/rest-0.8.1.tar.xz"
	librest_cmd_tar := exec.Command("curl", "-L", librest_tar_url, "-o", "package.tar.gz")
	err := librest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librest_zip_url := "https://download.gnome.org/sources/rest/0.8/rest-0.8.1.tar.xz"
	librest_cmd_zip := exec.Command("curl", "-L", librest_zip_url, "-o", "package.zip")
	err = librest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librest_bin_url := "https://download.gnome.org/sources/rest/0.8/rest-0.8.1.tar.xz"
	librest_cmd_bin := exec.Command("curl", "-L", librest_bin_url, "-o", "binary.bin")
	err = librest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librest_src_url := "https://download.gnome.org/sources/rest/0.8/rest-0.8.1.tar.xz"
	librest_cmd_src := exec.Command("curl", "-L", librest_src_url, "-o", "source.tar.gz")
	err = librest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librest_cmd_direct := exec.Command("./binary")
	err = librest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libsoup@2")
exec.Command("latte", "install", "libsoup@2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
