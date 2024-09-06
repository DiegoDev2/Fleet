package main

// NestopiaUe - NES emulator
// Homepage: http://0ldsk00l.ca/nestopia/

import (
	"fmt"
	
	"os/exec"
)

func installNestopiaUe() {
	// Método 1: Descargar y extraer .tar.gz
	nestopiaue_tar_url := "https://github.com/0ldsk00l/nestopia/archive/refs/tags/1.52.1.tar.gz"
	nestopiaue_cmd_tar := exec.Command("curl", "-L", nestopiaue_tar_url, "-o", "package.tar.gz")
	err := nestopiaue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nestopiaue_zip_url := "https://github.com/0ldsk00l/nestopia/archive/refs/tags/1.52.1.zip"
	nestopiaue_cmd_zip := exec.Command("curl", "-L", nestopiaue_zip_url, "-o", "package.zip")
	err = nestopiaue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nestopiaue_bin_url := "https://github.com/0ldsk00l/nestopia/archive/refs/tags/1.52.1.bin"
	nestopiaue_cmd_bin := exec.Command("curl", "-L", nestopiaue_bin_url, "-o", "binary.bin")
	err = nestopiaue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nestopiaue_src_url := "https://github.com/0ldsk00l/nestopia/archive/refs/tags/1.52.1.src.tar.gz"
	nestopiaue_cmd_src := exec.Command("curl", "-L", nestopiaue_src_url, "-o", "source.tar.gz")
	err = nestopiaue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nestopiaue_cmd_direct := exec.Command("./binary")
	err = nestopiaue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fltk")
exec.Command("latte", "install", "fltk")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
