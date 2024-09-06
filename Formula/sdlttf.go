package main

// SdlTtf - Library for using TrueType fonts in SDL applications
// Homepage: https://github.com/libsdl-org/SDL_ttf

import (
	"fmt"
	
	"os/exec"
)

func installSdlTtf() {
	// Método 1: Descargar y extraer .tar.gz
	sdlttf_tar_url := "https://www.libsdl.org/projects/SDL_ttf/release/SDL_ttf-2.0.11.tar.gz"
	sdlttf_cmd_tar := exec.Command("curl", "-L", sdlttf_tar_url, "-o", "package.tar.gz")
	err := sdlttf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdlttf_zip_url := "https://www.libsdl.org/projects/SDL_ttf/release/SDL_ttf-2.0.11.zip"
	sdlttf_cmd_zip := exec.Command("curl", "-L", sdlttf_zip_url, "-o", "package.zip")
	err = sdlttf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdlttf_bin_url := "https://www.libsdl.org/projects/SDL_ttf/release/SDL_ttf-2.0.11.bin"
	sdlttf_cmd_bin := exec.Command("curl", "-L", sdlttf_bin_url, "-o", "binary.bin")
	err = sdlttf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdlttf_src_url := "https://www.libsdl.org/projects/SDL_ttf/release/SDL_ttf-2.0.11.src.tar.gz"
	sdlttf_cmd_src := exec.Command("curl", "-L", sdlttf_src_url, "-o", "source.tar.gz")
	err = sdlttf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdlttf_cmd_direct := exec.Command("./binary")
	err = sdlttf_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
