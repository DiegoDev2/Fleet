package main

// Pdf2image - Convert PDFs to images
// Homepage: https://code.google.com/p/pdf2image/

import (
	"fmt"
	
	"os/exec"
)

func installPdf2image() {
	// Método 1: Descargar y extraer .tar.gz
	pdf2image_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/pdf2image/pdf2image-0.53-source.tar.gz"
	pdf2image_cmd_tar := exec.Command("curl", "-L", pdf2image_tar_url, "-o", "package.tar.gz")
	err := pdf2image_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdf2image_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/pdf2image/pdf2image-0.53-source.zip"
	pdf2image_cmd_zip := exec.Command("curl", "-L", pdf2image_zip_url, "-o", "package.zip")
	err = pdf2image_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdf2image_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/pdf2image/pdf2image-0.53-source.bin"
	pdf2image_cmd_bin := exec.Command("curl", "-L", pdf2image_bin_url, "-o", "binary.bin")
	err = pdf2image_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdf2image_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/pdf2image/pdf2image-0.53-source.src.tar.gz"
	pdf2image_cmd_src := exec.Command("curl", "-L", pdf2image_src_url, "-o", "source.tar.gz")
	err = pdf2image_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdf2image_cmd_direct := exec.Command("./binary")
	err = pdf2image_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
}
