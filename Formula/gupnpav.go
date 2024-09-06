package main

// GupnpAv - Library to help implement UPnP A/V profiles
// Homepage: https://wiki.gnome.org/GUPnP/

import (
	"fmt"
	
	"os/exec"
)

func installGupnpAv() {
	// Método 1: Descargar y extraer .tar.gz
	gupnpav_tar_url := "https://download.gnome.org/sources/gupnp-av/0.14/gupnp-av-0.14.1.tar.xz"
	gupnpav_cmd_tar := exec.Command("curl", "-L", gupnpav_tar_url, "-o", "package.tar.gz")
	err := gupnpav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gupnpav_zip_url := "https://download.gnome.org/sources/gupnp-av/0.14/gupnp-av-0.14.1.tar.xz"
	gupnpav_cmd_zip := exec.Command("curl", "-L", gupnpav_zip_url, "-o", "package.zip")
	err = gupnpav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gupnpav_bin_url := "https://download.gnome.org/sources/gupnp-av/0.14/gupnp-av-0.14.1.tar.xz"
	gupnpav_cmd_bin := exec.Command("curl", "-L", gupnpav_bin_url, "-o", "binary.bin")
	err = gupnpav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gupnpav_src_url := "https://download.gnome.org/sources/gupnp-av/0.14/gupnp-av-0.14.1.tar.xz"
	gupnpav_cmd_src := exec.Command("curl", "-L", gupnpav_src_url, "-o", "source.tar.gz")
	err = gupnpav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gupnpav_cmd_direct := exec.Command("./binary")
	err = gupnpav_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
