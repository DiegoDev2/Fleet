package main

// XcbUtil - Additional extensions to the XCB library
// Homepage: https://xcb.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installXcbUtil() {
	// Método 1: Descargar y extraer .tar.gz
	xcbutil_tar_url := "https://xcb.freedesktop.org/dist/xcb-util-0.4.1.tar.xz"
	xcbutil_cmd_tar := exec.Command("curl", "-L", xcbutil_tar_url, "-o", "package.tar.gz")
	err := xcbutil_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xcbutil_zip_url := "https://xcb.freedesktop.org/dist/xcb-util-0.4.1.tar.xz"
	xcbutil_cmd_zip := exec.Command("curl", "-L", xcbutil_zip_url, "-o", "package.zip")
	err = xcbutil_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xcbutil_bin_url := "https://xcb.freedesktop.org/dist/xcb-util-0.4.1.tar.xz"
	xcbutil_cmd_bin := exec.Command("curl", "-L", xcbutil_bin_url, "-o", "binary.bin")
	err = xcbutil_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xcbutil_src_url := "https://xcb.freedesktop.org/dist/xcb-util-0.4.1.tar.xz"
	xcbutil_cmd_src := exec.Command("curl", "-L", xcbutil_src_url, "-o", "source.tar.gz")
	err = xcbutil_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xcbutil_cmd_direct := exec.Command("./binary")
	err = xcbutil_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
}
