package main

// Portmidi - Cross-platform library for real-time MIDI I/O
// Homepage: https://github.com/PortMidi/portmidi

import (
	"fmt"
	
	"os/exec"
)

func installPortmidi() {
	// Método 1: Descargar y extraer .tar.gz
	portmidi_tar_url := "https://github.com/PortMidi/portmidi/archive/refs/tags/v2.0.4.tar.gz"
	portmidi_cmd_tar := exec.Command("curl", "-L", portmidi_tar_url, "-o", "package.tar.gz")
	err := portmidi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	portmidi_zip_url := "https://github.com/PortMidi/portmidi/archive/refs/tags/v2.0.4.zip"
	portmidi_cmd_zip := exec.Command("curl", "-L", portmidi_zip_url, "-o", "package.zip")
	err = portmidi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	portmidi_bin_url := "https://github.com/PortMidi/portmidi/archive/refs/tags/v2.0.4.bin"
	portmidi_cmd_bin := exec.Command("curl", "-L", portmidi_bin_url, "-o", "binary.bin")
	err = portmidi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	portmidi_src_url := "https://github.com/PortMidi/portmidi/archive/refs/tags/v2.0.4.src.tar.gz"
	portmidi_cmd_src := exec.Command("curl", "-L", portmidi_src_url, "-o", "source.tar.gz")
	err = portmidi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	portmidi_cmd_direct := exec.Command("./binary")
	err = portmidi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
