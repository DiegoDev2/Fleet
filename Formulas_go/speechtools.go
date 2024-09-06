package main

// SpeechTools - C++ speech software library from the University of Edinburgh
// Homepage: http://festvox.org/docs/speech_tools-2.4.0/

import (
	"fmt"
	
	"os/exec"
)

func installSpeechTools() {
	// Método 1: Descargar y extraer .tar.gz
	speechtools_tar_url := "http://festvox.org/packed/festival/2.5/speech_tools-2.5.0-release.tar.gz"
	speechtools_cmd_tar := exec.Command("curl", "-L", speechtools_tar_url, "-o", "package.tar.gz")
	err := speechtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	speechtools_zip_url := "http://festvox.org/packed/festival/2.5/speech_tools-2.5.0-release.zip"
	speechtools_cmd_zip := exec.Command("curl", "-L", speechtools_zip_url, "-o", "package.zip")
	err = speechtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	speechtools_bin_url := "http://festvox.org/packed/festival/2.5/speech_tools-2.5.0-release.bin"
	speechtools_cmd_bin := exec.Command("curl", "-L", speechtools_bin_url, "-o", "binary.bin")
	err = speechtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	speechtools_src_url := "http://festvox.org/packed/festival/2.5/speech_tools-2.5.0-release.src.tar.gz"
	speechtools_cmd_src := exec.Command("curl", "-L", speechtools_src_url, "-o", "source.tar.gz")
	err = speechtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	speechtools_cmd_direct := exec.Command("./binary")
	err = speechtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
