package main

// WhisperCpp - Port of OpenAI's Whisper model in C/C++
// Homepage: https://github.com/ggerganov/whisper.cpp

import (
	"fmt"
	
	"os/exec"
)

func installWhisperCpp() {
	// Método 1: Descargar y extraer .tar.gz
	whispercpp_tar_url := "https://github.com/ggerganov/whisper.cpp/archive/refs/tags/v1.5.4.tar.gz"
	whispercpp_cmd_tar := exec.Command("curl", "-L", whispercpp_tar_url, "-o", "package.tar.gz")
	err := whispercpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	whispercpp_zip_url := "https://github.com/ggerganov/whisper.cpp/archive/refs/tags/v1.5.4.zip"
	whispercpp_cmd_zip := exec.Command("curl", "-L", whispercpp_zip_url, "-o", "package.zip")
	err = whispercpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	whispercpp_bin_url := "https://github.com/ggerganov/whisper.cpp/archive/refs/tags/v1.5.4.bin"
	whispercpp_cmd_bin := exec.Command("curl", "-L", whispercpp_bin_url, "-o", "binary.bin")
	err = whispercpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	whispercpp_src_url := "https://github.com/ggerganov/whisper.cpp/archive/refs/tags/v1.5.4.src.tar.gz"
	whispercpp_cmd_src := exec.Command("curl", "-L", whispercpp_src_url, "-o", "source.tar.gz")
	err = whispercpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	whispercpp_cmd_direct := exec.Command("./binary")
	err = whispercpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
