package main

// Musikcube - Terminal-based audio engine, library, player and server
// Homepage: https://musikcube.com

import (
	"fmt"
	
	"os/exec"
)

func installMusikcube() {
	// Método 1: Descargar y extraer .tar.gz
	musikcube_tar_url := "https://github.com/clangen/musikcube/archive/refs/tags/3.0.4.tar.gz"
	musikcube_cmd_tar := exec.Command("curl", "-L", musikcube_tar_url, "-o", "package.tar.gz")
	err := musikcube_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	musikcube_zip_url := "https://github.com/clangen/musikcube/archive/refs/tags/3.0.4.zip"
	musikcube_cmd_zip := exec.Command("curl", "-L", musikcube_zip_url, "-o", "package.zip")
	err = musikcube_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	musikcube_bin_url := "https://github.com/clangen/musikcube/archive/refs/tags/3.0.4.bin"
	musikcube_cmd_bin := exec.Command("curl", "-L", musikcube_bin_url, "-o", "binary.bin")
	err = musikcube_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	musikcube_src_url := "https://github.com/clangen/musikcube/archive/refs/tags/3.0.4.src.tar.gz"
	musikcube_cmd_src := exec.Command("curl", "-L", musikcube_src_url, "-o", "source.tar.gz")
	err = musikcube_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	musikcube_cmd_direct := exec.Command("./binary")
	err = musikcube_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asio")
exec.Command("latte", "install", "asio")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: game-music-emu")
exec.Command("latte", "install", "game-music-emu")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: libmicrohttpd")
exec.Command("latte", "install", "libmicrohttpd")
	fmt.Println("Instalando dependencia: libopenmpt")
exec.Command("latte", "install", "libopenmpt")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: mpg123")
exec.Command("latte", "install", "mpg123")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
