package main

// XcbUtilWm - Client and window-manager helpers for EWMH and ICCCM
// Homepage: https://xcb.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installXcbUtilWm() {
	// Método 1: Descargar y extraer .tar.gz
	xcbutilwm_tar_url := "https://xcb.freedesktop.org/dist/xcb-util-wm-0.4.2.tar.gz"
	xcbutilwm_cmd_tar := exec.Command("curl", "-L", xcbutilwm_tar_url, "-o", "package.tar.gz")
	err := xcbutilwm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbutilwm_zip_url := "https://xcb.freedesktop.org/dist/xcb-util-wm-0.4.2.zip"
	xcbutilwm_cmd_zip := exec.Command("curl", "-L", xcbutilwm_zip_url, "-o", "package.zip")
	err = xcbutilwm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbutilwm_bin_url := "https://xcb.freedesktop.org/dist/xcb-util-wm-0.4.2.bin"
	xcbutilwm_cmd_bin := exec.Command("curl", "-L", xcbutilwm_bin_url, "-o", "binary.bin")
	err = xcbutilwm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbutilwm_src_url := "https://xcb.freedesktop.org/dist/xcb-util-wm-0.4.2.src.tar.gz"
	xcbutilwm_cmd_src := exec.Command("curl", "-L", xcbutilwm_src_url, "-o", "source.tar.gz")
	err = xcbutilwm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbutilwm_cmd_direct := exec.Command("./binary")
	err = xcbutilwm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
}
