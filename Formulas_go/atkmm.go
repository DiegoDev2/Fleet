package main

// Atkmm - Official C++ interface for the ATK accessibility toolkit library
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installAtkmm() {
	// Método 1: Descargar y extraer .tar.gz
	atkmm_tar_url := "https://download.gnome.org/sources/atkmm/2.36/atkmm-2.36.3.tar.xz"
	atkmm_cmd_tar := exec.Command("curl", "-L", atkmm_tar_url, "-o", "package.tar.gz")
	err := atkmm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atkmm_zip_url := "https://download.gnome.org/sources/atkmm/2.36/atkmm-2.36.3.tar.xz"
	atkmm_cmd_zip := exec.Command("curl", "-L", atkmm_zip_url, "-o", "package.zip")
	err = atkmm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atkmm_bin_url := "https://download.gnome.org/sources/atkmm/2.36/atkmm-2.36.3.tar.xz"
	atkmm_cmd_bin := exec.Command("curl", "-L", atkmm_bin_url, "-o", "binary.bin")
	err = atkmm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atkmm_src_url := "https://download.gnome.org/sources/atkmm/2.36/atkmm-2.36.3.tar.xz"
	atkmm_cmd_src := exec.Command("curl", "-L", atkmm_src_url, "-o", "source.tar.gz")
	err = atkmm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atkmm_cmd_direct := exec.Command("./binary")
	err = atkmm_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: glibmm")
exec.Command("latte", "install", "glibmm")
	fmt.Println("Instalando dependencia: libsigc++")
exec.Command("latte", "install", "libsigc++")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
