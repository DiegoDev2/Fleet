package main

// Tesseract - OCR (Optical Character Recognition) engine
// Homepage: https://github.com/tesseract-ocr/

import (
	"fmt"
	
	"os/exec"
)

func installTesseract() {
	// Método 1: Descargar y extraer .tar.gz
	tesseract_tar_url := "https://github.com/tesseract-ocr/tesseract/archive/refs/tags/5.4.1.tar.gz"
	tesseract_cmd_tar := exec.Command("curl", "-L", tesseract_tar_url, "-o", "package.tar.gz")
	err := tesseract_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tesseract_zip_url := "https://github.com/tesseract-ocr/tesseract/archive/refs/tags/5.4.1.zip"
	tesseract_cmd_zip := exec.Command("curl", "-L", tesseract_zip_url, "-o", "package.zip")
	err = tesseract_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tesseract_bin_url := "https://github.com/tesseract-ocr/tesseract/archive/refs/tags/5.4.1.bin"
	tesseract_cmd_bin := exec.Command("curl", "-L", tesseract_bin_url, "-o", "binary.bin")
	err = tesseract_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tesseract_src_url := "https://github.com/tesseract-ocr/tesseract/archive/refs/tags/5.4.1.src.tar.gz"
	tesseract_cmd_src := exec.Command("curl", "-L", tesseract_src_url, "-o", "source.tar.gz")
	err = tesseract_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tesseract_cmd_direct := exec.Command("./binary")
	err = tesseract_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: leptonica")
	exec.Command("latte", "install", "leptonica").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
