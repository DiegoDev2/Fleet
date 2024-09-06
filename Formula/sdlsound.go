package main

// SdlSound - Library to decode several popular sound file formats
// Homepage: https://icculus.org/SDL_sound/

import (
	"fmt"
	
	"os/exec"
)

func installSdlSound() {
	// Método 1: Descargar y extraer .tar.gz
	sdlsound_tar_url := "https://icculus.org/SDL_sound/downloads/SDL_sound-1.0.3.tar.gz"
	sdlsound_cmd_tar := exec.Command("curl", "-L", sdlsound_tar_url, "-o", "package.tar.gz")
	err := sdlsound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdlsound_zip_url := "https://icculus.org/SDL_sound/downloads/SDL_sound-1.0.3.zip"
	sdlsound_cmd_zip := exec.Command("curl", "-L", sdlsound_zip_url, "-o", "package.zip")
	err = sdlsound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdlsound_bin_url := "https://icculus.org/SDL_sound/downloads/SDL_sound-1.0.3.bin"
	sdlsound_cmd_bin := exec.Command("curl", "-L", sdlsound_bin_url, "-o", "binary.bin")
	err = sdlsound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdlsound_src_url := "https://icculus.org/SDL_sound/downloads/SDL_sound-1.0.3.src.tar.gz"
	sdlsound_cmd_src := exec.Command("curl", "-L", sdlsound_src_url, "-o", "source.tar.gz")
	err = sdlsound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdlsound_cmd_direct := exec.Command("./binary")
	err = sdlsound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
