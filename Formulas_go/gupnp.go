package main

// Gupnp - Framework for creating UPnP devices and control points
// Homepage: https://wiki.gnome.org/Projects/GUPnP

import (
	"fmt"
	
	"os/exec"
)

func installGupnp() {
	// Método 1: Descargar y extraer .tar.gz
	gupnp_tar_url := "https://download.gnome.org/sources/gupnp/1.6/gupnp-1.6.6.tar.xz"
	gupnp_cmd_tar := exec.Command("curl", "-L", gupnp_tar_url, "-o", "package.tar.gz")
	err := gupnp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gupnp_zip_url := "https://download.gnome.org/sources/gupnp/1.6/gupnp-1.6.6.tar.xz"
	gupnp_cmd_zip := exec.Command("curl", "-L", gupnp_zip_url, "-o", "package.zip")
	err = gupnp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gupnp_bin_url := "https://download.gnome.org/sources/gupnp/1.6/gupnp-1.6.6.tar.xz"
	gupnp_cmd_bin := exec.Command("curl", "-L", gupnp_bin_url, "-o", "binary.bin")
	err = gupnp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gupnp_src_url := "https://download.gnome.org/sources/gupnp/1.6/gupnp-1.6.6.tar.xz"
	gupnp_cmd_src := exec.Command("curl", "-L", gupnp_src_url, "-o", "source.tar.gz")
	err = gupnp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gupnp_cmd_direct := exec.Command("./binary")
	err = gupnp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
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
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gssdp")
exec.Command("latte", "install", "gssdp")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
