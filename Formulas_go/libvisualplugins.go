package main

// LibvisualPlugins - Audio Visualization tool and library
// Homepage: https://github.com/Libvisual/libvisual

import (
	"fmt"
	
	"os/exec"
)

func installLibvisualPlugins() {
	// Método 1: Descargar y extraer .tar.gz
	libvisualplugins_tar_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-plugins-0.4.2/libvisual-plugins-0.4.2.tar.gz"
	libvisualplugins_cmd_tar := exec.Command("curl", "-L", libvisualplugins_tar_url, "-o", "package.tar.gz")
	err := libvisualplugins_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvisualplugins_zip_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-plugins-0.4.2/libvisual-plugins-0.4.2.zip"
	libvisualplugins_cmd_zip := exec.Command("curl", "-L", libvisualplugins_zip_url, "-o", "package.zip")
	err = libvisualplugins_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvisualplugins_bin_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-plugins-0.4.2/libvisual-plugins-0.4.2.bin"
	libvisualplugins_cmd_bin := exec.Command("curl", "-L", libvisualplugins_bin_url, "-o", "binary.bin")
	err = libvisualplugins_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvisualplugins_src_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-plugins-0.4.2/libvisual-plugins-0.4.2.src.tar.gz"
	libvisualplugins_cmd_src := exec.Command("curl", "-L", libvisualplugins_src_url, "-o", "source.tar.gz")
	err = libvisualplugins_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvisualplugins_cmd_direct := exec.Command("./binary")
	err = libvisualplugins_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xorg-server")
exec.Command("latte", "install", "xorg-server")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: libvisual")
exec.Command("latte", "install", "libvisual")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
}
