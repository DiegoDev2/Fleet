package main

// Chromaprint - Core component of the AcoustID project (Audio fingerprinting)
// Homepage: https://acoustid.org/chromaprint

import (
	"fmt"
	
	"os/exec"
)

func installChromaprint() {
	// Método 1: Descargar y extraer .tar.gz
	chromaprint_tar_url := "https://github.com/acoustid/chromaprint/releases/download/v1.5.1/chromaprint-1.5.1.tar.gz"
	chromaprint_cmd_tar := exec.Command("curl", "-L", chromaprint_tar_url, "-o", "package.tar.gz")
	err := chromaprint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chromaprint_zip_url := "https://github.com/acoustid/chromaprint/releases/download/v1.5.1/chromaprint-1.5.1.zip"
	chromaprint_cmd_zip := exec.Command("curl", "-L", chromaprint_zip_url, "-o", "package.zip")
	err = chromaprint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chromaprint_bin_url := "https://github.com/acoustid/chromaprint/releases/download/v1.5.1/chromaprint-1.5.1.bin"
	chromaprint_cmd_bin := exec.Command("curl", "-L", chromaprint_bin_url, "-o", "binary.bin")
	err = chromaprint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chromaprint_src_url := "https://github.com/acoustid/chromaprint/releases/download/v1.5.1/chromaprint-1.5.1.src.tar.gz"
	chromaprint_cmd_src := exec.Command("curl", "-L", chromaprint_src_url, "-o", "source.tar.gz")
	err = chromaprint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chromaprint_cmd_direct := exec.Command("./binary")
	err = chromaprint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
}
