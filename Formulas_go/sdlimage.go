package main

// SdlImage - Image file loading library
// Homepage: https://github.com/libsdl-org/SDL_image

import (
	"fmt"
	
	"os/exec"
)

func installSdlImage() {
	// Método 1: Descargar y extraer .tar.gz
	sdlimage_tar_url := "https://www.libsdl.org/projects/SDL_image/release/SDL_image-1.2.12.tar.gz"
	sdlimage_cmd_tar := exec.Command("curl", "-L", sdlimage_tar_url, "-o", "package.tar.gz")
	err := sdlimage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdlimage_zip_url := "https://www.libsdl.org/projects/SDL_image/release/SDL_image-1.2.12.zip"
	sdlimage_cmd_zip := exec.Command("curl", "-L", sdlimage_zip_url, "-o", "package.zip")
	err = sdlimage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdlimage_bin_url := "https://www.libsdl.org/projects/SDL_image/release/SDL_image-1.2.12.bin"
	sdlimage_cmd_bin := exec.Command("curl", "-L", sdlimage_bin_url, "-o", "binary.bin")
	err = sdlimage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdlimage_src_url := "https://www.libsdl.org/projects/SDL_image/release/SDL_image-1.2.12.src.tar.gz"
	sdlimage_cmd_src := exec.Command("curl", "-L", sdlimage_src_url, "-o", "source.tar.gz")
	err = sdlimage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdlimage_cmd_direct := exec.Command("./binary")
	err = sdlimage_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
}
