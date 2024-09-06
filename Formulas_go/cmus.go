package main

// Cmus - Music player with an ncurses based interface
// Homepage: https://cmus.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCmus() {
	// Método 1: Descargar y extraer .tar.gz
	cmus_tar_url := "https://github.com/cmus/cmus/archive/refs/tags/v2.11.0.tar.gz"
	cmus_cmd_tar := exec.Command("curl", "-L", cmus_tar_url, "-o", "package.tar.gz")
	err := cmus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmus_zip_url := "https://github.com/cmus/cmus/archive/refs/tags/v2.11.0.zip"
	cmus_cmd_zip := exec.Command("curl", "-L", cmus_zip_url, "-o", "package.zip")
	err = cmus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmus_bin_url := "https://github.com/cmus/cmus/archive/refs/tags/v2.11.0.bin"
	cmus_cmd_bin := exec.Command("curl", "-L", cmus_bin_url, "-o", "binary.bin")
	err = cmus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmus_src_url := "https://github.com/cmus/cmus/archive/refs/tags/v2.11.0.src.tar.gz"
	cmus_cmd_src := exec.Command("curl", "-L", cmus_src_url, "-o", "source.tar.gz")
	err = cmus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmus_cmd_direct := exec.Command("./binary")
	err = cmus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: faad2")
exec.Command("latte", "install", "faad2")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
	fmt.Println("Instalando dependencia: libcue")
exec.Command("latte", "install", "libcue")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: mad")
exec.Command("latte", "install", "mad")
	fmt.Println("Instalando dependencia: mp4v2")
exec.Command("latte", "install", "mp4v2")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: opusfile")
exec.Command("latte", "install", "opusfile")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
}
