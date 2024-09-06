package main

// Mpd - Music Player Daemon
// Homepage: https://github.com/MusicPlayerDaemon/MPD

import (
	"fmt"
	
	"os/exec"
)

func installMpd() {
	// Método 1: Descargar y extraer .tar.gz
	mpd_tar_url := "https://github.com/MusicPlayerDaemon/MPD/archive/refs/tags/v0.23.15.tar.gz"
	mpd_cmd_tar := exec.Command("curl", "-L", mpd_tar_url, "-o", "package.tar.gz")
	err := mpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpd_zip_url := "https://github.com/MusicPlayerDaemon/MPD/archive/refs/tags/v0.23.15.zip"
	mpd_cmd_zip := exec.Command("curl", "-L", mpd_zip_url, "-o", "package.zip")
	err = mpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpd_bin_url := "https://github.com/MusicPlayerDaemon/MPD/archive/refs/tags/v0.23.15.bin"
	mpd_cmd_bin := exec.Command("curl", "-L", mpd_bin_url, "-o", "binary.bin")
	err = mpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpd_src_url := "https://github.com/MusicPlayerDaemon/MPD/archive/refs/tags/v0.23.15.src.tar.gz"
	mpd_cmd_src := exec.Command("curl", "-L", mpd_src_url, "-o", "source.tar.gz")
	err = mpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpd_cmd_direct := exec.Command("./binary")
	err = mpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: chromaprint")
exec.Command("latte", "install", "chromaprint")
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
	fmt.Println("Instalando dependencia: faad2")
exec.Command("latte", "install", "faad2")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: fluid-synth")
exec.Command("latte", "install", "fluid-synth")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libid3tag")
exec.Command("latte", "install", "libid3tag")
	fmt.Println("Instalando dependencia: libmpdclient")
exec.Command("latte", "install", "libmpdclient")
	fmt.Println("Instalando dependencia: libnfs")
exec.Command("latte", "install", "libnfs")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libshout")
exec.Command("latte", "install", "libshout")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libsoxr")
exec.Command("latte", "install", "libsoxr")
	fmt.Println("Instalando dependencia: libupnp")
exec.Command("latte", "install", "libupnp")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: mpg123")
exec.Command("latte", "install", "mpg123")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: wavpack")
exec.Command("latte", "install", "wavpack")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
