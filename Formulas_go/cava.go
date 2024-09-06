package main

// Cava - Console-based Audio Visualizer for ALSA
// Homepage: https://github.com/karlstav/cava

import (
	"fmt"
	
	"os/exec"
)

func installCava() {
	// Método 1: Descargar y extraer .tar.gz
	cava_tar_url := "https://github.com/karlstav/cava/archive/refs/tags/0.10.2.tar.gz"
	cava_cmd_tar := exec.Command("curl", "-L", cava_tar_url, "-o", "package.tar.gz")
	err := cava_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cava_zip_url := "https://github.com/karlstav/cava/archive/refs/tags/0.10.2.zip"
	cava_cmd_zip := exec.Command("curl", "-L", cava_zip_url, "-o", "package.zip")
	err = cava_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cava_bin_url := "https://github.com/karlstav/cava/archive/refs/tags/0.10.2.bin"
	cava_cmd_bin := exec.Command("curl", "-L", cava_bin_url, "-o", "binary.bin")
	err = cava_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cava_src_url := "https://github.com/karlstav/cava/archive/refs/tags/0.10.2.src.tar.gz"
	cava_cmd_src := exec.Command("curl", "-L", cava_src_url, "-o", "source.tar.gz")
	err = cava_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cava_cmd_direct := exec.Command("./binary")
	err = cava_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: iniparser")
exec.Command("latte", "install", "iniparser")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
}
