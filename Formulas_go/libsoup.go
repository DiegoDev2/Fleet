package main

// Libsoup - HTTP client/server library for GNOME
// Homepage: https://wiki.gnome.org/Projects/libsoup

import (
	"fmt"
	
	"os/exec"
)

func installLibsoup() {
	// Método 1: Descargar y extraer .tar.gz
	libsoup_tar_url := "https://download.gnome.org/sources/libsoup/3.6/libsoup-3.6.0.tar.xz"
	libsoup_cmd_tar := exec.Command("curl", "-L", libsoup_tar_url, "-o", "package.tar.gz")
	err := libsoup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsoup_zip_url := "https://download.gnome.org/sources/libsoup/3.6/libsoup-3.6.0.tar.xz"
	libsoup_cmd_zip := exec.Command("curl", "-L", libsoup_zip_url, "-o", "package.zip")
	err = libsoup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsoup_bin_url := "https://download.gnome.org/sources/libsoup/3.6/libsoup-3.6.0.tar.xz"
	libsoup_cmd_bin := exec.Command("curl", "-L", libsoup_bin_url, "-o", "binary.bin")
	err = libsoup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsoup_src_url := "https://download.gnome.org/sources/libsoup/3.6/libsoup-3.6.0.tar.xz"
	libsoup_cmd_src := exec.Command("curl", "-L", libsoup_src_url, "-o", "source.tar.gz")
	err = libsoup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsoup_cmd_direct := exec.Command("./binary")
	err = libsoup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: glib-networking")
exec.Command("latte", "install", "glib-networking")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: libpsl")
exec.Command("latte", "install", "libpsl")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
}
