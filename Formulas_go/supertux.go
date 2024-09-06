package main

// Supertux - Classic 2D jump'n run sidescroller game
// Homepage: https://www.supertux.org/

import (
	"fmt"
	
	"os/exec"
)

func installSupertux() {
	// Método 1: Descargar y extraer .tar.gz
	supertux_tar_url := "https://github.com/SuperTux/supertux/releases/download/v0.6.3/SuperTux-v0.6.3-Source.tar.gz"
	supertux_cmd_tar := exec.Command("curl", "-L", supertux_tar_url, "-o", "package.tar.gz")
	err := supertux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	supertux_zip_url := "https://github.com/SuperTux/supertux/releases/download/v0.6.3/SuperTux-v0.6.3-Source.zip"
	supertux_cmd_zip := exec.Command("curl", "-L", supertux_zip_url, "-o", "package.zip")
	err = supertux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	supertux_bin_url := "https://github.com/SuperTux/supertux/releases/download/v0.6.3/SuperTux-v0.6.3-Source.bin"
	supertux_cmd_bin := exec.Command("curl", "-L", supertux_bin_url, "-o", "binary.bin")
	err = supertux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	supertux_src_url := "https://github.com/SuperTux/supertux/releases/download/v0.6.3/SuperTux-v0.6.3-Source.src.tar.gz"
	supertux_cmd_src := exec.Command("curl", "-L", supertux_src_url, "-o", "source.tar.gz")
	err = supertux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	supertux_cmd_direct := exec.Command("./binary")
	err = supertux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: glm")
exec.Command("latte", "install", "glm")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: physfs")
exec.Command("latte", "install", "physfs")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: openal-soft")
exec.Command("latte", "install", "openal-soft")
}
