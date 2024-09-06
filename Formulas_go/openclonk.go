package main

// Openclonk - Multiplayer action game
// Homepage: https://www.openclonk.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenclonk() {
	// Método 1: Descargar y extraer .tar.gz
	openclonk_tar_url := "https://www.openclonk.org/builds/release/7.0/openclonk-7.0-src.tar.bz2"
	openclonk_cmd_tar := exec.Command("curl", "-L", openclonk_tar_url, "-o", "package.tar.gz")
	err := openclonk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openclonk_zip_url := "https://www.openclonk.org/builds/release/7.0/openclonk-7.0-src.tar.bz2"
	openclonk_cmd_zip := exec.Command("curl", "-L", openclonk_zip_url, "-o", "package.zip")
	err = openclonk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openclonk_bin_url := "https://www.openclonk.org/builds/release/7.0/openclonk-7.0-src.tar.bz2"
	openclonk_cmd_bin := exec.Command("curl", "-L", openclonk_bin_url, "-o", "binary.bin")
	err = openclonk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openclonk_src_url := "https://www.openclonk.org/builds/release/7.0/openclonk-7.0-src.tar.bz2"
	openclonk_cmd_src := exec.Command("curl", "-L", openclonk_src_url, "-o", "source.tar.gz")
	err = openclonk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openclonk_cmd_direct := exec.Command("./binary")
	err = openclonk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libepoxy")
exec.Command("latte", "install", "libepoxy")
	fmt.Println("Instalando dependencia: miniupnpc")
exec.Command("latte", "install", "miniupnpc")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: freealut")
exec.Command("latte", "install", "freealut")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: tinyxml")
exec.Command("latte", "install", "tinyxml")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: openal-soft")
exec.Command("latte", "install", "openal-soft")
}
