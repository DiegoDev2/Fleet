package main

// Wildmidi - Simple software midi player
// Homepage: https://github.com/Mindwerks/wildmidi

import (
	"fmt"
	
	"os/exec"
)

func installWildmidi() {
	// Método 1: Descargar y extraer .tar.gz
	wildmidi_tar_url := "https://github.com/Mindwerks/wildmidi/archive/refs/tags/wildmidi-0.4.6.tar.gz"
	wildmidi_cmd_tar := exec.Command("curl", "-L", wildmidi_tar_url, "-o", "package.tar.gz")
	err := wildmidi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wildmidi_zip_url := "https://github.com/Mindwerks/wildmidi/archive/refs/tags/wildmidi-0.4.6.zip"
	wildmidi_cmd_zip := exec.Command("curl", "-L", wildmidi_zip_url, "-o", "package.zip")
	err = wildmidi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wildmidi_bin_url := "https://github.com/Mindwerks/wildmidi/archive/refs/tags/wildmidi-0.4.6.bin"
	wildmidi_cmd_bin := exec.Command("curl", "-L", wildmidi_bin_url, "-o", "binary.bin")
	err = wildmidi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wildmidi_src_url := "https://github.com/Mindwerks/wildmidi/archive/refs/tags/wildmidi-0.4.6.src.tar.gz"
	wildmidi_cmd_src := exec.Command("curl", "-L", wildmidi_src_url, "-o", "source.tar.gz")
	err = wildmidi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wildmidi_cmd_direct := exec.Command("./binary")
	err = wildmidi_cmd_direct.Run()
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
