package main

// GsettingsDesktopSchemas - GSettings schemas for desktop components
// Homepage: https://download.gnome.org/sources/gsettings-desktop-schemas/

import (
	"fmt"
	
	"os/exec"
)

func installGsettingsDesktopSchemas() {
	// Método 1: Descargar y extraer .tar.gz
	gsettingsdesktopschemas_tar_url := "https://download.gnome.org/sources/gsettings-desktop-schemas/46/gsettings-desktop-schemas-46.1.tar.xz"
	gsettingsdesktopschemas_cmd_tar := exec.Command("curl", "-L", gsettingsdesktopschemas_tar_url, "-o", "package.tar.gz")
	err := gsettingsdesktopschemas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gsettingsdesktopschemas_zip_url := "https://download.gnome.org/sources/gsettings-desktop-schemas/46/gsettings-desktop-schemas-46.1.tar.xz"
	gsettingsdesktopschemas_cmd_zip := exec.Command("curl", "-L", gsettingsdesktopschemas_zip_url, "-o", "package.zip")
	err = gsettingsdesktopschemas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gsettingsdesktopschemas_bin_url := "https://download.gnome.org/sources/gsettings-desktop-schemas/46/gsettings-desktop-schemas-46.1.tar.xz"
	gsettingsdesktopschemas_cmd_bin := exec.Command("curl", "-L", gsettingsdesktopschemas_bin_url, "-o", "binary.bin")
	err = gsettingsdesktopschemas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gsettingsdesktopschemas_src_url := "https://download.gnome.org/sources/gsettings-desktop-schemas/46/gsettings-desktop-schemas-46.1.tar.xz"
	gsettingsdesktopschemas_cmd_src := exec.Command("curl", "-L", gsettingsdesktopschemas_src_url, "-o", "source.tar.gz")
	err = gsettingsdesktopschemas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gsettingsdesktopschemas_cmd_direct := exec.Command("./binary")
	err = gsettingsdesktopschemas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
}
