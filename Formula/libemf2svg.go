package main

// Libemf2svg - Microsoft (MS) EMF to SVG conversion library
// Homepage: https://github.com/kakwa/libemf2svg

import (
	"fmt"
	
	"os/exec"
)

func installLibemf2svg() {
	// Método 1: Descargar y extraer .tar.gz
	libemf2svg_tar_url := "https://github.com/kakwa/libemf2svg/archive/refs/tags/1.1.0.tar.gz"
	libemf2svg_cmd_tar := exec.Command("curl", "-L", libemf2svg_tar_url, "-o", "package.tar.gz")
	err := libemf2svg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libemf2svg_zip_url := "https://github.com/kakwa/libemf2svg/archive/refs/tags/1.1.0.zip"
	libemf2svg_cmd_zip := exec.Command("curl", "-L", libemf2svg_zip_url, "-o", "package.zip")
	err = libemf2svg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libemf2svg_bin_url := "https://github.com/kakwa/libemf2svg/archive/refs/tags/1.1.0.bin"
	libemf2svg_cmd_bin := exec.Command("curl", "-L", libemf2svg_bin_url, "-o", "binary.bin")
	err = libemf2svg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libemf2svg_src_url := "https://github.com/kakwa/libemf2svg/archive/refs/tags/1.1.0.src.tar.gz"
	libemf2svg_cmd_src := exec.Command("curl", "-L", libemf2svg_src_url, "-o", "source.tar.gz")
	err = libemf2svg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libemf2svg_cmd_direct := exec.Command("./binary")
	err = libemf2svg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: argp-standalone")
	exec.Command("latte", "install", "argp-standalone").Run()
}
