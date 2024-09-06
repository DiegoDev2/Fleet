package main

// Carla - Audio plugin host supporting LADSPA, LV2, VST2/3, SF2 and more
// Homepage: https://kx.studio/Applications:Carla

import (
	"fmt"
	
	"os/exec"
)

func installCarla() {
	// Método 1: Descargar y extraer .tar.gz
	carla_tar_url := "https://github.com/falkTX/Carla/archive/refs/tags/v2.5.8.tar.gz"
	carla_cmd_tar := exec.Command("curl", "-L", carla_tar_url, "-o", "package.tar.gz")
	err := carla_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	carla_zip_url := "https://github.com/falkTX/Carla/archive/refs/tags/v2.5.8.zip"
	carla_cmd_zip := exec.Command("curl", "-L", carla_zip_url, "-o", "package.zip")
	err = carla_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	carla_bin_url := "https://github.com/falkTX/Carla/archive/refs/tags/v2.5.8.bin"
	carla_cmd_bin := exec.Command("curl", "-L", carla_bin_url, "-o", "binary.bin")
	err = carla_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	carla_src_url := "https://github.com/falkTX/Carla/archive/refs/tags/v2.5.8.src.tar.gz"
	carla_cmd_src := exec.Command("curl", "-L", carla_src_url, "-o", "source.tar.gz")
	err = carla_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	carla_cmd_direct := exec.Command("./binary")
	err = carla_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pyqt@5")
exec.Command("latte", "install", "pyqt@5")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: pyqt")
exec.Command("latte", "install", "pyqt")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fluid-synth")
exec.Command("latte", "install", "fluid-synth")
	fmt.Println("Instalando dependencia: liblo")
exec.Command("latte", "install", "liblo")
	fmt.Println("Instalando dependencia: libmagic")
exec.Command("latte", "install", "libmagic")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
