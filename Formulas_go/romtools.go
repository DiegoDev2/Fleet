package main

// RomTools - Tools for Multiple Arcade Machine Emulator
// Homepage: https://mamedev.org/

import (
	"fmt"
	
	"os/exec"
)

func installRomTools() {
	// Método 1: Descargar y extraer .tar.gz
	romtools_tar_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.tar.gz"
	romtools_cmd_tar := exec.Command("curl", "-L", romtools_tar_url, "-o", "package.tar.gz")
	err := romtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	romtools_zip_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.zip"
	romtools_cmd_zip := exec.Command("curl", "-L", romtools_zip_url, "-o", "package.zip")
	err = romtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	romtools_bin_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.bin"
	romtools_cmd_bin := exec.Command("curl", "-L", romtools_bin_url, "-o", "binary.bin")
	err = romtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	romtools_src_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.src.tar.gz"
	romtools_cmd_src := exec.Command("curl", "-L", romtools_src_url, "-o", "source.tar.gz")
	err = romtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	romtools_cmd_direct := exec.Command("./binary")
	err = romtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asio")
exec.Command("latte", "install", "asio")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: utf8proc")
exec.Command("latte", "install", "utf8proc")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: portmidi")
exec.Command("latte", "install", "portmidi")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: sdl2_ttf")
exec.Command("latte", "install", "sdl2_ttf")
}
