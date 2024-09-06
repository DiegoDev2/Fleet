package main

// XcbUtilKeysyms - Standard X constants and conversion to/from keycodes
// Homepage: https://xcb.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installXcbUtilKeysyms() {
	// Método 1: Descargar y extraer .tar.gz
	xcbutilkeysyms_tar_url := "https://xcb.freedesktop.org/dist/xcb-util-keysyms-0.4.1.tar.gz"
	xcbutilkeysyms_cmd_tar := exec.Command("curl", "-L", xcbutilkeysyms_tar_url, "-o", "package.tar.gz")
	err := xcbutilkeysyms_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbutilkeysyms_zip_url := "https://xcb.freedesktop.org/dist/xcb-util-keysyms-0.4.1.zip"
	xcbutilkeysyms_cmd_zip := exec.Command("curl", "-L", xcbutilkeysyms_zip_url, "-o", "package.zip")
	err = xcbutilkeysyms_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbutilkeysyms_bin_url := "https://xcb.freedesktop.org/dist/xcb-util-keysyms-0.4.1.bin"
	xcbutilkeysyms_cmd_bin := exec.Command("curl", "-L", xcbutilkeysyms_bin_url, "-o", "binary.bin")
	err = xcbutilkeysyms_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbutilkeysyms_src_url := "https://xcb.freedesktop.org/dist/xcb-util-keysyms-0.4.1.src.tar.gz"
	xcbutilkeysyms_cmd_src := exec.Command("curl", "-L", xcbutilkeysyms_src_url, "-o", "source.tar.gz")
	err = xcbutilkeysyms_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbutilkeysyms_cmd_direct := exec.Command("./binary")
	err = xcbutilkeysyms_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
}
