package main

// Enigma - Puzzle game inspired by Oxyd and Rock'n'Roll
// Homepage: https://www.nongnu.org/enigma/

import (
	"fmt"
	
	"os/exec"
)

func installEnigma() {
	// Método 1: Descargar y extraer .tar.gz
	enigma_tar_url := "https://github.com/Enigma-Game/Enigma/releases/download/1.30/Enigma-1.30-src.tar.gz"
	enigma_cmd_tar := exec.Command("curl", "-L", enigma_tar_url, "-o", "package.tar.gz")
	err := enigma_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enigma_zip_url := "https://github.com/Enigma-Game/Enigma/releases/download/1.30/Enigma-1.30-src.zip"
	enigma_cmd_zip := exec.Command("curl", "-L", enigma_zip_url, "-o", "package.zip")
	err = enigma_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enigma_bin_url := "https://github.com/Enigma-Game/Enigma/releases/download/1.30/Enigma-1.30-src.bin"
	enigma_cmd_bin := exec.Command("curl", "-L", enigma_bin_url, "-o", "binary.bin")
	err = enigma_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enigma_src_url := "https://github.com/Enigma-Game/Enigma/releases/download/1.30/Enigma-1.30-src.src.tar.gz"
	enigma_cmd_src := exec.Command("curl", "-L", enigma_src_url, "-o", "source.tar.gz")
	err = enigma_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enigma_cmd_direct := exec.Command("./binary")
	err = enigma_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: texi2html")
exec.Command("latte", "install", "texi2html")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: enet")
exec.Command("latte", "install", "enet")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
	fmt.Println("Instalando dependencia: sdl2_ttf")
exec.Command("latte", "install", "sdl2_ttf")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
