package main

// Libidl - Library for creating CORBA IDL files
// Homepage: https://ftp.acc.umu.se/pub/gnome/sources/libIDL/0.8/

import (
	"fmt"
	
	"os/exec"
)

func installLibidl() {
	// Método 1: Descargar y extraer .tar.gz
	libidl_tar_url := "https://download.gnome.org/sources/libIDL/0.8/libIDL-0.8.14.tar.bz2"
	libidl_cmd_tar := exec.Command("curl", "-L", libidl_tar_url, "-o", "package.tar.gz")
	err := libidl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libidl_zip_url := "https://download.gnome.org/sources/libIDL/0.8/libIDL-0.8.14.tar.bz2"
	libidl_cmd_zip := exec.Command("curl", "-L", libidl_zip_url, "-o", "package.zip")
	err = libidl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libidl_bin_url := "https://download.gnome.org/sources/libIDL/0.8/libIDL-0.8.14.tar.bz2"
	libidl_cmd_bin := exec.Command("curl", "-L", libidl_bin_url, "-o", "binary.bin")
	err = libidl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libidl_src_url := "https://download.gnome.org/sources/libIDL/0.8/libIDL-0.8.14.tar.bz2"
	libidl_cmd_src := exec.Command("curl", "-L", libidl_src_url, "-o", "source.tar.gz")
	err = libidl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libidl_cmd_direct := exec.Command("./binary")
	err = libidl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
