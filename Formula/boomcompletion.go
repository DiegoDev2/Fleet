package main

// BoomCompletion - Bash and Zsh completion for Boom
// Homepage: https://zachholman.com/boom

import (
	"fmt"
	
	"os/exec"
)

func installBoomCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	boomcompletion_tar_url := "https://github.com/holman/boom/archive/refs/tags/v0.5.0.tar.gz"
	boomcompletion_cmd_tar := exec.Command("curl", "-L", boomcompletion_tar_url, "-o", "package.tar.gz")
	err := boomcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boomcompletion_zip_url := "https://github.com/holman/boom/archive/refs/tags/v0.5.0.zip"
	boomcompletion_cmd_zip := exec.Command("curl", "-L", boomcompletion_zip_url, "-o", "package.zip")
	err = boomcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boomcompletion_bin_url := "https://github.com/holman/boom/archive/refs/tags/v0.5.0.bin"
	boomcompletion_cmd_bin := exec.Command("curl", "-L", boomcompletion_bin_url, "-o", "binary.bin")
	err = boomcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boomcompletion_src_url := "https://github.com/holman/boom/archive/refs/tags/v0.5.0.src.tar.gz"
	boomcompletion_cmd_src := exec.Command("curl", "-L", boomcompletion_src_url, "-o", "source.tar.gz")
	err = boomcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boomcompletion_cmd_direct := exec.Command("./binary")
	err = boomcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
