package main

// XcbUtilImage - XCB port of Xlib's XImage and XShmImage
// Homepage: https://xcb.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installXcbUtilImage() {
	// Método 1: Descargar y extraer .tar.gz
	xcbutilimage_tar_url := "https://xcb.freedesktop.org/dist/xcb-util-image-0.4.1.tar.gz"
	xcbutilimage_cmd_tar := exec.Command("curl", "-L", xcbutilimage_tar_url, "-o", "package.tar.gz")
	err := xcbutilimage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbutilimage_zip_url := "https://xcb.freedesktop.org/dist/xcb-util-image-0.4.1.zip"
	xcbutilimage_cmd_zip := exec.Command("curl", "-L", xcbutilimage_zip_url, "-o", "package.zip")
	err = xcbutilimage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbutilimage_bin_url := "https://xcb.freedesktop.org/dist/xcb-util-image-0.4.1.bin"
	xcbutilimage_cmd_bin := exec.Command("curl", "-L", xcbutilimage_bin_url, "-o", "binary.bin")
	err = xcbutilimage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbutilimage_src_url := "https://xcb.freedesktop.org/dist/xcb-util-image-0.4.1.src.tar.gz"
	xcbutilimage_cmd_src := exec.Command("curl", "-L", xcbutilimage_src_url, "-o", "source.tar.gz")
	err = xcbutilimage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbutilimage_cmd_direct := exec.Command("./binary")
	err = xcbutilimage_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: xcb-util")
	exec.Command("latte", "install", "xcb-util").Run()
}
