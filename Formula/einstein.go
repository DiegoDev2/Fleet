package main

// Einstein - Remake of the old DOS game Sherlock
// Homepage: https://github.com/lksj/einstein-puzzle

import (
	"fmt"
	
	"os/exec"
)

func installEinstein() {
	// Método 1: Descargar y extraer .tar.gz
	einstein_tar_url := "https://github.com/lksj/einstein-puzzle/archive/refs/tags/v2.1.1.tar.gz"
	einstein_cmd_tar := exec.Command("curl", "-L", einstein_tar_url, "-o", "package.tar.gz")
	err := einstein_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	einstein_zip_url := "https://github.com/lksj/einstein-puzzle/archive/refs/tags/v2.1.1.zip"
	einstein_cmd_zip := exec.Command("curl", "-L", einstein_zip_url, "-o", "package.zip")
	err = einstein_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	einstein_bin_url := "https://github.com/lksj/einstein-puzzle/archive/refs/tags/v2.1.1.bin"
	einstein_cmd_bin := exec.Command("curl", "-L", einstein_bin_url, "-o", "binary.bin")
	err = einstein_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	einstein_src_url := "https://github.com/lksj/einstein-puzzle/archive/refs/tags/v2.1.1.src.tar.gz"
	einstein_cmd_src := exec.Command("curl", "-L", einstein_src_url, "-o", "source.tar.gz")
	err = einstein_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	einstein_cmd_direct := exec.Command("./binary")
	err = einstein_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl_mixer")
	exec.Command("latte", "install", "sdl_mixer").Run()
	fmt.Println("Instalando dependencia: sdl_ttf")
	exec.Command("latte", "install", "sdl_ttf").Run()
}
