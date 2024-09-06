package main

// Libxfont2 - X11 font rasterisation library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxfont2() {
	// Método 1: Descargar y extraer .tar.gz
	libxfont2_tar_url := "https://xorg.freedesktop.org/archive/individual/lib/libXfont2-2.0.7.tar.gz"
	libxfont2_cmd_tar := exec.Command("curl", "-L", libxfont2_tar_url, "-o", "package.tar.gz")
	err := libxfont2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxfont2_zip_url := "https://xorg.freedesktop.org/archive/individual/lib/libXfont2-2.0.7.zip"
	libxfont2_cmd_zip := exec.Command("curl", "-L", libxfont2_zip_url, "-o", "package.zip")
	err = libxfont2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxfont2_bin_url := "https://xorg.freedesktop.org/archive/individual/lib/libXfont2-2.0.7.bin"
	libxfont2_cmd_bin := exec.Command("curl", "-L", libxfont2_bin_url, "-o", "binary.bin")
	err = libxfont2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxfont2_src_url := "https://xorg.freedesktop.org/archive/individual/lib/libXfont2-2.0.7.src.tar.gz"
	libxfont2_cmd_src := exec.Command("curl", "-L", libxfont2_src_url, "-o", "source.tar.gz")
	err = libxfont2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxfont2_cmd_direct := exec.Command("./binary")
	err = libxfont2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
	fmt.Println("Instalando dependencia: xtrans")
exec.Command("latte", "install", "xtrans")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libfontenc")
exec.Command("latte", "install", "libfontenc")
}
