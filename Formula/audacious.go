package main

// Audacious - Lightweight and versatile audio player
// Homepage: https://audacious-media-player.org/

import (
	"fmt"
	
	"os/exec"
)

func installAudacious() {
	// Método 1: Descargar y extraer .tar.gz
	audacious_tar_url := "https://distfiles.audacious-media-player.org/audacious-4.4.tar.bz2"
	audacious_cmd_tar := exec.Command("curl", "-L", audacious_tar_url, "-o", "package.tar.gz")
	err := audacious_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	audacious_zip_url := "https://distfiles.audacious-media-player.org/audacious-4.4.tar.bz2"
	audacious_cmd_zip := exec.Command("curl", "-L", audacious_zip_url, "-o", "package.zip")
	err = audacious_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	audacious_bin_url := "https://distfiles.audacious-media-player.org/audacious-4.4.tar.bz2"
	audacious_cmd_bin := exec.Command("curl", "-L", audacious_bin_url, "-o", "binary.bin")
	err = audacious_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	audacious_src_url := "https://distfiles.audacious-media-player.org/audacious-4.4.tar.bz2"
	audacious_cmd_src := exec.Command("curl", "-L", audacious_src_url, "-o", "source.tar.gz")
	err = audacious_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	audacious_cmd_direct := exec.Command("./binary")
	err = audacious_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: faad2")
	exec.Command("latte", "install", "faad2").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: lame")
	exec.Command("latte", "install", "lame").Run()
	fmt.Println("Instalando dependencia: libbs2b")
	exec.Command("latte", "install", "libbs2b").Run()
	fmt.Println("Instalando dependencia: libcue")
	exec.Command("latte", "install", "libcue").Run()
	fmt.Println("Instalando dependencia: libmodplug")
	exec.Command("latte", "install", "libmodplug").Run()
	fmt.Println("Instalando dependencia: libnotify")
	exec.Command("latte", "install", "libnotify").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libopenmpt")
	exec.Command("latte", "install", "libopenmpt").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libsoxr")
	exec.Command("latte", "install", "libsoxr").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: mpg123")
	exec.Command("latte", "install", "mpg123").Run()
	fmt.Println("Instalando dependencia: neon")
	exec.Command("latte", "install", "neon").Run()
	fmt.Println("Instalando dependencia: opusfile")
	exec.Command("latte", "install", "opusfile").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: wavpack")
	exec.Command("latte", "install", "wavpack").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: jack")
	exec.Command("latte", "install", "jack").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
}
