package main

// Xkeyboardconfig - Keyboard configuration database for the X Window System
// Homepage: https://www.freedesktop.org/wiki/Software/XKeyboardConfig/

import (
	"fmt"
	
	"os/exec"
)

func installXkeyboardconfig() {
	// Método 1: Descargar y extraer .tar.gz
	xkeyboardconfig_tar_url := "https://xorg.freedesktop.org/archive/individual/data/xkeyboard-config/xkeyboard-config-2.42.tar.xz"
	xkeyboardconfig_cmd_tar := exec.Command("curl", "-L", xkeyboardconfig_tar_url, "-o", "package.tar.gz")
	err := xkeyboardconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xkeyboardconfig_zip_url := "https://xorg.freedesktop.org/archive/individual/data/xkeyboard-config/xkeyboard-config-2.42.tar.xz"
	xkeyboardconfig_cmd_zip := exec.Command("curl", "-L", xkeyboardconfig_zip_url, "-o", "package.zip")
	err = xkeyboardconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xkeyboardconfig_bin_url := "https://xorg.freedesktop.org/archive/individual/data/xkeyboard-config/xkeyboard-config-2.42.tar.xz"
	xkeyboardconfig_cmd_bin := exec.Command("curl", "-L", xkeyboardconfig_bin_url, "-o", "binary.bin")
	err = xkeyboardconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xkeyboardconfig_src_url := "https://xorg.freedesktop.org/archive/individual/data/xkeyboard-config/xkeyboard-config-2.42.tar.xz"
	xkeyboardconfig_cmd_src := exec.Command("curl", "-L", xkeyboardconfig_src_url, "-o", "source.tar.gz")
	err = xkeyboardconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xkeyboardconfig_cmd_direct := exec.Command("./binary")
	err = xkeyboardconfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
