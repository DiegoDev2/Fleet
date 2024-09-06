package main

// Sdl12Compat - SDL 1.2 compatibility layer that uses SDL 2.0 behind the scenes
// Homepage: https://github.com/libsdl-org/sdl12-compat

import (
	"fmt"
	
	"os/exec"
)

func installSdl12Compat() {
	// Método 1: Descargar y extraer .tar.gz
	sdl12compat_tar_url := "https://github.com/libsdl-org/sdl12-compat/archive/refs/tags/release-1.2.68.tar.gz"
	sdl12compat_cmd_tar := exec.Command("curl", "-L", sdl12compat_tar_url, "-o", "package.tar.gz")
	err := sdl12compat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl12compat_zip_url := "https://github.com/libsdl-org/sdl12-compat/archive/refs/tags/release-1.2.68.zip"
	sdl12compat_cmd_zip := exec.Command("curl", "-L", sdl12compat_zip_url, "-o", "package.zip")
	err = sdl12compat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl12compat_bin_url := "https://github.com/libsdl-org/sdl12-compat/archive/refs/tags/release-1.2.68.bin"
	sdl12compat_cmd_bin := exec.Command("curl", "-L", sdl12compat_bin_url, "-o", "binary.bin")
	err = sdl12compat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl12compat_src_url := "https://github.com/libsdl-org/sdl12-compat/archive/refs/tags/release-1.2.68.src.tar.gz"
	sdl12compat_cmd_src := exec.Command("curl", "-L", sdl12compat_src_url, "-o", "source.tar.gz")
	err = sdl12compat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl12compat_cmd_direct := exec.Command("./binary")
	err = sdl12compat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
