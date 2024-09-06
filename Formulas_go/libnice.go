package main

// Libnice - GLib ICE implementation
// Homepage: https://wiki.freedesktop.org/nice/

import (
	"fmt"
	
	"os/exec"
)

func installLibnice() {
	// Método 1: Descargar y extraer .tar.gz
	libnice_tar_url := "https://libnice.freedesktop.org/releases/libnice-0.1.22.tar.gz"
	libnice_cmd_tar := exec.Command("curl", "-L", libnice_tar_url, "-o", "package.tar.gz")
	err := libnice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnice_zip_url := "https://libnice.freedesktop.org/releases/libnice-0.1.22.zip"
	libnice_cmd_zip := exec.Command("curl", "-L", libnice_zip_url, "-o", "package.zip")
	err = libnice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnice_bin_url := "https://libnice.freedesktop.org/releases/libnice-0.1.22.bin"
	libnice_cmd_bin := exec.Command("curl", "-L", libnice_bin_url, "-o", "binary.bin")
	err = libnice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnice_src_url := "https://libnice.freedesktop.org/releases/libnice-0.1.22.src.tar.gz"
	libnice_cmd_src := exec.Command("curl", "-L", libnice_src_url, "-o", "source.tar.gz")
	err = libnice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnice_cmd_direct := exec.Command("./binary")
	err = libnice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: gstreamer")
exec.Command("latte", "install", "gstreamer")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
}
