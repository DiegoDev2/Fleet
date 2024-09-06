package main

// Espeak - Text to speech, software speech synthesizer
// Homepage: https://espeak.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installEspeak() {
	// Método 1: Descargar y extraer .tar.gz
	espeak_tar_url := "https://downloads.sourceforge.net/project/espeak/espeak/espeak-1.48/espeak-1.48.04-source.zip"
	espeak_cmd_tar := exec.Command("curl", "-L", espeak_tar_url, "-o", "package.tar.gz")
	err := espeak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	espeak_zip_url := "https://downloads.sourceforge.net/project/espeak/espeak/espeak-1.48/espeak-1.48.04-source.zip"
	espeak_cmd_zip := exec.Command("curl", "-L", espeak_zip_url, "-o", "package.zip")
	err = espeak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	espeak_bin_url := "https://downloads.sourceforge.net/project/espeak/espeak/espeak-1.48/espeak-1.48.04-source.zip"
	espeak_cmd_bin := exec.Command("curl", "-L", espeak_bin_url, "-o", "binary.bin")
	err = espeak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	espeak_src_url := "https://downloads.sourceforge.net/project/espeak/espeak/espeak-1.48/espeak-1.48.04-source.zip"
	espeak_cmd_src := exec.Command("curl", "-L", espeak_src_url, "-o", "source.tar.gz")
	err = espeak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	espeak_cmd_direct := exec.Command("./binary")
	err = espeak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: portaudio")
	exec.Command("latte", "install", "portaudio").Run()
}
