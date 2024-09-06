package main

// TesseractLang - Enables extra languages support for Tesseract
// Homepage: https://github.com/tesseract-ocr/tessdata_fast/

import (
	"fmt"
	
	"os/exec"
)

func installTesseractLang() {
	// Método 1: Descargar y extraer .tar.gz
	tesseractlang_tar_url := "https://github.com/tesseract-ocr/tessdata_fast/archive/refs/tags/4.1.0.tar.gz"
	tesseractlang_cmd_tar := exec.Command("curl", "-L", tesseractlang_tar_url, "-o", "package.tar.gz")
	err := tesseractlang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tesseractlang_zip_url := "https://github.com/tesseract-ocr/tessdata_fast/archive/refs/tags/4.1.0.zip"
	tesseractlang_cmd_zip := exec.Command("curl", "-L", tesseractlang_zip_url, "-o", "package.zip")
	err = tesseractlang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tesseractlang_bin_url := "https://github.com/tesseract-ocr/tessdata_fast/archive/refs/tags/4.1.0.bin"
	tesseractlang_cmd_bin := exec.Command("curl", "-L", tesseractlang_bin_url, "-o", "binary.bin")
	err = tesseractlang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tesseractlang_src_url := "https://github.com/tesseract-ocr/tessdata_fast/archive/refs/tags/4.1.0.src.tar.gz"
	tesseractlang_cmd_src := exec.Command("curl", "-L", tesseractlang_src_url, "-o", "source.tar.gz")
	err = tesseractlang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tesseractlang_cmd_direct := exec.Command("./binary")
	err = tesseractlang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tesseract")
exec.Command("latte", "install", "tesseract")
}
