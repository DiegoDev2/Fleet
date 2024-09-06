package main

// GlibNetworking - Network related modules for glib
// Homepage: https://gitlab.gnome.org/GNOME/glib-networking

import (
	"fmt"
	
	"os/exec"
)

func installGlibNetworking() {
	// Método 1: Descargar y extraer .tar.gz
	glibnetworking_tar_url := "https://download.gnome.org/sources/glib-networking/2.80/glib-networking-2.80.0.tar.xz"
	glibnetworking_cmd_tar := exec.Command("curl", "-L", glibnetworking_tar_url, "-o", "package.tar.gz")
	err := glibnetworking_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glibnetworking_zip_url := "https://download.gnome.org/sources/glib-networking/2.80/glib-networking-2.80.0.tar.xz"
	glibnetworking_cmd_zip := exec.Command("curl", "-L", glibnetworking_zip_url, "-o", "package.zip")
	err = glibnetworking_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glibnetworking_bin_url := "https://download.gnome.org/sources/glib-networking/2.80/glib-networking-2.80.0.tar.xz"
	glibnetworking_cmd_bin := exec.Command("curl", "-L", glibnetworking_bin_url, "-o", "binary.bin")
	err = glibnetworking_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glibnetworking_src_url := "https://download.gnome.org/sources/glib-networking/2.80/glib-networking-2.80.0.tar.xz"
	glibnetworking_cmd_src := exec.Command("curl", "-L", glibnetworking_src_url, "-o", "source.tar.gz")
	err = glibnetworking_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glibnetworking_cmd_direct := exec.Command("./binary")
	err = glibnetworking_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: gsettings-desktop-schemas")
	exec.Command("latte", "install", "gsettings-desktop-schemas").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
