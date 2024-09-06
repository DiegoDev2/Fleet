package main

// Libxfont - X.Org: Core of the legacy X11 font system
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxfont() {
	// Método 1: Descargar y extraer .tar.gz
	libxfont_tar_url := "https://www.x.org/archive/individual/lib/libXfont-1.5.4.tar.bz2"
	libxfont_cmd_tar := exec.Command("curl", "-L", libxfont_tar_url, "-o", "package.tar.gz")
	err := libxfont_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxfont_zip_url := "https://www.x.org/archive/individual/lib/libXfont-1.5.4.tar.bz2"
	libxfont_cmd_zip := exec.Command("curl", "-L", libxfont_zip_url, "-o", "package.zip")
	err = libxfont_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxfont_bin_url := "https://www.x.org/archive/individual/lib/libXfont-1.5.4.tar.bz2"
	libxfont_cmd_bin := exec.Command("curl", "-L", libxfont_bin_url, "-o", "binary.bin")
	err = libxfont_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxfont_src_url := "https://www.x.org/archive/individual/lib/libXfont-1.5.4.tar.bz2"
	libxfont_cmd_src := exec.Command("curl", "-L", libxfont_src_url, "-o", "source.tar.gz")
	err = libxfont_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxfont_cmd_direct := exec.Command("./binary")
	err = libxfont_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
	fmt.Println("Instalando dependencia: xtrans")
exec.Command("latte", "install", "xtrans")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libfontenc")
exec.Command("latte", "install", "libfontenc")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
