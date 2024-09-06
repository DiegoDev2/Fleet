package main

// Rtmidi - API for realtime MIDI input/output
// Homepage: https://www.music.mcgill.ca/~gary/rtmidi/

import (
	"fmt"
	
	"os/exec"
)

func installRtmidi() {
	// Método 1: Descargar y extraer .tar.gz
	rtmidi_tar_url := "https://www.music.mcgill.ca/~gary/rtmidi/release/rtmidi-6.0.0.tar.gz"
	rtmidi_cmd_tar := exec.Command("curl", "-L", rtmidi_tar_url, "-o", "package.tar.gz")
	err := rtmidi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rtmidi_zip_url := "https://www.music.mcgill.ca/~gary/rtmidi/release/rtmidi-6.0.0.zip"
	rtmidi_cmd_zip := exec.Command("curl", "-L", rtmidi_zip_url, "-o", "package.zip")
	err = rtmidi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rtmidi_bin_url := "https://www.music.mcgill.ca/~gary/rtmidi/release/rtmidi-6.0.0.bin"
	rtmidi_cmd_bin := exec.Command("curl", "-L", rtmidi_bin_url, "-o", "binary.bin")
	err = rtmidi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rtmidi_src_url := "https://www.music.mcgill.ca/~gary/rtmidi/release/rtmidi-6.0.0.src.tar.gz"
	rtmidi_cmd_src := exec.Command("curl", "-L", rtmidi_src_url, "-o", "source.tar.gz")
	err = rtmidi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rtmidi_cmd_direct := exec.Command("./binary")
	err = rtmidi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: jack")
	exec.Command("latte", "install", "jack").Run()
}
