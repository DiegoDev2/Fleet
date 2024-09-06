package main

// JpegTurbo - JPEG image codec that aids compression and decompression
// Homepage: https://www.libjpeg-turbo.org/

import (
	"fmt"
	
	"os/exec"
)

func installJpegTurbo() {
	// Método 1: Descargar y extraer .tar.gz
	jpegturbo_tar_url := "https://github.com/libjpeg-turbo/libjpeg-turbo/releases/download/3.0.3/libjpeg-turbo-3.0.3.tar.gz"
	jpegturbo_cmd_tar := exec.Command("curl", "-L", jpegturbo_tar_url, "-o", "package.tar.gz")
	err := jpegturbo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpegturbo_zip_url := "https://github.com/libjpeg-turbo/libjpeg-turbo/releases/download/3.0.3/libjpeg-turbo-3.0.3.zip"
	jpegturbo_cmd_zip := exec.Command("curl", "-L", jpegturbo_zip_url, "-o", "package.zip")
	err = jpegturbo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpegturbo_bin_url := "https://github.com/libjpeg-turbo/libjpeg-turbo/releases/download/3.0.3/libjpeg-turbo-3.0.3.bin"
	jpegturbo_cmd_bin := exec.Command("curl", "-L", jpegturbo_bin_url, "-o", "binary.bin")
	err = jpegturbo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpegturbo_src_url := "https://github.com/libjpeg-turbo/libjpeg-turbo/releases/download/3.0.3/libjpeg-turbo-3.0.3.src.tar.gz"
	jpegturbo_cmd_src := exec.Command("curl", "-L", jpegturbo_src_url, "-o", "source.tar.gz")
	err = jpegturbo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpegturbo_cmd_direct := exec.Command("./binary")
	err = jpegturbo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
