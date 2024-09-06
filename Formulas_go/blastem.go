package main

// Blastem - Fast and accurate Genesis emulator
// Homepage: https://www.retrodev.com/blastem/

import (
	"fmt"
	
	"os/exec"
)

func installBlastem() {
	// Método 1: Descargar y extraer .tar.gz
	blastem_tar_url := "https://www.retrodev.com/repos/blastem/archive/v0.6.2.tar.gz"
	blastem_cmd_tar := exec.Command("curl", "-L", blastem_tar_url, "-o", "package.tar.gz")
	err := blastem_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blastem_zip_url := "https://www.retrodev.com/repos/blastem/archive/v0.6.2.zip"
	blastem_cmd_zip := exec.Command("curl", "-L", blastem_zip_url, "-o", "package.zip")
	err = blastem_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blastem_bin_url := "https://www.retrodev.com/repos/blastem/archive/v0.6.2.bin"
	blastem_cmd_bin := exec.Command("curl", "-L", blastem_bin_url, "-o", "binary.bin")
	err = blastem_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blastem_src_url := "https://www.retrodev.com/repos/blastem/archive/v0.6.2.src.tar.gz"
	blastem_cmd_src := exec.Command("curl", "-L", blastem_src_url, "-o", "source.tar.gz")
	err = blastem_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blastem_cmd_direct := exec.Command("./binary")
	err = blastem_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: glew")
exec.Command("latte", "install", "glew")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
