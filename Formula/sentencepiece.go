package main

// Sentencepiece - Unsupervised text tokenizer and detokenizer
// Homepage: https://github.com/google/sentencepiece

import (
	"fmt"
	
	"os/exec"
)

func installSentencepiece() {
	// Método 1: Descargar y extraer .tar.gz
	sentencepiece_tar_url := "https://github.com/google/sentencepiece/archive/refs/tags/v0.2.0.tar.gz"
	sentencepiece_cmd_tar := exec.Command("curl", "-L", sentencepiece_tar_url, "-o", "package.tar.gz")
	err := sentencepiece_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sentencepiece_zip_url := "https://github.com/google/sentencepiece/archive/refs/tags/v0.2.0.zip"
	sentencepiece_cmd_zip := exec.Command("curl", "-L", sentencepiece_zip_url, "-o", "package.zip")
	err = sentencepiece_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sentencepiece_bin_url := "https://github.com/google/sentencepiece/archive/refs/tags/v0.2.0.bin"
	sentencepiece_cmd_bin := exec.Command("curl", "-L", sentencepiece_bin_url, "-o", "binary.bin")
	err = sentencepiece_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sentencepiece_src_url := "https://github.com/google/sentencepiece/archive/refs/tags/v0.2.0.src.tar.gz"
	sentencepiece_cmd_src := exec.Command("curl", "-L", sentencepiece_src_url, "-o", "source.tar.gz")
	err = sentencepiece_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sentencepiece_cmd_direct := exec.Command("./binary")
	err = sentencepiece_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
