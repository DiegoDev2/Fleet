package main

// FfmpegAT28 - Play, record, convert, and stream audio and video
// Homepage: https://ffmpeg.org/

import (
	"fmt"
	
	"os/exec"
)

func installFfmpegAT28() {
	// Método 1: Descargar y extraer .tar.gz
	ffmpegat28_tar_url := "https://ffmpeg.org/releases/ffmpeg-2.8.22.tar.xz"
	ffmpegat28_cmd_tar := exec.Command("curl", "-L", ffmpegat28_tar_url, "-o", "package.tar.gz")
	err := ffmpegat28_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffmpegat28_zip_url := "https://ffmpeg.org/releases/ffmpeg-2.8.22.tar.xz"
	ffmpegat28_cmd_zip := exec.Command("curl", "-L", ffmpegat28_zip_url, "-o", "package.zip")
	err = ffmpegat28_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffmpegat28_bin_url := "https://ffmpeg.org/releases/ffmpeg-2.8.22.tar.xz"
	ffmpegat28_cmd_bin := exec.Command("curl", "-L", ffmpegat28_bin_url, "-o", "binary.bin")
	err = ffmpegat28_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffmpegat28_src_url := "https://ffmpeg.org/releases/ffmpeg-2.8.22.tar.xz"
	ffmpegat28_cmd_src := exec.Command("curl", "-L", ffmpegat28_src_url, "-o", "source.tar.gz")
	err = ffmpegat28_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffmpegat28_cmd_direct := exec.Command("./binary")
	err = ffmpegat28_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: texi2html")
exec.Command("latte", "install", "texi2html")
	fmt.Println("Instalando dependencia: yasm")
exec.Command("latte", "install", "yasm")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: frei0r")
exec.Command("latte", "install", "frei0r")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libass")
exec.Command("latte", "install", "libass")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvo-aacenc")
exec.Command("latte", "install", "libvo-aacenc")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: libvpx")
exec.Command("latte", "install", "libvpx")
	fmt.Println("Instalando dependencia: opencore-amr")
exec.Command("latte", "install", "opencore-amr")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: rtmpdump")
exec.Command("latte", "install", "rtmpdump")
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: speex")
exec.Command("latte", "install", "speex")
	fmt.Println("Instalando dependencia: theora")
exec.Command("latte", "install", "theora")
	fmt.Println("Instalando dependencia: x264")
exec.Command("latte", "install", "x264")
	fmt.Println("Instalando dependencia: x265")
exec.Command("latte", "install", "x265")
	fmt.Println("Instalando dependencia: xvid")
exec.Command("latte", "install", "xvid")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
