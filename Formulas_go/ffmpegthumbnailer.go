package main

// Ffmpegthumbnailer - Create thumbnails for your video files
// Homepage: https://github.com/dirkvdb/ffmpegthumbnailer

import (
	"fmt"
	
	"os/exec"
)

func installFfmpegthumbnailer() {
	// Método 1: Descargar y extraer .tar.gz
	ffmpegthumbnailer_tar_url := "https://github.com/dirkvdb/ffmpegthumbnailer/archive/refs/tags/2.2.2.tar.gz"
	ffmpegthumbnailer_cmd_tar := exec.Command("curl", "-L", ffmpegthumbnailer_tar_url, "-o", "package.tar.gz")
	err := ffmpegthumbnailer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffmpegthumbnailer_zip_url := "https://github.com/dirkvdb/ffmpegthumbnailer/archive/refs/tags/2.2.2.zip"
	ffmpegthumbnailer_cmd_zip := exec.Command("curl", "-L", ffmpegthumbnailer_zip_url, "-o", "package.zip")
	err = ffmpegthumbnailer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffmpegthumbnailer_bin_url := "https://github.com/dirkvdb/ffmpegthumbnailer/archive/refs/tags/2.2.2.bin"
	ffmpegthumbnailer_cmd_bin := exec.Command("curl", "-L", ffmpegthumbnailer_bin_url, "-o", "binary.bin")
	err = ffmpegthumbnailer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffmpegthumbnailer_src_url := "https://github.com/dirkvdb/ffmpegthumbnailer/archive/refs/tags/2.2.2.src.tar.gz"
	ffmpegthumbnailer_cmd_src := exec.Command("curl", "-L", ffmpegthumbnailer_src_url, "-o", "source.tar.gz")
	err = ffmpegthumbnailer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffmpegthumbnailer_cmd_direct := exec.Command("./binary")
	err = ffmpegthumbnailer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
