package main

// Openmsx - MSX emulator
// Homepage: https://openmsx.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenmsx() {
	// Método 1: Descargar y extraer .tar.gz
	openmsx_tar_url := "https://github.com/openMSX/openMSX/releases/download/RELEASE_19_1/openmsx-19.1.tar.gz"
	openmsx_cmd_tar := exec.Command("curl", "-L", openmsx_tar_url, "-o", "package.tar.gz")
	err := openmsx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openmsx_zip_url := "https://github.com/openMSX/openMSX/releases/download/RELEASE_19_1/openmsx-19.1.zip"
	openmsx_cmd_zip := exec.Command("curl", "-L", openmsx_zip_url, "-o", "package.zip")
	err = openmsx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openmsx_bin_url := "https://github.com/openMSX/openMSX/releases/download/RELEASE_19_1/openmsx-19.1.bin"
	openmsx_cmd_bin := exec.Command("curl", "-L", openmsx_bin_url, "-o", "binary.bin")
	err = openmsx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openmsx_src_url := "https://github.com/openMSX/openMSX/releases/download/RELEASE_19_1/openmsx-19.1.src.tar.gz"
	openmsx_cmd_src := exec.Command("curl", "-L", openmsx_src_url, "-o", "source.tar.gz")
	err = openmsx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openmsx_cmd_direct := exec.Command("./binary")
	err = openmsx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_ttf")
exec.Command("latte", "install", "sdl2_ttf")
	fmt.Println("Instalando dependencia: theora")
exec.Command("latte", "install", "theora")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
