package main

// Ecasound - Multitrack-capable audio recorder and effect processor
// Homepage: https://nosignal.fi/ecasound/

import (
	"fmt"
	
	"os/exec"
)

func installEcasound() {
	// Método 1: Descargar y extraer .tar.gz
	ecasound_tar_url := "https://nosignal.fi/download/ecasound-2.9.3.tar.gz"
	ecasound_cmd_tar := exec.Command("curl", "-L", ecasound_tar_url, "-o", "package.tar.gz")
	err := ecasound_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ecasound_zip_url := "https://nosignal.fi/download/ecasound-2.9.3.zip"
	ecasound_cmd_zip := exec.Command("curl", "-L", ecasound_zip_url, "-o", "package.zip")
	err = ecasound_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ecasound_bin_url := "https://nosignal.fi/download/ecasound-2.9.3.bin"
	ecasound_cmd_bin := exec.Command("curl", "-L", ecasound_bin_url, "-o", "binary.bin")
	err = ecasound_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ecasound_src_url := "https://nosignal.fi/download/ecasound-2.9.3.src.tar.gz"
	ecasound_cmd_src := exec.Command("curl", "-L", ecasound_src_url, "-o", "source.tar.gz")
	err = ecasound_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ecasound_cmd_direct := exec.Command("./binary")
	err = ecasound_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jack")
	exec.Command("latte", "install", "jack").Run()
	fmt.Println("Instalando dependencia: libsamplerate")
	exec.Command("latte", "install", "libsamplerate").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
