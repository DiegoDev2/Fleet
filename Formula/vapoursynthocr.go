package main

// VapoursynthOcr - VapourSynth filters - Tesseract OCR filter
// Homepage: https://www.vapoursynth.com

import (
	"fmt"
	
	"os/exec"
)

func installVapoursynthOcr() {
	// Método 1: Descargar y extraer .tar.gz
	vapoursynthocr_tar_url := "https://github.com/vapoursynth/vs-ocr/archive/refs/tags/R3.tar.gz"
	vapoursynthocr_cmd_tar := exec.Command("curl", "-L", vapoursynthocr_tar_url, "-o", "package.tar.gz")
	err := vapoursynthocr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vapoursynthocr_zip_url := "https://github.com/vapoursynth/vs-ocr/archive/refs/tags/R3.zip"
	vapoursynthocr_cmd_zip := exec.Command("curl", "-L", vapoursynthocr_zip_url, "-o", "package.zip")
	err = vapoursynthocr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vapoursynthocr_bin_url := "https://github.com/vapoursynth/vs-ocr/archive/refs/tags/R3.bin"
	vapoursynthocr_cmd_bin := exec.Command("curl", "-L", vapoursynthocr_bin_url, "-o", "binary.bin")
	err = vapoursynthocr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vapoursynthocr_src_url := "https://github.com/vapoursynth/vs-ocr/archive/refs/tags/R3.src.tar.gz"
	vapoursynthocr_cmd_src := exec.Command("curl", "-L", vapoursynthocr_src_url, "-o", "source.tar.gz")
	err = vapoursynthocr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vapoursynthocr_cmd_direct := exec.Command("./binary")
	err = vapoursynthocr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: tesseract")
	exec.Command("latte", "install", "tesseract").Run()
	fmt.Println("Instalando dependencia: vapoursynth")
	exec.Command("latte", "install", "vapoursynth").Run()
}
