package main

// Homeshick - Git dotfiles synchronizer written in bash
// Homepage: https://github.com/andsens/homeshick

import (
	"fmt"
	
	"os/exec"
)

func installHomeshick() {
	// Método 1: Descargar y extraer .tar.gz
	homeshick_tar_url := "https://github.com/andsens/homeshick/archive/refs/tags/v2.0.1.tar.gz"
	homeshick_cmd_tar := exec.Command("curl", "-L", homeshick_tar_url, "-o", "package.tar.gz")
	err := homeshick_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	homeshick_zip_url := "https://github.com/andsens/homeshick/archive/refs/tags/v2.0.1.zip"
	homeshick_cmd_zip := exec.Command("curl", "-L", homeshick_zip_url, "-o", "package.zip")
	err = homeshick_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	homeshick_bin_url := "https://github.com/andsens/homeshick/archive/refs/tags/v2.0.1.bin"
	homeshick_cmd_bin := exec.Command("curl", "-L", homeshick_bin_url, "-o", "binary.bin")
	err = homeshick_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	homeshick_src_url := "https://github.com/andsens/homeshick/archive/refs/tags/v2.0.1.src.tar.gz"
	homeshick_cmd_src := exec.Command("curl", "-L", homeshick_src_url, "-o", "source.tar.gz")
	err = homeshick_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	homeshick_cmd_direct := exec.Command("./binary")
	err = homeshick_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
