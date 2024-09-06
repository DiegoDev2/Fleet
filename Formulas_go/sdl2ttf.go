package main

// Sdl2Ttf - Library for using TrueType fonts in SDL applications
// Homepage: https://github.com/libsdl-org/SDL_ttf

import (
	"fmt"
	
	"os/exec"
)

func installSdl2Ttf() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2ttf_tar_url := "https://github.com/libsdl-org/SDL_ttf/releases/download/release-2.22.0/SDL2_ttf-2.22.0.tar.gz"
	sdl2ttf_cmd_tar := exec.Command("curl", "-L", sdl2ttf_tar_url, "-o", "package.tar.gz")
	err := sdl2ttf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2ttf_zip_url := "https://github.com/libsdl-org/SDL_ttf/releases/download/release-2.22.0/SDL2_ttf-2.22.0.zip"
	sdl2ttf_cmd_zip := exec.Command("curl", "-L", sdl2ttf_zip_url, "-o", "package.zip")
	err = sdl2ttf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2ttf_bin_url := "https://github.com/libsdl-org/SDL_ttf/releases/download/release-2.22.0/SDL2_ttf-2.22.0.bin"
	sdl2ttf_cmd_bin := exec.Command("curl", "-L", sdl2ttf_bin_url, "-o", "binary.bin")
	err = sdl2ttf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2ttf_src_url := "https://github.com/libsdl-org/SDL_ttf/releases/download/release-2.22.0/SDL2_ttf-2.22.0.src.tar.gz"
	sdl2ttf_cmd_src := exec.Command("curl", "-L", sdl2ttf_src_url, "-o", "source.tar.gz")
	err = sdl2ttf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2ttf_cmd_direct := exec.Command("./binary")
	err = sdl2ttf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
