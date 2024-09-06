package main

// Dirt - Experimental sample playback
// Homepage: https://github.com/tidalcycles/Dirt

import (
	"fmt"
	
	"os/exec"
)

func installDirt() {
	// Método 1: Descargar y extraer .tar.gz
	dirt_tar_url := "https://github.com/tidalcycles/Dirt/archive/refs/tags/1.1.tar.gz"
	dirt_cmd_tar := exec.Command("curl", "-L", dirt_tar_url, "-o", "package.tar.gz")
	err := dirt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dirt_zip_url := "https://github.com/tidalcycles/Dirt/archive/refs/tags/1.1.zip"
	dirt_cmd_zip := exec.Command("curl", "-L", dirt_zip_url, "-o", "package.zip")
	err = dirt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dirt_bin_url := "https://github.com/tidalcycles/Dirt/archive/refs/tags/1.1.bin"
	dirt_cmd_bin := exec.Command("curl", "-L", dirt_bin_url, "-o", "binary.bin")
	err = dirt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dirt_src_url := "https://github.com/tidalcycles/Dirt/archive/refs/tags/1.1.src.tar.gz"
	dirt_cmd_src := exec.Command("curl", "-L", dirt_src_url, "-o", "source.tar.gz")
	err = dirt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dirt_cmd_direct := exec.Command("./binary")
	err = dirt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: liblo")
exec.Command("latte", "install", "liblo")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
}
