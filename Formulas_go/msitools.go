package main

// Msitools - Windows installer (.MSI) tool
// Homepage: https://wiki.gnome.org/msitools

import (
	"fmt"
	
	"os/exec"
)

func installMsitools() {
	// Método 1: Descargar y extraer .tar.gz
	msitools_tar_url := "https://download.gnome.org/sources/msitools/0.103/msitools-0.103.tar.xz"
	msitools_cmd_tar := exec.Command("curl", "-L", msitools_tar_url, "-o", "package.tar.gz")
	err := msitools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msitools_zip_url := "https://download.gnome.org/sources/msitools/0.103/msitools-0.103.tar.xz"
	msitools_cmd_zip := exec.Command("curl", "-L", msitools_zip_url, "-o", "package.zip")
	err = msitools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msitools_bin_url := "https://download.gnome.org/sources/msitools/0.103/msitools-0.103.tar.xz"
	msitools_cmd_bin := exec.Command("curl", "-L", msitools_bin_url, "-o", "binary.bin")
	err = msitools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msitools_src_url := "https://download.gnome.org/sources/msitools/0.103/msitools-0.103.tar.xz"
	msitools_cmd_src := exec.Command("curl", "-L", msitools_src_url, "-o", "source.tar.gz")
	err = msitools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msitools_cmd_direct := exec.Command("./binary")
	err = msitools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: gcab")
exec.Command("latte", "install", "gcab")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libgsf")
exec.Command("latte", "install", "libgsf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
