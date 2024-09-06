package main

// Portaudio - Cross-platform library for audio I/O
// Homepage: https://www.portaudio.com

import (
	"fmt"
	
	"os/exec"
)

func installPortaudio() {
	// Método 1: Descargar y extraer .tar.gz
	portaudio_tar_url := "https://files.portaudio.com/archives/pa_stable_v190700_20210406.tgz"
	portaudio_cmd_tar := exec.Command("curl", "-L", portaudio_tar_url, "-o", "package.tar.gz")
	err := portaudio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	portaudio_zip_url := "https://files.portaudio.com/archives/pa_stable_v190700_20210406.tgz"
	portaudio_cmd_zip := exec.Command("curl", "-L", portaudio_zip_url, "-o", "package.zip")
	err = portaudio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	portaudio_bin_url := "https://files.portaudio.com/archives/pa_stable_v190700_20210406.tgz"
	portaudio_cmd_bin := exec.Command("curl", "-L", portaudio_bin_url, "-o", "binary.bin")
	err = portaudio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	portaudio_src_url := "https://files.portaudio.com/archives/pa_stable_v190700_20210406.tgz"
	portaudio_cmd_src := exec.Command("curl", "-L", portaudio_src_url, "-o", "source.tar.gz")
	err = portaudio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	portaudio_cmd_direct := exec.Command("./binary")
	err = portaudio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
}
