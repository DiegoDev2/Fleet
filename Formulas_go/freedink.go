package main

// Freedink - Portable version of the Dink Smallwood game engine
// Homepage: https://www.gnu.org/software/freedink/

import (
	"fmt"
	
	"os/exec"
)

func installFreedink() {
	// Método 1: Descargar y extraer .tar.gz
	freedink_tar_url := "https://ftp.gnu.org/gnu/freedink/freedink-109.6.tar.gz"
	freedink_cmd_tar := exec.Command("curl", "-L", freedink_tar_url, "-o", "package.tar.gz")
	err := freedink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freedink_zip_url := "https://ftp.gnu.org/gnu/freedink/freedink-109.6.zip"
	freedink_cmd_zip := exec.Command("curl", "-L", freedink_zip_url, "-o", "package.zip")
	err = freedink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freedink_bin_url := "https://ftp.gnu.org/gnu/freedink/freedink-109.6.bin"
	freedink_cmd_bin := exec.Command("curl", "-L", freedink_bin_url, "-o", "binary.bin")
	err = freedink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freedink_src_url := "https://ftp.gnu.org/gnu/freedink/freedink-109.6.src.tar.gz"
	freedink_cmd_src := exec.Command("curl", "-L", freedink_src_url, "-o", "source.tar.gz")
	err = freedink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freedink_cmd_direct := exec.Command("./binary")
	err = freedink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glm")
exec.Command("latte", "install", "glm")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: check")
exec.Command("latte", "install", "check")
	fmt.Println("Instalando dependencia: cxxtest")
exec.Command("latte", "install", "cxxtest")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_gfx")
exec.Command("latte", "install", "sdl2_gfx")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
	fmt.Println("Instalando dependencia: sdl2_ttf")
exec.Command("latte", "install", "sdl2_ttf")
}
