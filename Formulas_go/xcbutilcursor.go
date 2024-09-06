package main

// XcbUtilCursor - XCB cursor library (replacement for libXcursor)
// Homepage: https://xcb.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installXcbUtilCursor() {
	// Método 1: Descargar y extraer .tar.gz
	xcbutilcursor_tar_url := "https://xcb.freedesktop.org/dist/xcb-util-cursor-0.1.4.tar.xz"
	xcbutilcursor_cmd_tar := exec.Command("curl", "-L", xcbutilcursor_tar_url, "-o", "package.tar.gz")
	err := xcbutilcursor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbutilcursor_zip_url := "https://xcb.freedesktop.org/dist/xcb-util-cursor-0.1.4.tar.xz"
	xcbutilcursor_cmd_zip := exec.Command("curl", "-L", xcbutilcursor_zip_url, "-o", "package.zip")
	err = xcbutilcursor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbutilcursor_bin_url := "https://xcb.freedesktop.org/dist/xcb-util-cursor-0.1.4.tar.xz"
	xcbutilcursor_cmd_bin := exec.Command("curl", "-L", xcbutilcursor_bin_url, "-o", "binary.bin")
	err = xcbutilcursor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbutilcursor_src_url := "https://xcb.freedesktop.org/dist/xcb-util-cursor-0.1.4.tar.xz"
	xcbutilcursor_cmd_src := exec.Command("curl", "-L", xcbutilcursor_src_url, "-o", "source.tar.gz")
	err = xcbutilcursor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbutilcursor_cmd_direct := exec.Command("./binary")
	err = xcbutilcursor_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: xcb-util")
exec.Command("latte", "install", "xcb-util")
	fmt.Println("Instalando dependencia: xcb-util-image")
exec.Command("latte", "install", "xcb-util-image")
	fmt.Println("Instalando dependencia: xcb-util-renderutil")
exec.Command("latte", "install", "xcb-util-renderutil")
}
