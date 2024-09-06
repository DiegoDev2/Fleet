package main

// Sdlpop - Open-source port of Prince of Persia
// Homepage: https://github.com/NagyD/SDLPoP

import (
	"fmt"
	
	"os/exec"
)

func installSdlpop() {
	// Método 1: Descargar y extraer .tar.gz
	sdlpop_tar_url := "https://github.com/NagyD/SDLPoP/archive/refs/tags/v1.23.tar.gz"
	sdlpop_cmd_tar := exec.Command("curl", "-L", sdlpop_tar_url, "-o", "package.tar.gz")
	err := sdlpop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdlpop_zip_url := "https://github.com/NagyD/SDLPoP/archive/refs/tags/v1.23.zip"
	sdlpop_cmd_zip := exec.Command("curl", "-L", sdlpop_zip_url, "-o", "package.zip")
	err = sdlpop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdlpop_bin_url := "https://github.com/NagyD/SDLPoP/archive/refs/tags/v1.23.bin"
	sdlpop_cmd_bin := exec.Command("curl", "-L", sdlpop_bin_url, "-o", "binary.bin")
	err = sdlpop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdlpop_src_url := "https://github.com/NagyD/SDLPoP/archive/refs/tags/v1.23.src.tar.gz"
	sdlpop_cmd_src := exec.Command("curl", "-L", sdlpop_src_url, "-o", "source.tar.gz")
	err = sdlpop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdlpop_cmd_direct := exec.Command("./binary")
	err = sdlpop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
}
