package main

// Spaceship - Zsh prompt for Astronauts
// Homepage: https://spaceship-prompt.sh

import (
	"fmt"
	
	"os/exec"
)

func installSpaceship() {
	// Método 1: Descargar y extraer .tar.gz
	spaceship_tar_url := "https://github.com/spaceship-prompt/spaceship-prompt/archive/refs/tags/v4.16.0.tar.gz"
	spaceship_cmd_tar := exec.Command("curl", "-L", spaceship_tar_url, "-o", "package.tar.gz")
	err := spaceship_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spaceship_zip_url := "https://github.com/spaceship-prompt/spaceship-prompt/archive/refs/tags/v4.16.0.zip"
	spaceship_cmd_zip := exec.Command("curl", "-L", spaceship_zip_url, "-o", "package.zip")
	err = spaceship_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spaceship_bin_url := "https://github.com/spaceship-prompt/spaceship-prompt/archive/refs/tags/v4.16.0.bin"
	spaceship_cmd_bin := exec.Command("curl", "-L", spaceship_bin_url, "-o", "binary.bin")
	err = spaceship_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spaceship_src_url := "https://github.com/spaceship-prompt/spaceship-prompt/archive/refs/tags/v4.16.0.src.tar.gz"
	spaceship_cmd_src := exec.Command("curl", "-L", spaceship_src_url, "-o", "source.tar.gz")
	err = spaceship_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spaceship_cmd_direct := exec.Command("./binary")
	err = spaceship_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zsh-async")
exec.Command("latte", "install", "zsh-async")
}
