package main

// Img2pdf - Convert images to PDF via direct JPEG inclusion
// Homepage: https://gitlab.mister-muffin.de/josch/img2pdf

import (
	"fmt"
	
	"os/exec"
)

func installImg2pdf() {
	// Método 1: Descargar y extraer .tar.gz
	img2pdf_tar_url := "https://files.pythonhosted.org/packages/36/92/6ac4d61951ba507b499f674c90dfa7b48fa776b56f6f068507f8751c03f1/img2pdf-0.5.1.tar.gz"
	img2pdf_cmd_tar := exec.Command("curl", "-L", img2pdf_tar_url, "-o", "package.tar.gz")
	err := img2pdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	img2pdf_zip_url := "https://files.pythonhosted.org/packages/36/92/6ac4d61951ba507b499f674c90dfa7b48fa776b56f6f068507f8751c03f1/img2pdf-0.5.1.zip"
	img2pdf_cmd_zip := exec.Command("curl", "-L", img2pdf_zip_url, "-o", "package.zip")
	err = img2pdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	img2pdf_bin_url := "https://files.pythonhosted.org/packages/36/92/6ac4d61951ba507b499f674c90dfa7b48fa776b56f6f068507f8751c03f1/img2pdf-0.5.1.bin"
	img2pdf_cmd_bin := exec.Command("curl", "-L", img2pdf_bin_url, "-o", "binary.bin")
	err = img2pdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	img2pdf_src_url := "https://files.pythonhosted.org/packages/36/92/6ac4d61951ba507b499f674c90dfa7b48fa776b56f6f068507f8751c03f1/img2pdf-0.5.1.src.tar.gz"
	img2pdf_cmd_src := exec.Command("curl", "-L", img2pdf_src_url, "-o", "source.tar.gz")
	err = img2pdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	img2pdf_cmd_direct := exec.Command("./binary")
	err = img2pdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qpdf")
exec.Command("latte", "install", "qpdf")
}
