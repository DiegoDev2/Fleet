package main

// Sdl2Net - Small sample cross-platform networking library
// Homepage: https://github.com/libsdl-org/SDL_net

import (
	"fmt"
	
	"os/exec"
)

func installSdl2Net() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2net_tar_url := "https://github.com/libsdl-org/SDL_net/releases/download/release-2.2.0/SDL2_net-2.2.0.tar.gz"
	sdl2net_cmd_tar := exec.Command("curl", "-L", sdl2net_tar_url, "-o", "package.tar.gz")
	err := sdl2net_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2net_zip_url := "https://github.com/libsdl-org/SDL_net/releases/download/release-2.2.0/SDL2_net-2.2.0.zip"
	sdl2net_cmd_zip := exec.Command("curl", "-L", sdl2net_zip_url, "-o", "package.zip")
	err = sdl2net_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2net_bin_url := "https://github.com/libsdl-org/SDL_net/releases/download/release-2.2.0/SDL2_net-2.2.0.bin"
	sdl2net_cmd_bin := exec.Command("curl", "-L", sdl2net_bin_url, "-o", "binary.bin")
	err = sdl2net_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2net_src_url := "https://github.com/libsdl-org/SDL_net/releases/download/release-2.2.0/SDL2_net-2.2.0.src.tar.gz"
	sdl2net_cmd_src := exec.Command("curl", "-L", sdl2net_src_url, "-o", "source.tar.gz")
	err = sdl2net_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2net_cmd_direct := exec.Command("./binary")
	err = sdl2net_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
