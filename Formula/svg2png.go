package main

// Svg2png - SVG to PNG converter
// Homepage: https://cairographics.org/

import (
	"fmt"
	
	"os/exec"
)

func installSvg2png() {
	// Método 1: Descargar y extraer .tar.gz
	svg2png_tar_url := "https://cairographics.org/snapshots/svg2png-0.1.3.tar.gz"
	svg2png_cmd_tar := exec.Command("curl", "-L", svg2png_tar_url, "-o", "package.tar.gz")
	err := svg2png_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svg2png_zip_url := "https://cairographics.org/snapshots/svg2png-0.1.3.zip"
	svg2png_cmd_zip := exec.Command("curl", "-L", svg2png_zip_url, "-o", "package.zip")
	err = svg2png_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svg2png_bin_url := "https://cairographics.org/snapshots/svg2png-0.1.3.bin"
	svg2png_cmd_bin := exec.Command("curl", "-L", svg2png_bin_url, "-o", "binary.bin")
	err = svg2png_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svg2png_src_url := "https://cairographics.org/snapshots/svg2png-0.1.3.src.tar.gz"
	svg2png_cmd_src := exec.Command("curl", "-L", svg2png_src_url, "-o", "source.tar.gz")
	err = svg2png_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svg2png_cmd_direct := exec.Command("./binary")
	err = svg2png_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsvg")
	exec.Command("latte", "install", "libsvg").Run()
	fmt.Println("Instalando dependencia: libsvg-cairo")
	exec.Command("latte", "install", "libsvg-cairo").Run()
}
