package main

// Zenity - GTK+ dialog boxes for the command-line
// Homepage: https://wiki.gnome.org/Projects/Zenity

import (
	"fmt"
	
	"os/exec"
)

func installZenity() {
	// Método 1: Descargar y extraer .tar.gz
	zenity_tar_url := "https://download.gnome.org/sources/zenity/4.0/zenity-4.0.2.tar.xz"
	zenity_cmd_tar := exec.Command("curl", "-L", zenity_tar_url, "-o", "package.tar.gz")
	err := zenity_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zenity_zip_url := "https://download.gnome.org/sources/zenity/4.0/zenity-4.0.2.tar.xz"
	zenity_cmd_zip := exec.Command("curl", "-L", zenity_zip_url, "-o", "package.zip")
	err = zenity_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zenity_bin_url := "https://download.gnome.org/sources/zenity/4.0/zenity-4.0.2.tar.xz"
	zenity_cmd_bin := exec.Command("curl", "-L", zenity_bin_url, "-o", "binary.bin")
	err = zenity_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zenity_src_url := "https://download.gnome.org/sources/zenity/4.0/zenity-4.0.2.tar.xz"
	zenity_cmd_src := exec.Command("curl", "-L", zenity_src_url, "-o", "source.tar.gz")
	err = zenity_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zenity_cmd_direct := exec.Command("./binary")
	err = zenity_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: itstool")
	exec.Command("latte", "install", "itstool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: libadwaita")
	exec.Command("latte", "install", "libadwaita").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
