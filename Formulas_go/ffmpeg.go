package main

// Ffmpeg - Play, record, convert, and stream audio and video
// Homepage: https://ffmpeg.org/

import (
	"fmt"
	
	"os/exec"
)

func installFfmpeg() {
	// Método 1: Descargar y extraer .tar.gz
	ffmpeg_tar_url := "https://ffmpeg.org/releases/ffmpeg-7.0.2.tar.xz"
	ffmpeg_cmd_tar := exec.Command("curl", "-L", ffmpeg_tar_url, "-o", "package.tar.gz")
	err := ffmpeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffmpeg_zip_url := "https://ffmpeg.org/releases/ffmpeg-7.0.2.tar.xz"
	ffmpeg_cmd_zip := exec.Command("curl", "-L", ffmpeg_zip_url, "-o", "package.zip")
	err = ffmpeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffmpeg_bin_url := "https://ffmpeg.org/releases/ffmpeg-7.0.2.tar.xz"
	ffmpeg_cmd_bin := exec.Command("curl", "-L", ffmpeg_bin_url, "-o", "binary.bin")
	err = ffmpeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffmpeg_src_url := "https://ffmpeg.org/releases/ffmpeg-7.0.2.tar.xz"
	ffmpeg_cmd_src := exec.Command("curl", "-L", ffmpeg_src_url, "-o", "source.tar.gz")
	err = ffmpeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffmpeg_cmd_direct := exec.Command("./binary")
	err = ffmpeg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: aom")
exec.Command("latte", "install", "aom")
	fmt.Println("Instalando dependencia: aribb24")
exec.Command("latte", "install", "aribb24")
	fmt.Println("Instalando dependencia: dav1d")
exec.Command("latte", "install", "dav1d")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: frei0r")
exec.Command("latte", "install", "frei0r")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: jpeg-xl")
exec.Command("latte", "install", "jpeg-xl")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libass")
exec.Command("latte", "install", "libass")
	fmt.Println("Instalando dependencia: libbluray")
exec.Command("latte", "install", "libbluray")
	fmt.Println("Instalando dependencia: librist")
exec.Command("latte", "install", "librist")
	fmt.Println("Instalando dependencia: libsoxr")
exec.Command("latte", "install", "libsoxr")
	fmt.Println("Instalando dependencia: libssh")
exec.Command("latte", "install", "libssh")
	fmt.Println("Instalando dependencia: libvidstab")
exec.Command("latte", "install", "libvidstab")
	fmt.Println("Instalando dependencia: libvmaf")
exec.Command("latte", "install", "libvmaf")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: libvpx")
exec.Command("latte", "install", "libvpx")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: opencore-amr")
exec.Command("latte", "install", "opencore-amr")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: rav1e")
exec.Command("latte", "install", "rav1e")
	fmt.Println("Instalando dependencia: rubberband")
exec.Command("latte", "install", "rubberband")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: snappy")
exec.Command("latte", "install", "snappy")
	fmt.Println("Instalando dependencia: speex")
exec.Command("latte", "install", "speex")
	fmt.Println("Instalando dependencia: srt")
exec.Command("latte", "install", "srt")
	fmt.Println("Instalando dependencia: svt-av1")
exec.Command("latte", "install", "svt-av1")
	fmt.Println("Instalando dependencia: tesseract")
exec.Command("latte", "install", "tesseract")
	fmt.Println("Instalando dependencia: theora")
exec.Command("latte", "install", "theora")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: x264")
exec.Command("latte", "install", "x264")
	fmt.Println("Instalando dependencia: x265")
exec.Command("latte", "install", "x265")
	fmt.Println("Instalando dependencia: xvid")
exec.Command("latte", "install", "xvid")
	fmt.Println("Instalando dependencia: xz")
exec.Command("latte", "install", "xz")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
	fmt.Println("Instalando dependencia: zimg")
exec.Command("latte", "install", "zimg")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: libxext")
exec.Command("latte", "install", "libxext")
	fmt.Println("Instalando dependencia: libxv")
exec.Command("latte", "install", "libxv")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
