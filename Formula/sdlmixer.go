package main

// SdlMixer - Sample multi-channel audio mixer library
// Homepage: https://www.libsdl.org/projects/SDL_mixer/release-1.2.html

import (
	"fmt"
	
	"os/exec"
)

func installSdlMixer() {
	// Método 1: Descargar y extraer .tar.gz
	sdlmixer_tar_url := "https://www.libsdl.org/projects/SDL_mixer/release/SDL_mixer-1.2.12.tar.gz"
	sdlmixer_cmd_tar := exec.Command("curl", "-L", sdlmixer_tar_url, "-o", "package.tar.gz")
	err := sdlmixer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdlmixer_zip_url := "https://www.libsdl.org/projects/SDL_mixer/release/SDL_mixer-1.2.12.zip"
	sdlmixer_cmd_zip := exec.Command("curl", "-L", sdlmixer_zip_url, "-o", "package.zip")
	err = sdlmixer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdlmixer_bin_url := "https://www.libsdl.org/projects/SDL_mixer/release/SDL_mixer-1.2.12.bin"
	sdlmixer_cmd_bin := exec.Command("curl", "-L", sdlmixer_bin_url, "-o", "binary.bin")
	err = sdlmixer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdlmixer_src_url := "https://www.libsdl.org/projects/SDL_mixer/release/SDL_mixer-1.2.12.src.tar.gz"
	sdlmixer_cmd_src := exec.Command("curl", "-L", sdlmixer_src_url, "-o", "source.tar.gz")
	err = sdlmixer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdlmixer_cmd_direct := exec.Command("./binary")
	err = sdlmixer_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libmikmod")
	exec.Command("latte", "install", "libmikmod").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
