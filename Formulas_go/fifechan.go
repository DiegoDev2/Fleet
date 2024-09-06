package main

// Fifechan - C++ GUI library designed for games
// Homepage: https://fifengine.github.io/fifechan/

import (
	"fmt"
	
	"os/exec"
)

func installFifechan() {
	// Método 1: Descargar y extraer .tar.gz
	fifechan_tar_url := "https://github.com/fifengine/fifechan/archive/refs/tags/0.1.5.tar.gz"
	fifechan_cmd_tar := exec.Command("curl", "-L", fifechan_tar_url, "-o", "package.tar.gz")
	err := fifechan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fifechan_zip_url := "https://github.com/fifengine/fifechan/archive/refs/tags/0.1.5.zip"
	fifechan_cmd_zip := exec.Command("curl", "-L", fifechan_zip_url, "-o", "package.zip")
	err = fifechan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fifechan_bin_url := "https://github.com/fifengine/fifechan/archive/refs/tags/0.1.5.bin"
	fifechan_cmd_bin := exec.Command("curl", "-L", fifechan_bin_url, "-o", "binary.bin")
	err = fifechan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fifechan_src_url := "https://github.com/fifengine/fifechan/archive/refs/tags/0.1.5.src.tar.gz"
	fifechan_cmd_src := exec.Command("curl", "-L", fifechan_src_url, "-o", "source.tar.gz")
	err = fifechan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fifechan_cmd_direct := exec.Command("./binary")
	err = fifechan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: allegro")
exec.Command("latte", "install", "allegro")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_image")
exec.Command("latte", "install", "sdl2_image")
	fmt.Println("Instalando dependencia: sdl2_ttf")
exec.Command("latte", "install", "sdl2_ttf")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
