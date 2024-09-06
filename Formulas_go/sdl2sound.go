package main

// Sdl2Sound - Abstract soundfile decoder for SDL
// Homepage: https://icculus.org/SDL_sound/

import (
	"fmt"
	
	"os/exec"
)

func installSdl2Sound() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2sound_tar_url := "https://github.com/icculus/SDL_sound/releases/download/v2.0.2/SDL2_sound-2.0.2.tar.gz"
	sdl2sound_cmd_tar := exec.Command("curl", "-L", sdl2sound_tar_url, "-o", "package.tar.gz")
	err := sdl2sound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2sound_zip_url := "https://github.com/icculus/SDL_sound/releases/download/v2.0.2/SDL2_sound-2.0.2.zip"
	sdl2sound_cmd_zip := exec.Command("curl", "-L", sdl2sound_zip_url, "-o", "package.zip")
	err = sdl2sound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2sound_bin_url := "https://github.com/icculus/SDL_sound/releases/download/v2.0.2/SDL2_sound-2.0.2.bin"
	sdl2sound_cmd_bin := exec.Command("curl", "-L", sdl2sound_bin_url, "-o", "binary.bin")
	err = sdl2sound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2sound_src_url := "https://github.com/icculus/SDL_sound/releases/download/v2.0.2/SDL2_sound-2.0.2.src.tar.gz"
	sdl2sound_cmd_src := exec.Command("curl", "-L", sdl2sound_src_url, "-o", "source.tar.gz")
	err = sdl2sound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2sound_cmd_direct := exec.Command("./binary")
	err = sdl2sound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
