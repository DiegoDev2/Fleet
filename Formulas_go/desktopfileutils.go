package main

// DesktopFileUtils - Command-line utilities for working with desktop entries
// Homepage: https://wiki.freedesktop.org/www/Software/desktop-file-utils/

import (
	"fmt"
	
	"os/exec"
)

func installDesktopFileUtils() {
	// Método 1: Descargar y extraer .tar.gz
	desktopfileutils_tar_url := "https://www.freedesktop.org/software/desktop-file-utils/releases/desktop-file-utils-0.27.tar.xz"
	desktopfileutils_cmd_tar := exec.Command("curl", "-L", desktopfileutils_tar_url, "-o", "package.tar.gz")
	err := desktopfileutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	desktopfileutils_zip_url := "https://www.freedesktop.org/software/desktop-file-utils/releases/desktop-file-utils-0.27.tar.xz"
	desktopfileutils_cmd_zip := exec.Command("curl", "-L", desktopfileutils_zip_url, "-o", "package.zip")
	err = desktopfileutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	desktopfileutils_bin_url := "https://www.freedesktop.org/software/desktop-file-utils/releases/desktop-file-utils-0.27.tar.xz"
	desktopfileutils_cmd_bin := exec.Command("curl", "-L", desktopfileutils_bin_url, "-o", "binary.bin")
	err = desktopfileutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	desktopfileutils_src_url := "https://www.freedesktop.org/software/desktop-file-utils/releases/desktop-file-utils-0.27.tar.xz"
	desktopfileutils_cmd_src := exec.Command("curl", "-L", desktopfileutils_src_url, "-o", "source.tar.gz")
	err = desktopfileutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	desktopfileutils_cmd_direct := exec.Command("./binary")
	err = desktopfileutils_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
