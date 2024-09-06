package main

// AtkmmAT228 - Official C++ interface for the ATK accessibility toolkit library
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installAtkmmAT228() {
	// Método 1: Descargar y extraer .tar.gz
	atkmmat228_tar_url := "https://download.gnome.org/sources/atkmm/2.28/atkmm-2.28.4.tar.xz"
	atkmmat228_cmd_tar := exec.Command("curl", "-L", atkmmat228_tar_url, "-o", "package.tar.gz")
	err := atkmmat228_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atkmmat228_zip_url := "https://download.gnome.org/sources/atkmm/2.28/atkmm-2.28.4.tar.xz"
	atkmmat228_cmd_zip := exec.Command("curl", "-L", atkmmat228_zip_url, "-o", "package.zip")
	err = atkmmat228_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atkmmat228_bin_url := "https://download.gnome.org/sources/atkmm/2.28/atkmm-2.28.4.tar.xz"
	atkmmat228_cmd_bin := exec.Command("curl", "-L", atkmmat228_bin_url, "-o", "binary.bin")
	err = atkmmat228_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atkmmat228_src_url := "https://download.gnome.org/sources/atkmm/2.28/atkmm-2.28.4.tar.xz"
	atkmmat228_cmd_src := exec.Command("curl", "-L", atkmmat228_src_url, "-o", "source.tar.gz")
	err = atkmmat228_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atkmmat228_cmd_direct := exec.Command("./binary")
	err = atkmmat228_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: glibmm@2.66")
	exec.Command("latte", "install", "glibmm@2.66").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
