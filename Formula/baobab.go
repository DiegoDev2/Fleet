package main

// Baobab - Gnome disk usage analyzer
// Homepage: https://apps.gnome.org/Baobab/

import (
	"fmt"
	
	"os/exec"
)

func installBaobab() {
	// Método 1: Descargar y extraer .tar.gz
	baobab_tar_url := "https://download.gnome.org/sources/baobab/46/baobab-46.0.tar.xz"
	baobab_cmd_tar := exec.Command("curl", "-L", baobab_tar_url, "-o", "package.tar.gz")
	err := baobab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	baobab_zip_url := "https://download.gnome.org/sources/baobab/46/baobab-46.0.tar.xz"
	baobab_cmd_zip := exec.Command("curl", "-L", baobab_zip_url, "-o", "package.zip")
	err = baobab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	baobab_bin_url := "https://download.gnome.org/sources/baobab/46/baobab-46.0.tar.xz"
	baobab_cmd_bin := exec.Command("curl", "-L", baobab_bin_url, "-o", "binary.bin")
	err = baobab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	baobab_src_url := "https://download.gnome.org/sources/baobab/46/baobab-46.0.tar.xz"
	baobab_cmd_src := exec.Command("curl", "-L", baobab_src_url, "-o", "source.tar.gz")
	err = baobab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	baobab_cmd_direct := exec.Command("./binary")
	err = baobab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: desktop-file-utils")
	exec.Command("latte", "install", "desktop-file-utils").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: itstool")
	exec.Command("latte", "install", "itstool").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk4")
	exec.Command("latte", "install", "gtk4").Run()
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
	exec.Command("latte", "install", "hicolor-icon-theme").Run()
	fmt.Println("Instalando dependencia: libadwaita")
	exec.Command("latte", "install", "libadwaita").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
