package main

// DsdaDoom - Fork of prboom+ with a focus on speedrunning
// Homepage: https://github.com/kraflab/dsda-doom

import (
	"fmt"
	
	"os/exec"
)

func installDsdaDoom() {
	// Método 1: Descargar y extraer .tar.gz
	dsdadoom_tar_url := "https://github.com/kraflab/dsda-doom/archive/refs/tags/v0.28.1.tar.gz"
	dsdadoom_cmd_tar := exec.Command("curl", "-L", dsdadoom_tar_url, "-o", "package.tar.gz")
	err := dsdadoom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dsdadoom_zip_url := "https://github.com/kraflab/dsda-doom/archive/refs/tags/v0.28.1.zip"
	dsdadoom_cmd_zip := exec.Command("curl", "-L", dsdadoom_zip_url, "-o", "package.zip")
	err = dsdadoom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dsdadoom_bin_url := "https://github.com/kraflab/dsda-doom/archive/refs/tags/v0.28.1.bin"
	dsdadoom_cmd_bin := exec.Command("curl", "-L", dsdadoom_bin_url, "-o", "binary.bin")
	err = dsdadoom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dsdadoom_src_url := "https://github.com/kraflab/dsda-doom/archive/refs/tags/v0.28.1.src.tar.gz"
	dsdadoom_cmd_src := exec.Command("curl", "-L", dsdadoom_src_url, "-o", "source.tar.gz")
	err = dsdadoom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dsdadoom_cmd_direct := exec.Command("./binary")
	err = dsdadoom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: dumb")
	exec.Command("latte", "install", "dumb").Run()
	fmt.Println("Instalando dependencia: fluid-synth")
	exec.Command("latte", "install", "fluid-synth").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: mad")
	exec.Command("latte", "install", "mad").Run()
	fmt.Println("Instalando dependencia: portmidi")
	exec.Command("latte", "install", "portmidi").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
