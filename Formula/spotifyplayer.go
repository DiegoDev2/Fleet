package main

// SpotifyPlayer - Command driven spotify player
// Homepage: https://github.com/aome510/spotify-player

import (
	"fmt"
	
	"os/exec"
)

func installSpotifyPlayer() {
	// Método 1: Descargar y extraer .tar.gz
	spotifyplayer_tar_url := "https://github.com/aome510/spotify-player/archive/refs/tags/v0.19.1.tar.gz"
	spotifyplayer_cmd_tar := exec.Command("curl", "-L", spotifyplayer_tar_url, "-o", "package.tar.gz")
	err := spotifyplayer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spotifyplayer_zip_url := "https://github.com/aome510/spotify-player/archive/refs/tags/v0.19.1.zip"
	spotifyplayer_cmd_zip := exec.Command("curl", "-L", spotifyplayer_zip_url, "-o", "package.zip")
	err = spotifyplayer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spotifyplayer_bin_url := "https://github.com/aome510/spotify-player/archive/refs/tags/v0.19.1.bin"
	spotifyplayer_cmd_bin := exec.Command("curl", "-L", spotifyplayer_bin_url, "-o", "binary.bin")
	err = spotifyplayer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spotifyplayer_src_url := "https://github.com/aome510/spotify-player/archive/refs/tags/v0.19.1.src.tar.gz"
	spotifyplayer_cmd_src := exec.Command("curl", "-L", spotifyplayer_src_url, "-o", "source.tar.gz")
	err = spotifyplayer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spotifyplayer_cmd_direct := exec.Command("./binary")
	err = spotifyplayer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: dbus")
	exec.Command("latte", "install", "dbus").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
