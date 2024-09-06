package main

// GupnpTools - Free replacements of Intel's UPnP tools
// Homepage: https://wiki.gnome.org/GUPnP/

import (
	"fmt"
	
	"os/exec"
)

func installGupnpTools() {
	// Método 1: Descargar y extraer .tar.gz
	gupnptools_tar_url := "https://download.gnome.org/sources/gupnp-tools/0.12/gupnp-tools-0.12.1.tar.xz"
	gupnptools_cmd_tar := exec.Command("curl", "-L", gupnptools_tar_url, "-o", "package.tar.gz")
	err := gupnptools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gupnptools_zip_url := "https://download.gnome.org/sources/gupnp-tools/0.12/gupnp-tools-0.12.1.tar.xz"
	gupnptools_cmd_zip := exec.Command("curl", "-L", gupnptools_zip_url, "-o", "package.zip")
	err = gupnptools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gupnptools_bin_url := "https://download.gnome.org/sources/gupnp-tools/0.12/gupnp-tools-0.12.1.tar.xz"
	gupnptools_cmd_bin := exec.Command("curl", "-L", gupnptools_bin_url, "-o", "binary.bin")
	err = gupnptools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gupnptools_src_url := "https://download.gnome.org/sources/gupnp-tools/0.12/gupnp-tools-0.12.1.tar.xz"
	gupnptools_cmd_src := exec.Command("curl", "-L", gupnptools_src_url, "-o", "source.tar.gz")
	err = gupnptools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gupnptools_cmd_direct := exec.Command("./binary")
	err = gupnptools_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gssdp")
exec.Command("latte", "install", "gssdp")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: gtksourceview4")
exec.Command("latte", "install", "gtksourceview4")
	fmt.Println("Instalando dependencia: gupnp")
exec.Command("latte", "install", "gupnp")
	fmt.Println("Instalando dependencia: gupnp-av")
exec.Command("latte", "install", "gupnp-av")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
