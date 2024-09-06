package main

// Brogue - Roguelike game
// Homepage: https://sites.google.com/site/broguegame/

import (
	"fmt"
	
	"os/exec"
)

func installBrogue() {
	// Método 1: Descargar y extraer .tar.gz
	brogue_tar_url := "https://github.com/tmewett/BrogueCE/archive/refs/tags/v1.14.1.tar.gz"
	brogue_cmd_tar := exec.Command("curl", "-L", brogue_tar_url, "-o", "package.tar.gz")
	err := brogue_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brogue_zip_url := "https://github.com/tmewett/BrogueCE/archive/refs/tags/v1.14.1.zip"
	brogue_cmd_zip := exec.Command("curl", "-L", brogue_zip_url, "-o", "package.zip")
	err = brogue_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brogue_bin_url := "https://github.com/tmewett/BrogueCE/archive/refs/tags/v1.14.1.bin"
	brogue_cmd_bin := exec.Command("curl", "-L", brogue_bin_url, "-o", "binary.bin")
	err = brogue_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brogue_src_url := "https://github.com/tmewett/BrogueCE/archive/refs/tags/v1.14.1.src.tar.gz"
	brogue_cmd_src := exec.Command("curl", "-L", brogue_src_url, "-o", "source.tar.gz")
	err = brogue_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brogue_cmd_direct := exec.Command("./binary")
	err = brogue_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_image")
	exec.Command("latte", "install", "sdl2_image").Run()
}
