package main

// Gtkmm - C++ interfaces for GTK+ and GNOME
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installGtkmm() {
	// Método 1: Descargar y extraer .tar.gz
	gtkmm_tar_url := "https://download.gnome.org/sources/gtkmm/2.24/gtkmm-2.24.5.tar.xz"
	gtkmm_cmd_tar := exec.Command("curl", "-L", gtkmm_tar_url, "-o", "package.tar.gz")
	err := gtkmm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkmm_zip_url := "https://download.gnome.org/sources/gtkmm/2.24/gtkmm-2.24.5.tar.xz"
	gtkmm_cmd_zip := exec.Command("curl", "-L", gtkmm_zip_url, "-o", "package.zip")
	err = gtkmm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkmm_bin_url := "https://download.gnome.org/sources/gtkmm/2.24/gtkmm-2.24.5.tar.xz"
	gtkmm_cmd_bin := exec.Command("curl", "-L", gtkmm_bin_url, "-o", "binary.bin")
	err = gtkmm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkmm_src_url := "https://download.gnome.org/sources/gtkmm/2.24/gtkmm-2.24.5.tar.xz"
	gtkmm_cmd_src := exec.Command("curl", "-L", gtkmm_src_url, "-o", "source.tar.gz")
	err = gtkmm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkmm_cmd_direct := exec.Command("./binary")
	err = gtkmm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: atkmm@2.28")
	exec.Command("latte", "install", "atkmm@2.28").Run()
	fmt.Println("Instalando dependencia: cairomm@1.14")
	exec.Command("latte", "install", "cairomm@1.14").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: glibmm@2.66")
	exec.Command("latte", "install", "glibmm@2.66").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
	fmt.Println("Instalando dependencia: pangomm@2.46")
	exec.Command("latte", "install", "pangomm@2.46").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
