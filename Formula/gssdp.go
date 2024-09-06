package main

// Gssdp - GUPnP library for resource discovery and announcement over SSDP
// Homepage: https://wiki.gnome.org/GUPnP/

import (
	"fmt"
	
	"os/exec"
)

func installGssdp() {
	// Método 1: Descargar y extraer .tar.gz
	gssdp_tar_url := "https://download.gnome.org/sources/gssdp/1.6/gssdp-1.6.3.tar.xz"
	gssdp_cmd_tar := exec.Command("curl", "-L", gssdp_tar_url, "-o", "package.tar.gz")
	err := gssdp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gssdp_zip_url := "https://download.gnome.org/sources/gssdp/1.6/gssdp-1.6.3.tar.xz"
	gssdp_cmd_zip := exec.Command("curl", "-L", gssdp_zip_url, "-o", "package.zip")
	err = gssdp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gssdp_bin_url := "https://download.gnome.org/sources/gssdp/1.6/gssdp-1.6.3.tar.xz"
	gssdp_cmd_bin := exec.Command("curl", "-L", gssdp_bin_url, "-o", "binary.bin")
	err = gssdp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gssdp_src_url := "https://download.gnome.org/sources/gssdp/1.6/gssdp-1.6.3.tar.xz"
	gssdp_cmd_src := exec.Command("curl", "-L", gssdp_src_url, "-o", "source.tar.gz")
	err = gssdp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gssdp_cmd_direct := exec.Command("./binary")
	err = gssdp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pandoc")
	exec.Command("latte", "install", "pandoc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: libsoup")
	exec.Command("latte", "install", "libsoup").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
