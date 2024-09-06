package main

// Ffms2 - Libav/ffmpeg based source library and Avisynth plugin
// Homepage: https://github.com/FFMS/ffms2

import (
	"fmt"
	
	"os/exec"
)

func installFfms2() {
	// Método 1: Descargar y extraer .tar.gz
	ffms2_tar_url := "https://github.com/FFMS/ffms2/archive/refs/tags/5.0.tar.gz"
	ffms2_cmd_tar := exec.Command("curl", "-L", ffms2_tar_url, "-o", "package.tar.gz")
	err := ffms2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffms2_zip_url := "https://github.com/FFMS/ffms2/archive/refs/tags/5.0.zip"
	ffms2_cmd_zip := exec.Command("curl", "-L", ffms2_zip_url, "-o", "package.zip")
	err = ffms2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffms2_bin_url := "https://github.com/FFMS/ffms2/archive/refs/tags/5.0.bin"
	ffms2_cmd_bin := exec.Command("curl", "-L", ffms2_bin_url, "-o", "binary.bin")
	err = ffms2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffms2_src_url := "https://github.com/FFMS/ffms2/archive/refs/tags/5.0.src.tar.gz"
	ffms2_cmd_src := exec.Command("curl", "-L", ffms2_src_url, "-o", "source.tar.gz")
	err = ffms2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffms2_cmd_direct := exec.Command("./binary")
	err = ffms2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
}
