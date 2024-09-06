package main

// HicolorIconTheme - Fallback theme for FreeDesktop.org icon themes
// Homepage: https://wiki.freedesktop.org/www/Software/icon-theme/

import (
	"fmt"
	
	"os/exec"
)

func installHicolorIconTheme() {
	// Método 1: Descargar y extraer .tar.gz
	hicoloricontheme_tar_url := "https://icon-theme.freedesktop.org/releases/hicolor-icon-theme-0.18.tar.xz"
	hicoloricontheme_cmd_tar := exec.Command("curl", "-L", hicoloricontheme_tar_url, "-o", "package.tar.gz")
	err := hicoloricontheme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hicoloricontheme_zip_url := "https://icon-theme.freedesktop.org/releases/hicolor-icon-theme-0.18.tar.xz"
	hicoloricontheme_cmd_zip := exec.Command("curl", "-L", hicoloricontheme_zip_url, "-o", "package.zip")
	err = hicoloricontheme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hicoloricontheme_bin_url := "https://icon-theme.freedesktop.org/releases/hicolor-icon-theme-0.18.tar.xz"
	hicoloricontheme_cmd_bin := exec.Command("curl", "-L", hicoloricontheme_bin_url, "-o", "binary.bin")
	err = hicoloricontheme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hicoloricontheme_src_url := "https://icon-theme.freedesktop.org/releases/hicolor-icon-theme-0.18.tar.xz"
	hicoloricontheme_cmd_src := exec.Command("curl", "-L", hicoloricontheme_src_url, "-o", "source.tar.gz")
	err = hicoloricontheme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hicoloricontheme_cmd_direct := exec.Command("./binary")
	err = hicoloricontheme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
}
