package main

// Svg2pdf - Renders SVG images to a PDF file (using Cairo)
// Homepage: https://cairographics.org/

import (
	"fmt"
	
	"os/exec"
)

func installSvg2pdf() {
	// Método 1: Descargar y extraer .tar.gz
	svg2pdf_tar_url := "https://cairographics.org/snapshots/svg2pdf-0.1.3.tar.gz"
	svg2pdf_cmd_tar := exec.Command("curl", "-L", svg2pdf_tar_url, "-o", "package.tar.gz")
	err := svg2pdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	svg2pdf_zip_url := "https://cairographics.org/snapshots/svg2pdf-0.1.3.zip"
	svg2pdf_cmd_zip := exec.Command("curl", "-L", svg2pdf_zip_url, "-o", "package.zip")
	err = svg2pdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	svg2pdf_bin_url := "https://cairographics.org/snapshots/svg2pdf-0.1.3.bin"
	svg2pdf_cmd_bin := exec.Command("curl", "-L", svg2pdf_bin_url, "-o", "binary.bin")
	err = svg2pdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	svg2pdf_src_url := "https://cairographics.org/snapshots/svg2pdf-0.1.3.src.tar.gz"
	svg2pdf_cmd_src := exec.Command("curl", "-L", svg2pdf_src_url, "-o", "source.tar.gz")
	err = svg2pdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	svg2pdf_cmd_direct := exec.Command("./binary")
	err = svg2pdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsvg-cairo")
exec.Command("latte", "install", "libsvg-cairo")
}
