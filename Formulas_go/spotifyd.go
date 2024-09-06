package main

// Spotifyd - Spotify daemon
// Homepage: https://github.com/Spotifyd/spotifyd

import (
	"fmt"
	
	"os/exec"
)

func installSpotifyd() {
	// Método 1: Descargar y extraer .tar.gz
	spotifyd_tar_url := "https://github.com/Spotifyd/spotifyd/archive/refs/tags/v0.3.5.tar.gz"
	spotifyd_cmd_tar := exec.Command("curl", "-L", spotifyd_tar_url, "-o", "package.tar.gz")
	err := spotifyd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spotifyd_zip_url := "https://github.com/Spotifyd/spotifyd/archive/refs/tags/v0.3.5.zip"
	spotifyd_cmd_zip := exec.Command("curl", "-L", spotifyd_zip_url, "-o", "package.zip")
	err = spotifyd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spotifyd_bin_url := "https://github.com/Spotifyd/spotifyd/archive/refs/tags/v0.3.5.bin"
	spotifyd_cmd_bin := exec.Command("curl", "-L", spotifyd_bin_url, "-o", "binary.bin")
	err = spotifyd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spotifyd_src_url := "https://github.com/Spotifyd/spotifyd/archive/refs/tags/v0.3.5.src.tar.gz"
	spotifyd_cmd_src := exec.Command("curl", "-L", spotifyd_src_url, "-o", "source.tar.gz")
	err = spotifyd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spotifyd_cmd_direct := exec.Command("./binary")
	err = spotifyd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: dbus")
exec.Command("latte", "install", "dbus")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
}
