package main

// LibvirtGlib - Libvirt API for glib-based programs
// Homepage: https://libvirt.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibvirtGlib() {
	// Método 1: Descargar y extraer .tar.gz
	libvirtglib_tar_url := "https://download.libvirt.org/glib/libvirt-glib-5.0.0.tar.xz"
	libvirtglib_cmd_tar := exec.Command("curl", "-L", libvirtglib_tar_url, "-o", "package.tar.gz")
	err := libvirtglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvirtglib_zip_url := "https://download.libvirt.org/glib/libvirt-glib-5.0.0.tar.xz"
	libvirtglib_cmd_zip := exec.Command("curl", "-L", libvirtglib_zip_url, "-o", "package.zip")
	err = libvirtglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvirtglib_bin_url := "https://download.libvirt.org/glib/libvirt-glib-5.0.0.tar.xz"
	libvirtglib_cmd_bin := exec.Command("curl", "-L", libvirtglib_bin_url, "-o", "binary.bin")
	err = libvirtglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvirtglib_src_url := "https://download.libvirt.org/glib/libvirt-glib-5.0.0.tar.xz"
	libvirtglib_cmd_src := exec.Command("curl", "-L", libvirtglib_src_url, "-o", "source.tar.gz")
	err = libvirtglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvirtglib_cmd_direct := exec.Command("./binary")
	err = libvirtglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libvirt")
exec.Command("latte", "install", "libvirt")
}
