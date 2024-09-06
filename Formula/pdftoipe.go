package main

// Pdftoipe - Reads arbitrary PDF files and generates an XML file readable by Ipe
// Homepage: https://github.com/otfried/ipe-tools

import (
	"fmt"
	
	"os/exec"
)

func installPdftoipe() {
	// Método 1: Descargar y extraer .tar.gz
	pdftoipe_tar_url := "https://github.com/otfried/ipe-tools/archive/refs/tags/v7.2.24.1.tar.gz"
	pdftoipe_cmd_tar := exec.Command("curl", "-L", pdftoipe_tar_url, "-o", "package.tar.gz")
	err := pdftoipe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdftoipe_zip_url := "https://github.com/otfried/ipe-tools/archive/refs/tags/v7.2.24.1.zip"
	pdftoipe_cmd_zip := exec.Command("curl", "-L", pdftoipe_zip_url, "-o", "package.zip")
	err = pdftoipe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdftoipe_bin_url := "https://github.com/otfried/ipe-tools/archive/refs/tags/v7.2.24.1.bin"
	pdftoipe_cmd_bin := exec.Command("curl", "-L", pdftoipe_bin_url, "-o", "binary.bin")
	err = pdftoipe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdftoipe_src_url := "https://github.com/otfried/ipe-tools/archive/refs/tags/v7.2.24.1.src.tar.gz"
	pdftoipe_cmd_src := exec.Command("curl", "-L", pdftoipe_src_url, "-o", "source.tar.gz")
	err = pdftoipe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdftoipe_cmd_direct := exec.Command("./binary")
	err = pdftoipe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: poppler")
	exec.Command("latte", "install", "poppler").Run()
}
