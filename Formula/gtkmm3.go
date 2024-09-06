package main

// Gtkmm3 - C++ interfaces for GTK+ and GNOME
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installGtkmm3() {
	// Método 1: Descargar y extraer .tar.gz
	gtkmm3_tar_url := "https://download.gnome.org/sources/gtkmm/3.24/gtkmm-3.24.9.tar.xz"
	gtkmm3_cmd_tar := exec.Command("curl", "-L", gtkmm3_tar_url, "-o", "package.tar.gz")
	err := gtkmm3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkmm3_zip_url := "https://download.gnome.org/sources/gtkmm/3.24/gtkmm-3.24.9.tar.xz"
	gtkmm3_cmd_zip := exec.Command("curl", "-L", gtkmm3_zip_url, "-o", "package.zip")
	err = gtkmm3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkmm3_bin_url := "https://download.gnome.org/sources/gtkmm/3.24/gtkmm-3.24.9.tar.xz"
	gtkmm3_cmd_bin := exec.Command("curl", "-L", gtkmm3_bin_url, "-o", "binary.bin")
	err = gtkmm3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkmm3_src_url := "https://download.gnome.org/sources/gtkmm/3.24/gtkmm-3.24.9.tar.xz"
	gtkmm3_cmd_src := exec.Command("curl", "-L", gtkmm3_src_url, "-o", "source.tar.gz")
	err = gtkmm3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkmm3_cmd_direct := exec.Command("./binary")
	err = gtkmm3_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: libsigc++@2")
	exec.Command("latte", "install", "libsigc++@2").Run()
	fmt.Println("Instalando dependencia: pangomm@2.46")
	exec.Command("latte", "install", "pangomm@2.46").Run()
}
