package main

// Gtkmm4 - C++ interfaces for GTK+ and GNOME
// Homepage: https://www.gtkmm.org/

import (
	"fmt"
	
	"os/exec"
)

func installGtkmm4() {
	// Método 1: Descargar y extraer .tar.gz
	gtkmm4_tar_url := "https://download.gnome.org/sources/gtkmm/4.14/gtkmm-4.14.0.tar.xz"
	gtkmm4_cmd_tar := exec.Command("curl", "-L", gtkmm4_tar_url, "-o", "package.tar.gz")
	err := gtkmm4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gtkmm4_zip_url := "https://download.gnome.org/sources/gtkmm/4.14/gtkmm-4.14.0.tar.xz"
	gtkmm4_cmd_zip := exec.Command("curl", "-L", gtkmm4_zip_url, "-o", "package.zip")
	err = gtkmm4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gtkmm4_bin_url := "https://download.gnome.org/sources/gtkmm/4.14/gtkmm-4.14.0.tar.xz"
	gtkmm4_cmd_bin := exec.Command("curl", "-L", gtkmm4_bin_url, "-o", "binary.bin")
	err = gtkmm4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gtkmm4_src_url := "https://download.gnome.org/sources/gtkmm/4.14/gtkmm-4.14.0.tar.xz"
	gtkmm4_cmd_src := exec.Command("curl", "-L", gtkmm4_src_url, "-o", "source.tar.gz")
	err = gtkmm4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gtkmm4_cmd_direct := exec.Command("./binary")
	err = gtkmm4_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: cairomm")
	exec.Command("latte", "install", "cairomm").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: glibmm")
	exec.Command("latte", "install", "glibmm").Run()
	fmt.Println("Instalando dependencia: graphene")
	exec.Command("latte", "install", "graphene").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: libsigc++")
	exec.Command("latte", "install", "libsigc++").Run()
	fmt.Println("Instalando dependencia: pangomm")
	exec.Command("latte", "install", "pangomm").Run()
}
