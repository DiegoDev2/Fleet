package main

// Corsixth - Open source clone of Theme Hospital
// Homepage: https://github.com/CorsixTH/CorsixTH

import (
	"fmt"
	
	"os/exec"
)

func installCorsixth() {
	// Método 1: Descargar y extraer .tar.gz
	corsixth_tar_url := "https://github.com/CorsixTH/CorsixTH/archive/refs/tags/v0.67.tar.gz"
	corsixth_cmd_tar := exec.Command("curl", "-L", corsixth_tar_url, "-o", "package.tar.gz")
	err := corsixth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	corsixth_zip_url := "https://github.com/CorsixTH/CorsixTH/archive/refs/tags/v0.67.zip"
	corsixth_cmd_zip := exec.Command("curl", "-L", corsixth_zip_url, "-o", "package.zip")
	err = corsixth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	corsixth_bin_url := "https://github.com/CorsixTH/CorsixTH/archive/refs/tags/v0.67.bin"
	corsixth_cmd_bin := exec.Command("curl", "-L", corsixth_bin_url, "-o", "binary.bin")
	err = corsixth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	corsixth_src_url := "https://github.com/CorsixTH/CorsixTH/archive/refs/tags/v0.67.src.tar.gz"
	corsixth_cmd_src := exec.Command("curl", "-L", corsixth_src_url, "-o", "source.tar.gz")
	err = corsixth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	corsixth_cmd_direct := exec.Command("./binary")
	err = corsixth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: luarocks")
exec.Command("latte", "install", "luarocks")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: lpeg")
exec.Command("latte", "install", "lpeg")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
