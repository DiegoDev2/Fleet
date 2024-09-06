package main

// Mupen64plus - Cross-platform plugin-based N64 emulator
// Homepage: https://www.mupen64plus.org/

import (
	"fmt"
	
	"os/exec"
)

func installMupen64plus() {
	// Método 1: Descargar y extraer .tar.gz
	mupen64plus_tar_url := "https://github.com/mupen64plus/mupen64plus-core/releases/download/2.5/mupen64plus-bundle-src-2.5.tar.gz"
	mupen64plus_cmd_tar := exec.Command("curl", "-L", mupen64plus_tar_url, "-o", "package.tar.gz")
	err := mupen64plus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mupen64plus_zip_url := "https://github.com/mupen64plus/mupen64plus-core/releases/download/2.5/mupen64plus-bundle-src-2.5.zip"
	mupen64plus_cmd_zip := exec.Command("curl", "-L", mupen64plus_zip_url, "-o", "package.zip")
	err = mupen64plus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mupen64plus_bin_url := "https://github.com/mupen64plus/mupen64plus-core/releases/download/2.5/mupen64plus-bundle-src-2.5.bin"
	mupen64plus_cmd_bin := exec.Command("curl", "-L", mupen64plus_bin_url, "-o", "binary.bin")
	err = mupen64plus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mupen64plus_src_url := "https://github.com/mupen64plus/mupen64plus-core/releases/download/2.5/mupen64plus-bundle-src-2.5.src.tar.gz"
	mupen64plus_cmd_src := exec.Command("curl", "-L", mupen64plus_src_url, "-o", "source.tar.gz")
	err = mupen64plus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mupen64plus_cmd_direct := exec.Command("./binary")
	err = mupen64plus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
