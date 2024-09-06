package main

// Rubberband - Audio time stretcher tool and library
// Homepage: https://breakfastquay.com/rubberband/

import (
	"fmt"
	
	"os/exec"
)

func installRubberband() {
	// Método 1: Descargar y extraer .tar.gz
	rubberband_tar_url := "https://breakfastquay.com/files/releases/rubberband-3.3.0.tar.bz2"
	rubberband_cmd_tar := exec.Command("curl", "-L", rubberband_tar_url, "-o", "package.tar.gz")
	err := rubberband_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rubberband_zip_url := "https://breakfastquay.com/files/releases/rubberband-3.3.0.tar.bz2"
	rubberband_cmd_zip := exec.Command("curl", "-L", rubberband_zip_url, "-o", "package.zip")
	err = rubberband_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rubberband_bin_url := "https://breakfastquay.com/files/releases/rubberband-3.3.0.tar.bz2"
	rubberband_cmd_bin := exec.Command("curl", "-L", rubberband_bin_url, "-o", "binary.bin")
	err = rubberband_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rubberband_src_url := "https://breakfastquay.com/files/releases/rubberband-3.3.0.tar.bz2"
	rubberband_cmd_src := exec.Command("curl", "-L", rubberband_src_url, "-o", "source.tar.gz")
	err = rubberband_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rubberband_cmd_direct := exec.Command("./binary")
	err = rubberband_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: ladspa-sdk")
exec.Command("latte", "install", "ladspa-sdk")
	fmt.Println("Instalando dependencia: vamp-plugin-sdk")
exec.Command("latte", "install", "vamp-plugin-sdk")
}
