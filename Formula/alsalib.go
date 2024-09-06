package main

// AlsaLib - Provides audio and MIDI functionality to the Linux operating system
// Homepage: https://www.alsa-project.org/

import (
	"fmt"
	
	"os/exec"
)

func installAlsaLib() {
	// Método 1: Descargar y extraer .tar.gz
	alsalib_tar_url := "https://www.alsa-project.org/files/pub/lib/alsa-lib-1.2.12.tar.bz2"
	alsalib_cmd_tar := exec.Command("curl", "-L", alsalib_tar_url, "-o", "package.tar.gz")
	err := alsalib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alsalib_zip_url := "https://www.alsa-project.org/files/pub/lib/alsa-lib-1.2.12.tar.bz2"
	alsalib_cmd_zip := exec.Command("curl", "-L", alsalib_zip_url, "-o", "package.zip")
	err = alsalib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alsalib_bin_url := "https://www.alsa-project.org/files/pub/lib/alsa-lib-1.2.12.tar.bz2"
	alsalib_cmd_bin := exec.Command("curl", "-L", alsalib_bin_url, "-o", "binary.bin")
	err = alsalib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alsalib_src_url := "https://www.alsa-project.org/files/pub/lib/alsa-lib-1.2.12.tar.bz2"
	alsalib_cmd_src := exec.Command("curl", "-L", alsalib_src_url, "-o", "source.tar.gz")
	err = alsalib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alsalib_cmd_direct := exec.Command("./binary")
	err = alsalib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
