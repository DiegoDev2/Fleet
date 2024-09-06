package main

// Mlt - Author, manage, and run multitrack audio/video compositions
// Homepage: https://www.mltframework.org/

import (
	"fmt"
	
	"os/exec"
)

func installMlt() {
	// Método 1: Descargar y extraer .tar.gz
	mlt_tar_url := "https://github.com/mltframework/mlt/releases/download/v7.26.0/mlt-7.26.0.tar.gz"
	mlt_cmd_tar := exec.Command("curl", "-L", mlt_tar_url, "-o", "package.tar.gz")
	err := mlt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mlt_zip_url := "https://github.com/mltframework/mlt/releases/download/v7.26.0/mlt-7.26.0.zip"
	mlt_cmd_zip := exec.Command("curl", "-L", mlt_zip_url, "-o", "package.zip")
	err = mlt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mlt_bin_url := "https://github.com/mltframework/mlt/releases/download/v7.26.0/mlt-7.26.0.bin"
	mlt_cmd_bin := exec.Command("curl", "-L", mlt_bin_url, "-o", "binary.bin")
	err = mlt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mlt_src_url := "https://github.com/mltframework/mlt/releases/download/v7.26.0/mlt-7.26.0.src.tar.gz"
	mlt_cmd_src := exec.Command("curl", "-L", mlt_src_url, "-o", "source.tar.gz")
	err = mlt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mlt_cmd_direct := exec.Command("./binary")
	err = mlt_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: frei0r")
exec.Command("latte", "install", "frei0r")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libdv")
exec.Command("latte", "install", "libdv")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libvidstab")
exec.Command("latte", "install", "libvidstab")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: opencv")
exec.Command("latte", "install", "opencv")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: rubberband")
exec.Command("latte", "install", "rubberband")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sox")
exec.Command("latte", "install", "sox")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
}
