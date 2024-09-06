package main

// Dromeaudio - Small C++ audio manipulation and playback library
// Homepage: https://github.com/joshb/dromeaudio/

import (
	"fmt"
	
	"os/exec"
)

func installDromeaudio() {
	// Método 1: Descargar y extraer .tar.gz
	dromeaudio_tar_url := "https://github.com/joshb/DromeAudio/archive/refs/tags/v0.3.0.tar.gz"
	dromeaudio_cmd_tar := exec.Command("curl", "-L", dromeaudio_tar_url, "-o", "package.tar.gz")
	err := dromeaudio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dromeaudio_zip_url := "https://github.com/joshb/DromeAudio/archive/refs/tags/v0.3.0.zip"
	dromeaudio_cmd_zip := exec.Command("curl", "-L", dromeaudio_zip_url, "-o", "package.zip")
	err = dromeaudio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dromeaudio_bin_url := "https://github.com/joshb/DromeAudio/archive/refs/tags/v0.3.0.bin"
	dromeaudio_cmd_bin := exec.Command("curl", "-L", dromeaudio_bin_url, "-o", "binary.bin")
	err = dromeaudio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dromeaudio_src_url := "https://github.com/joshb/DromeAudio/archive/refs/tags/v0.3.0.src.tar.gz"
	dromeaudio_cmd_src := exec.Command("curl", "-L", dromeaudio_src_url, "-o", "source.tar.gz")
	err = dromeaudio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dromeaudio_cmd_direct := exec.Command("./binary")
	err = dromeaudio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
