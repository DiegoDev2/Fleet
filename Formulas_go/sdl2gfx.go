package main

// Sdl2Gfx - SDL2 graphics drawing primitives and other support functions
// Homepage: https://www.ferzkopp.net/wordpress/2016/01/02/sdl_gfx-sdl2_gfx/

import (
	"fmt"
	
	"os/exec"
)

func installSdl2Gfx() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2gfx_tar_url := "https://www.ferzkopp.net/Software/SDL2_gfx/SDL2_gfx-1.0.4.tar.gz"
	sdl2gfx_cmd_tar := exec.Command("curl", "-L", sdl2gfx_tar_url, "-o", "package.tar.gz")
	err := sdl2gfx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2gfx_zip_url := "https://www.ferzkopp.net/Software/SDL2_gfx/SDL2_gfx-1.0.4.zip"
	sdl2gfx_cmd_zip := exec.Command("curl", "-L", sdl2gfx_zip_url, "-o", "package.zip")
	err = sdl2gfx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2gfx_bin_url := "https://www.ferzkopp.net/Software/SDL2_gfx/SDL2_gfx-1.0.4.bin"
	sdl2gfx_cmd_bin := exec.Command("curl", "-L", sdl2gfx_bin_url, "-o", "binary.bin")
	err = sdl2gfx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2gfx_src_url := "https://www.ferzkopp.net/Software/SDL2_gfx/SDL2_gfx-1.0.4.src.tar.gz"
	sdl2gfx_cmd_src := exec.Command("curl", "-L", sdl2gfx_src_url, "-o", "source.tar.gz")
	err = sdl2gfx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2gfx_cmd_direct := exec.Command("./binary")
	err = sdl2gfx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
