package main

// Libgtop - Library for portably obtaining information about processes
// Homepage: https://gitlab.gnome.org/GNOME/libgtop

import (
	"fmt"
	
	"os/exec"
)

func installLibgtop() {
	// Método 1: Descargar y extraer .tar.gz
	libgtop_tar_url := "https://download.gnome.org/sources/libgtop/2.40/libgtop-2.40.0.tar.xz"
	libgtop_cmd_tar := exec.Command("curl", "-L", libgtop_tar_url, "-o", "package.tar.gz")
	err := libgtop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgtop_zip_url := "https://download.gnome.org/sources/libgtop/2.40/libgtop-2.40.0.tar.xz"
	libgtop_cmd_zip := exec.Command("curl", "-L", libgtop_zip_url, "-o", "package.zip")
	err = libgtop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgtop_bin_url := "https://download.gnome.org/sources/libgtop/2.40/libgtop-2.40.0.tar.xz"
	libgtop_cmd_bin := exec.Command("curl", "-L", libgtop_bin_url, "-o", "binary.bin")
	err = libgtop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgtop_src_url := "https://download.gnome.org/sources/libgtop/2.40/libgtop-2.40.0.tar.xz"
	libgtop_cmd_src := exec.Command("curl", "-L", libgtop_src_url, "-o", "source.tar.gz")
	err = libgtop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgtop_cmd_direct := exec.Command("./binary")
	err = libgtop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libxau")
	exec.Command("latte", "install", "libxau").Run()
}
