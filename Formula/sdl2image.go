package main

// Sdl2Image - Library for loading images as SDL surfaces and textures
// Homepage: https://github.com/libsdl-org/SDL_image

import (
	"fmt"
	
	"os/exec"
)

func installSdl2Image() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2image_tar_url := "https://github.com/libsdl-org/SDL_image/releases/download/release-2.8.2/SDL2_image-2.8.2.tar.gz"
	sdl2image_cmd_tar := exec.Command("curl", "-L", sdl2image_tar_url, "-o", "package.tar.gz")
	err := sdl2image_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2image_zip_url := "https://github.com/libsdl-org/SDL_image/releases/download/release-2.8.2/SDL2_image-2.8.2.zip"
	sdl2image_cmd_zip := exec.Command("curl", "-L", sdl2image_zip_url, "-o", "package.zip")
	err = sdl2image_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2image_bin_url := "https://github.com/libsdl-org/SDL_image/releases/download/release-2.8.2/SDL2_image-2.8.2.bin"
	sdl2image_cmd_bin := exec.Command("curl", "-L", sdl2image_bin_url, "-o", "binary.bin")
	err = sdl2image_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2image_src_url := "https://github.com/libsdl-org/SDL_image/releases/download/release-2.8.2/SDL2_image-2.8.2.src.tar.gz"
	sdl2image_cmd_src := exec.Command("curl", "-L", sdl2image_src_url, "-o", "source.tar.gz")
	err = sdl2image_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2image_cmd_direct := exec.Command("./binary")
	err = sdl2image_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: jpeg-xl")
	exec.Command("latte", "install", "jpeg-xl").Run()
	fmt.Println("Instalando dependencia: libavif")
	exec.Command("latte", "install", "libavif").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: webp")
	exec.Command("latte", "install", "webp").Run()
}
