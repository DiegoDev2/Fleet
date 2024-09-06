package main

// Qmmp - Qt-based Multimedia Player
// Homepage: https://qmmp.ylsoftware.com/

import (
	"fmt"
	
	"os/exec"
)

func installQmmp() {
	// Método 1: Descargar y extraer .tar.gz
	qmmp_tar_url := "https://qmmp.ylsoftware.com/files/qmmp/2.1/qmmp-2.1.9.tar.bz2"
	qmmp_cmd_tar := exec.Command("curl", "-L", qmmp_tar_url, "-o", "package.tar.gz")
	err := qmmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qmmp_zip_url := "https://qmmp.ylsoftware.com/files/qmmp/2.1/qmmp-2.1.9.tar.bz2"
	qmmp_cmd_zip := exec.Command("curl", "-L", qmmp_zip_url, "-o", "package.zip")
	err = qmmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qmmp_bin_url := "https://qmmp.ylsoftware.com/files/qmmp/2.1/qmmp-2.1.9.tar.bz2"
	qmmp_cmd_bin := exec.Command("curl", "-L", qmmp_bin_url, "-o", "binary.bin")
	err = qmmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qmmp_src_url := "https://qmmp.ylsoftware.com/files/qmmp/2.1/qmmp-2.1.9.tar.bz2"
	qmmp_cmd_src := exec.Command("curl", "-L", qmmp_src_url, "-o", "source.tar.gz")
	err = qmmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qmmp_cmd_direct := exec.Command("./binary")
	err = qmmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: faad2")
	exec.Command("latte", "install", "faad2").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: game-music-emu")
	exec.Command("latte", "install", "game-music-emu").Run()
	fmt.Println("Instalando dependencia: jack")
	exec.Command("latte", "install", "jack").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: libbs2b")
	exec.Command("latte", "install", "libbs2b").Run()
	fmt.Println("Instalando dependencia: libcddb")
	exec.Command("latte", "install", "libcddb").Run()
	fmt.Println("Instalando dependencia: libcdio")
	exec.Command("latte", "install", "libcdio").Run()
	fmt.Println("Instalando dependencia: libmms")
	exec.Command("latte", "install", "libmms").Run()
	fmt.Println("Instalando dependencia: libmodplug")
	exec.Command("latte", "install", "libmodplug").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: libshout")
	exec.Command("latte", "install", "libshout").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libsoxr")
	exec.Command("latte", "install", "libsoxr").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: libxmp")
	exec.Command("latte", "install", "libxmp").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
	fmt.Println("Instalando dependencia: mpg123")
	exec.Command("latte", "install", "mpg123").Run()
	fmt.Println("Instalando dependencia: mplayer")
	exec.Command("latte", "install", "mplayer").Run()
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
	fmt.Println("Instalando dependencia: opusfile")
	exec.Command("latte", "install", "opusfile").Run()
	fmt.Println("Instalando dependencia: projectm")
	exec.Command("latte", "install", "projectm").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: taglib")
	exec.Command("latte", "install", "taglib").Run()
	fmt.Println("Instalando dependencia: wavpack")
	exec.Command("latte", "install", "wavpack").Run()
	fmt.Println("Instalando dependencia: wildmidi")
	exec.Command("latte", "install", "wildmidi").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: musepack")
	exec.Command("latte", "install", "musepack").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
