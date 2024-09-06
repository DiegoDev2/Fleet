package main

// FfmpegAT6 - Play, record, convert, and stream audio and video
// Homepage: https://ffmpeg.org/

import (
	"fmt"
	
	"os/exec"
)

func installFfmpegAT6() {
	// Método 1: Descargar y extraer .tar.gz
	ffmpegat6_tar_url := "https://ffmpeg.org/releases/ffmpeg-6.1.2.tar.xz"
	ffmpegat6_cmd_tar := exec.Command("curl", "-L", ffmpegat6_tar_url, "-o", "package.tar.gz")
	err := ffmpegat6_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffmpegat6_zip_url := "https://ffmpeg.org/releases/ffmpeg-6.1.2.tar.xz"
	ffmpegat6_cmd_zip := exec.Command("curl", "-L", ffmpegat6_zip_url, "-o", "package.zip")
	err = ffmpegat6_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffmpegat6_bin_url := "https://ffmpeg.org/releases/ffmpeg-6.1.2.tar.xz"
	ffmpegat6_cmd_bin := exec.Command("curl", "-L", ffmpegat6_bin_url, "-o", "binary.bin")
	err = ffmpegat6_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffmpegat6_src_url := "https://ffmpeg.org/releases/ffmpeg-6.1.2.tar.xz"
	ffmpegat6_cmd_src := exec.Command("curl", "-L", ffmpegat6_src_url, "-o", "source.tar.gz")
	err = ffmpegat6_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffmpegat6_cmd_direct := exec.Command("./binary")
	err = ffmpegat6_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: aom")
	exec.Command("latte", "install", "aom").Run()
	fmt.Println("Instalando dependencia: aribb24")
	exec.Command("latte", "install", "aribb24").Run()
	fmt.Println("Instalando dependencia: dav1d")
	exec.Command("latte", "install", "dav1d").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: frei0r")
	exec.Command("latte", "install", "frei0r").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: jpeg-xl")
	exec.Command("latte", "install", "jpeg-xl").Run()
	fmt.Println("Instalando dependencia: lame")
	exec.Command("latte", "install", "lame").Run()
	fmt.Println("Instalando dependencia: libass")
	exec.Command("latte", "install", "libass").Run()
	fmt.Println("Instalando dependencia: libbluray")
	exec.Command("latte", "install", "libbluray").Run()
	fmt.Println("Instalando dependencia: librist")
	exec.Command("latte", "install", "librist").Run()
	fmt.Println("Instalando dependencia: libsoxr")
	exec.Command("latte", "install", "libsoxr").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: libvidstab")
	exec.Command("latte", "install", "libvidstab").Run()
	fmt.Println("Instalando dependencia: libvmaf")
	exec.Command("latte", "install", "libvmaf").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libvpx")
	exec.Command("latte", "install", "libvpx").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: opencore-amr")
	exec.Command("latte", "install", "opencore-amr").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: openvino")
	exec.Command("latte", "install", "openvino").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: rav1e")
	exec.Command("latte", "install", "rav1e").Run()
	fmt.Println("Instalando dependencia: rubberband")
	exec.Command("latte", "install", "rubberband").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: speex")
	exec.Command("latte", "install", "speex").Run()
	fmt.Println("Instalando dependencia: srt")
	exec.Command("latte", "install", "srt").Run()
	fmt.Println("Instalando dependencia: svt-av1")
	exec.Command("latte", "install", "svt-av1").Run()
	fmt.Println("Instalando dependencia: tesseract")
	exec.Command("latte", "install", "tesseract").Run()
	fmt.Println("Instalando dependencia: theora")
	exec.Command("latte", "install", "theora").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
	fmt.Println("Instalando dependencia: x264")
	exec.Command("latte", "install", "x264").Run()
	fmt.Println("Instalando dependencia: x265")
	exec.Command("latte", "install", "x265").Run()
	fmt.Println("Instalando dependencia: xvid")
	exec.Command("latte", "install", "xvid").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
	fmt.Println("Instalando dependencia: zimg")
	exec.Command("latte", "install", "zimg").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: pugixml")
	exec.Command("latte", "install", "pugixml").Run()
	fmt.Println("Instalando dependencia: tbb")
	exec.Command("latte", "install", "tbb").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxv")
	exec.Command("latte", "install", "libxv").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
}
