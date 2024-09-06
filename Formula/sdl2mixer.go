package main

// Sdl2Mixer - Sample multi-channel audio mixer library
// Homepage: https://github.com/libsdl-org/SDL_mixer

import (
	"fmt"
	
	"os/exec"
)

func installSdl2Mixer() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2mixer_tar_url := "https://github.com/libsdl-org/SDL_mixer/releases/download/release-2.8.0/SDL2_mixer-2.8.0.tar.gz"
	sdl2mixer_cmd_tar := exec.Command("curl", "-L", sdl2mixer_tar_url, "-o", "package.tar.gz")
	err := sdl2mixer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2mixer_zip_url := "https://github.com/libsdl-org/SDL_mixer/releases/download/release-2.8.0/SDL2_mixer-2.8.0.zip"
	sdl2mixer_cmd_zip := exec.Command("curl", "-L", sdl2mixer_zip_url, "-o", "package.zip")
	err = sdl2mixer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2mixer_bin_url := "https://github.com/libsdl-org/SDL_mixer/releases/download/release-2.8.0/SDL2_mixer-2.8.0.bin"
	sdl2mixer_cmd_bin := exec.Command("curl", "-L", sdl2mixer_bin_url, "-o", "binary.bin")
	err = sdl2mixer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2mixer_src_url := "https://github.com/libsdl-org/SDL_mixer/releases/download/release-2.8.0/SDL2_mixer-2.8.0.src.tar.gz"
	sdl2mixer_cmd_src := exec.Command("curl", "-L", sdl2mixer_src_url, "-o", "source.tar.gz")
	err = sdl2mixer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2mixer_cmd_direct := exec.Command("./binary")
	err = sdl2mixer_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: game-music-emu")
	exec.Command("latte", "install", "game-music-emu").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libxmp")
	exec.Command("latte", "install", "libxmp").Run()
	fmt.Println("Instalando dependencia: mpg123")
	exec.Command("latte", "install", "mpg123").Run()
	fmt.Println("Instalando dependencia: opusfile")
	exec.Command("latte", "install", "opusfile").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: wavpack")
	exec.Command("latte", "install", "wavpack").Run()
}
