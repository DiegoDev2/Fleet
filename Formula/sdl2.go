package main

// Sdl2 - Low-level access to audio, keyboard, mouse, joystick, and graphics
// Homepage: https://www.libsdl.org/

import (
	"fmt"
	
	"os/exec"
)

func installSdl2() {
	// Método 1: Descargar y extraer .tar.gz
	sdl2_tar_url := "https://github.com/libsdl-org/SDL/releases/download/release-2.30.7/SDL2-2.30.7.tar.gz"
	sdl2_cmd_tar := exec.Command("curl", "-L", sdl2_tar_url, "-o", "package.tar.gz")
	err := sdl2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdl2_zip_url := "https://github.com/libsdl-org/SDL/releases/download/release-2.30.7/SDL2-2.30.7.zip"
	sdl2_cmd_zip := exec.Command("curl", "-L", sdl2_zip_url, "-o", "package.zip")
	err = sdl2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdl2_bin_url := "https://github.com/libsdl-org/SDL/releases/download/release-2.30.7/SDL2-2.30.7.bin"
	sdl2_cmd_bin := exec.Command("curl", "-L", sdl2_bin_url, "-o", "binary.bin")
	err = sdl2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdl2_src_url := "https://github.com/libsdl-org/SDL/releases/download/release-2.30.7/SDL2-2.30.7.src.tar.gz"
	sdl2_cmd_src := exec.Command("curl", "-L", sdl2_src_url, "-o", "source.tar.gz")
	err = sdl2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdl2_cmd_direct := exec.Command("./binary")
	err = sdl2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libxcursor")
	exec.Command("latte", "install", "libxcursor").Run()
	fmt.Println("Instalando dependencia: libxscrnsaver")
	exec.Command("latte", "install", "libxscrnsaver").Run()
	fmt.Println("Instalando dependencia: libxxf86vm")
	exec.Command("latte", "install", "libxxf86vm").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
	fmt.Println("Instalando dependencia: xinput")
	exec.Command("latte", "install", "xinput").Run()
}
