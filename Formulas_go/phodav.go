package main

// Phodav - WebDav server implementation using libsoup (RFC 4918)
// Homepage: https://gitlab.gnome.org/GNOME/phodav

import (
	"fmt"
	
	"os/exec"
)

func installPhodav() {
	// Método 1: Descargar y extraer .tar.gz
	phodav_tar_url := "https://download.gnome.org/sources/phodav/3.0/phodav-3.0.tar.xz"
	phodav_cmd_tar := exec.Command("curl", "-L", phodav_tar_url, "-o", "package.tar.gz")
	err := phodav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phodav_zip_url := "https://download.gnome.org/sources/phodav/3.0/phodav-3.0.tar.xz"
	phodav_cmd_zip := exec.Command("curl", "-L", phodav_zip_url, "-o", "package.zip")
	err = phodav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phodav_bin_url := "https://download.gnome.org/sources/phodav/3.0/phodav-3.0.tar.xz"
	phodav_cmd_bin := exec.Command("curl", "-L", phodav_bin_url, "-o", "binary.bin")
	err = phodav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phodav_src_url := "https://download.gnome.org/sources/phodav/3.0/phodav-3.0.tar.xz"
	phodav_cmd_src := exec.Command("curl", "-L", phodav_src_url, "-o", "source.tar.gz")
	err = phodav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phodav_cmd_direct := exec.Command("./binary")
	err = phodav_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
