package main

// Julius - Two-pass large vocabulary continuous speech recognition engine
// Homepage: https://github.com/julius-speech/julius

import (
	"fmt"
	
	"os/exec"
)

func installJulius() {
	// Método 1: Descargar y extraer .tar.gz
	julius_tar_url := "https://github.com/julius-speech/julius/archive/refs/tags/v4.6.tar.gz"
	julius_cmd_tar := exec.Command("curl", "-L", julius_tar_url, "-o", "package.tar.gz")
	err := julius_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	julius_zip_url := "https://github.com/julius-speech/julius/archive/refs/tags/v4.6.zip"
	julius_cmd_zip := exec.Command("curl", "-L", julius_zip_url, "-o", "package.zip")
	err = julius_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	julius_bin_url := "https://github.com/julius-speech/julius/archive/refs/tags/v4.6.bin"
	julius_cmd_bin := exec.Command("curl", "-L", julius_bin_url, "-o", "binary.bin")
	err = julius_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	julius_src_url := "https://github.com/julius-speech/julius/archive/refs/tags/v4.6.src.tar.gz"
	julius_cmd_src := exec.Command("curl", "-L", julius_src_url, "-o", "source.tar.gz")
	err = julius_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	julius_cmd_direct := exec.Command("./binary")
	err = julius_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
}
