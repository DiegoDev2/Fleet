package main

// Abcmidi - Converts abc music notation files to MIDI files
// Homepage: https://ifdo.ca/~seymour/runabc/top.html

import (
	"fmt"
	
	"os/exec"
)

func installAbcmidi() {
	// Método 1: Descargar y extraer .tar.gz
	abcmidi_tar_url := "https://ifdo.ca/~seymour/runabc/abcMIDI-2024.08.13.zip"
	abcmidi_cmd_tar := exec.Command("curl", "-L", abcmidi_tar_url, "-o", "package.tar.gz")
	err := abcmidi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abcmidi_zip_url := "https://ifdo.ca/~seymour/runabc/abcMIDI-2024.08.13.zip"
	abcmidi_cmd_zip := exec.Command("curl", "-L", abcmidi_zip_url, "-o", "package.zip")
	err = abcmidi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abcmidi_bin_url := "https://ifdo.ca/~seymour/runabc/abcMIDI-2024.08.13.zip"
	abcmidi_cmd_bin := exec.Command("curl", "-L", abcmidi_bin_url, "-o", "binary.bin")
	err = abcmidi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abcmidi_src_url := "https://ifdo.ca/~seymour/runabc/abcMIDI-2024.08.13.zip"
	abcmidi_cmd_src := exec.Command("curl", "-L", abcmidi_src_url, "-o", "source.tar.gz")
	err = abcmidi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abcmidi_cmd_direct := exec.Command("./binary")
	err = abcmidi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
