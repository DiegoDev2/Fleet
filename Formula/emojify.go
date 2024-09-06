package main

// Emojify - Emoji on the command-line :scream:
// Homepage: https://github.com/mrowa44/emojify

import (
	"fmt"
	
	"os/exec"
)

func installEmojify() {
	// Método 1: Descargar y extraer .tar.gz
	emojify_tar_url := "https://github.com/mrowa44/emojify/archive/refs/tags/2.2.0.tar.gz"
	emojify_cmd_tar := exec.Command("curl", "-L", emojify_tar_url, "-o", "package.tar.gz")
	err := emojify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emojify_zip_url := "https://github.com/mrowa44/emojify/archive/refs/tags/2.2.0.zip"
	emojify_cmd_zip := exec.Command("curl", "-L", emojify_zip_url, "-o", "package.zip")
	err = emojify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emojify_bin_url := "https://github.com/mrowa44/emojify/archive/refs/tags/2.2.0.bin"
	emojify_cmd_bin := exec.Command("curl", "-L", emojify_bin_url, "-o", "binary.bin")
	err = emojify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emojify_src_url := "https://github.com/mrowa44/emojify/archive/refs/tags/2.2.0.src.tar.gz"
	emojify_cmd_src := exec.Command("curl", "-L", emojify_src_url, "-o", "source.tar.gz")
	err = emojify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emojify_cmd_direct := exec.Command("./binary")
	err = emojify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
}
