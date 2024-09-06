package main

// Stella - Atari 2600 VCS emulator
// Homepage: https://stella-emu.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installStella() {
	// Método 1: Descargar y extraer .tar.gz
	stella_tar_url := "https://github.com/stella-emu/stella/archive/refs/tags/6.7.1.tar.gz"
	stella_cmd_tar := exec.Command("curl", "-L", stella_tar_url, "-o", "package.tar.gz")
	err := stella_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stella_zip_url := "https://github.com/stella-emu/stella/archive/refs/tags/6.7.1.zip"
	stella_cmd_zip := exec.Command("curl", "-L", stella_zip_url, "-o", "package.zip")
	err = stella_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stella_bin_url := "https://github.com/stella-emu/stella/archive/refs/tags/6.7.1.bin"
	stella_cmd_bin := exec.Command("curl", "-L", stella_bin_url, "-o", "binary.bin")
	err = stella_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stella_src_url := "https://github.com/stella-emu/stella/archive/refs/tags/6.7.1.src.tar.gz"
	stella_cmd_src := exec.Command("curl", "-L", stella_src_url, "-o", "source.tar.gz")
	err = stella_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stella_cmd_direct := exec.Command("./binary")
	err = stella_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
