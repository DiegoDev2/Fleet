package main

// SdlNet - Sample cross-platform networking library
// Homepage: https://www.libsdl.org/projects/SDL_net/release-1.2.html

import (
	"fmt"
	
	"os/exec"
)

func installSdlNet() {
	// Método 1: Descargar y extraer .tar.gz
	sdlnet_tar_url := "https://www.libsdl.org/projects/SDL_net/release/SDL_net-1.2.8.tar.gz"
	sdlnet_cmd_tar := exec.Command("curl", "-L", sdlnet_tar_url, "-o", "package.tar.gz")
	err := sdlnet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdlnet_zip_url := "https://www.libsdl.org/projects/SDL_net/release/SDL_net-1.2.8.zip"
	sdlnet_cmd_zip := exec.Command("curl", "-L", sdlnet_zip_url, "-o", "package.zip")
	err = sdlnet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdlnet_bin_url := "https://www.libsdl.org/projects/SDL_net/release/SDL_net-1.2.8.bin"
	sdlnet_cmd_bin := exec.Command("curl", "-L", sdlnet_bin_url, "-o", "binary.bin")
	err = sdlnet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdlnet_src_url := "https://www.libsdl.org/projects/SDL_net/release/SDL_net-1.2.8.src.tar.gz"
	sdlnet_cmd_src := exec.Command("curl", "-L", sdlnet_src_url, "-o", "source.tar.gz")
	err = sdlnet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdlnet_cmd_direct := exec.Command("./binary")
	err = sdlnet_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
