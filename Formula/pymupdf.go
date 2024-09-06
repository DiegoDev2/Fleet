package main

// Pymupdf - Python bindings for the PDF toolkit and renderer MuPDF
// Homepage: https://pymupdf.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installPymupdf() {
	// Método 1: Descargar y extraer .tar.gz
	pymupdf_tar_url := "https://files.pythonhosted.org/packages/09/48/862dcbe3cc3f11394c2fc9c5021bf8023b4c917213b63553fb8f15764c95/PyMuPDF-1.24.9.tar.gz"
	pymupdf_cmd_tar := exec.Command("curl", "-L", pymupdf_tar_url, "-o", "package.tar.gz")
	err := pymupdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pymupdf_zip_url := "https://files.pythonhosted.org/packages/09/48/862dcbe3cc3f11394c2fc9c5021bf8023b4c917213b63553fb8f15764c95/PyMuPDF-1.24.9.zip"
	pymupdf_cmd_zip := exec.Command("curl", "-L", pymupdf_zip_url, "-o", "package.zip")
	err = pymupdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pymupdf_bin_url := "https://files.pythonhosted.org/packages/09/48/862dcbe3cc3f11394c2fc9c5021bf8023b4c917213b63553fb8f15764c95/PyMuPDF-1.24.9.bin"
	pymupdf_cmd_bin := exec.Command("curl", "-L", pymupdf_bin_url, "-o", "binary.bin")
	err = pymupdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pymupdf_src_url := "https://files.pythonhosted.org/packages/09/48/862dcbe3cc3f11394c2fc9c5021bf8023b4c917213b63553fb8f15764c95/PyMuPDF-1.24.9.src.tar.gz"
	pymupdf_cmd_src := exec.Command("curl", "-L", pymupdf_src_url, "-o", "source.tar.gz")
	err = pymupdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pymupdf_cmd_direct := exec.Command("./binary")
	err = pymupdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: mupdf")
	exec.Command("latte", "install", "mupdf").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
