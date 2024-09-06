package main

// VideoCompare - Split screen video comparison tool using FFmpeg and SDL2
// Homepage: https://github.com/pixop/video-compare

import (
	"fmt"
	
	"os/exec"
)

func installVideoCompare() {
	// Método 1: Descargar y extraer .tar.gz
	videocompare_tar_url := "https://github.com/pixop/video-compare/archive/refs/tags/20240818.tar.gz"
	videocompare_cmd_tar := exec.Command("curl", "-L", videocompare_tar_url, "-o", "package.tar.gz")
	err := videocompare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	videocompare_zip_url := "https://github.com/pixop/video-compare/archive/refs/tags/20240818.zip"
	videocompare_cmd_zip := exec.Command("curl", "-L", videocompare_zip_url, "-o", "package.zip")
	err = videocompare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	videocompare_bin_url := "https://github.com/pixop/video-compare/archive/refs/tags/20240818.bin"
	videocompare_cmd_bin := exec.Command("curl", "-L", videocompare_bin_url, "-o", "binary.bin")
	err = videocompare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	videocompare_src_url := "https://github.com/pixop/video-compare/archive/refs/tags/20240818.src.tar.gz"
	videocompare_cmd_src := exec.Command("curl", "-L", videocompare_src_url, "-o", "source.tar.gz")
	err = videocompare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	videocompare_cmd_direct := exec.Command("./binary")
	err = videocompare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_ttf")
	exec.Command("latte", "install", "sdl2_ttf").Run()
}
