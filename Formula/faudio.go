package main

// Faudio - Accuracy-focused XAudio reimplementation for open platforms
// Homepage: https://fna-xna.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installFaudio() {
	// Método 1: Descargar y extraer .tar.gz
	faudio_tar_url := "https://github.com/FNA-XNA/FAudio/archive/refs/tags/24.09.tar.gz"
	faudio_cmd_tar := exec.Command("curl", "-L", faudio_tar_url, "-o", "package.tar.gz")
	err := faudio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	faudio_zip_url := "https://github.com/FNA-XNA/FAudio/archive/refs/tags/24.09.zip"
	faudio_cmd_zip := exec.Command("curl", "-L", faudio_zip_url, "-o", "package.zip")
	err = faudio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	faudio_bin_url := "https://github.com/FNA-XNA/FAudio/archive/refs/tags/24.09.bin"
	faudio_cmd_bin := exec.Command("curl", "-L", faudio_bin_url, "-o", "binary.bin")
	err = faudio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	faudio_src_url := "https://github.com/FNA-XNA/FAudio/archive/refs/tags/24.09.src.tar.gz"
	faudio_cmd_src := exec.Command("curl", "-L", faudio_src_url, "-o", "source.tar.gz")
	err = faudio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	faudio_cmd_direct := exec.Command("./binary")
	err = faudio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
}
