package main

// Pdf2svg - PDF converter to SVG
// Homepage: https://cityinthesky.co.uk/opensource/pdf2svg

import (
	"fmt"
	
	"os/exec"
)

func installPdf2svg() {
	// Método 1: Descargar y extraer .tar.gz
	pdf2svg_tar_url := "https://github.com/dawbarton/pdf2svg/archive/refs/tags/v0.2.3.tar.gz"
	pdf2svg_cmd_tar := exec.Command("curl", "-L", pdf2svg_tar_url, "-o", "package.tar.gz")
	err := pdf2svg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdf2svg_zip_url := "https://github.com/dawbarton/pdf2svg/archive/refs/tags/v0.2.3.zip"
	pdf2svg_cmd_zip := exec.Command("curl", "-L", pdf2svg_zip_url, "-o", "package.zip")
	err = pdf2svg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdf2svg_bin_url := "https://github.com/dawbarton/pdf2svg/archive/refs/tags/v0.2.3.bin"
	pdf2svg_cmd_bin := exec.Command("curl", "-L", pdf2svg_bin_url, "-o", "binary.bin")
	err = pdf2svg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdf2svg_src_url := "https://github.com/dawbarton/pdf2svg/archive/refs/tags/v0.2.3.src.tar.gz"
	pdf2svg_cmd_src := exec.Command("curl", "-L", pdf2svg_src_url, "-o", "source.tar.gz")
	err = pdf2svg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdf2svg_cmd_direct := exec.Command("./binary")
	err = pdf2svg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
