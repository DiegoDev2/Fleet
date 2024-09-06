package main

// SpotifyTui - Terminal-based client for Spotify
// Homepage: https://github.com/Rigellute/spotify-tui

import (
	"fmt"
	
	"os/exec"
)

func installSpotifyTui() {
	// Método 1: Descargar y extraer .tar.gz
	spotifytui_tar_url := "https://github.com/Rigellute/spotify-tui/archive/refs/tags/v0.25.0.tar.gz"
	spotifytui_cmd_tar := exec.Command("curl", "-L", spotifytui_tar_url, "-o", "package.tar.gz")
	err := spotifytui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spotifytui_zip_url := "https://github.com/Rigellute/spotify-tui/archive/refs/tags/v0.25.0.zip"
	spotifytui_cmd_zip := exec.Command("curl", "-L", spotifytui_zip_url, "-o", "package.zip")
	err = spotifytui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spotifytui_bin_url := "https://github.com/Rigellute/spotify-tui/archive/refs/tags/v0.25.0.bin"
	spotifytui_cmd_bin := exec.Command("curl", "-L", spotifytui_bin_url, "-o", "binary.bin")
	err = spotifytui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spotifytui_src_url := "https://github.com/Rigellute/spotify-tui/archive/refs/tags/v0.25.0.src.tar.gz"
	spotifytui_cmd_src := exec.Command("curl", "-L", spotifytui_src_url, "-o", "source.tar.gz")
	err = spotifytui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spotifytui_cmd_direct := exec.Command("./binary")
	err = spotifytui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: openssl@1.1")
exec.Command("latte", "install", "openssl@1.1")
}
