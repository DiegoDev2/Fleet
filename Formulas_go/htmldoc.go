package main

// Htmldoc - Convert HTML to PDF or PostScript
// Homepage: https://www.msweet.org/htmldoc/

import (
	"fmt"
	
	"os/exec"
)

func installHtmldoc() {
	// Método 1: Descargar y extraer .tar.gz
	htmldoc_tar_url := "https://github.com/michaelrsweet/htmldoc/archive/refs/tags/v1.9.18.tar.gz"
	htmldoc_cmd_tar := exec.Command("curl", "-L", htmldoc_tar_url, "-o", "package.tar.gz")
	err := htmldoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htmldoc_zip_url := "https://github.com/michaelrsweet/htmldoc/archive/refs/tags/v1.9.18.zip"
	htmldoc_cmd_zip := exec.Command("curl", "-L", htmldoc_zip_url, "-o", "package.zip")
	err = htmldoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htmldoc_bin_url := "https://github.com/michaelrsweet/htmldoc/archive/refs/tags/v1.9.18.bin"
	htmldoc_cmd_bin := exec.Command("curl", "-L", htmldoc_bin_url, "-o", "binary.bin")
	err = htmldoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htmldoc_src_url := "https://github.com/michaelrsweet/htmldoc/archive/refs/tags/v1.9.18.src.tar.gz"
	htmldoc_cmd_src := exec.Command("curl", "-L", htmldoc_src_url, "-o", "source.tar.gz")
	err = htmldoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htmldoc_cmd_direct := exec.Command("./binary")
	err = htmldoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
}
