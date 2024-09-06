package main

// SwitchaudioOsx - Change macOS audio source from the command-line
// Homepage: https://github.com/deweller/switchaudio-osx/

import (
	"fmt"
	
	"os/exec"
)

func installSwitchaudioOsx() {
	// Método 1: Descargar y extraer .tar.gz
	switchaudioosx_tar_url := "https://github.com/deweller/switchaudio-osx/archive/refs/tags/1.2.2.tar.gz"
	switchaudioosx_cmd_tar := exec.Command("curl", "-L", switchaudioosx_tar_url, "-o", "package.tar.gz")
	err := switchaudioosx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	switchaudioosx_zip_url := "https://github.com/deweller/switchaudio-osx/archive/refs/tags/1.2.2.zip"
	switchaudioosx_cmd_zip := exec.Command("curl", "-L", switchaudioosx_zip_url, "-o", "package.zip")
	err = switchaudioosx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	switchaudioosx_bin_url := "https://github.com/deweller/switchaudio-osx/archive/refs/tags/1.2.2.bin"
	switchaudioosx_cmd_bin := exec.Command("curl", "-L", switchaudioosx_bin_url, "-o", "binary.bin")
	err = switchaudioosx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	switchaudioosx_src_url := "https://github.com/deweller/switchaudio-osx/archive/refs/tags/1.2.2.src.tar.gz"
	switchaudioosx_cmd_src := exec.Command("curl", "-L", switchaudioosx_src_url, "-o", "source.tar.gz")
	err = switchaudioosx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	switchaudioosx_cmd_direct := exec.Command("./binary")
	err = switchaudioosx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
