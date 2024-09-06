package main

// Ocrmypdf - Adds an OCR text layer to scanned PDF files
// Homepage: https://ocrmypdf.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installOcrmypdf() {
	// Método 1: Descargar y extraer .tar.gz
	ocrmypdf_tar_url := "https://files.pythonhosted.org/packages/2c/85/439deb5b418261dbfd844dbae8c1d8713c9c7bc5dd36a59577a77d7ffbbb/ocrmypdf-16.5.0.tar.gz"
	ocrmypdf_cmd_tar := exec.Command("curl", "-L", ocrmypdf_tar_url, "-o", "package.tar.gz")
	err := ocrmypdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ocrmypdf_zip_url := "https://files.pythonhosted.org/packages/2c/85/439deb5b418261dbfd844dbae8c1d8713c9c7bc5dd36a59577a77d7ffbbb/ocrmypdf-16.5.0.zip"
	ocrmypdf_cmd_zip := exec.Command("curl", "-L", ocrmypdf_zip_url, "-o", "package.zip")
	err = ocrmypdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ocrmypdf_bin_url := "https://files.pythonhosted.org/packages/2c/85/439deb5b418261dbfd844dbae8c1d8713c9c7bc5dd36a59577a77d7ffbbb/ocrmypdf-16.5.0.bin"
	ocrmypdf_cmd_bin := exec.Command("curl", "-L", ocrmypdf_bin_url, "-o", "binary.bin")
	err = ocrmypdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ocrmypdf_src_url := "https://files.pythonhosted.org/packages/2c/85/439deb5b418261dbfd844dbae8c1d8713c9c7bc5dd36a59577a77d7ffbbb/ocrmypdf-16.5.0.src.tar.gz"
	ocrmypdf_cmd_src := exec.Command("curl", "-L", ocrmypdf_src_url, "-o", "source.tar.gz")
	err = ocrmypdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ocrmypdf_cmd_direct := exec.Command("./binary")
	err = ocrmypdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: img2pdf")
exec.Command("latte", "install", "img2pdf")
	fmt.Println("Instalando dependencia: jbig2enc")
exec.Command("latte", "install", "jbig2enc")
	fmt.Println("Instalando dependencia: libheif")
exec.Command("latte", "install", "libheif")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: pngquant")
exec.Command("latte", "install", "pngquant")
	fmt.Println("Instalando dependencia: pybind11")
exec.Command("latte", "install", "pybind11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: qpdf")
exec.Command("latte", "install", "qpdf")
	fmt.Println("Instalando dependencia: tesseract")
exec.Command("latte", "install", "tesseract")
	fmt.Println("Instalando dependencia: unpaper")
exec.Command("latte", "install", "unpaper")
}
