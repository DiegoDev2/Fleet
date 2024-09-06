package main

// GemCompletion - Bash completion for gem
// Homepage: https://github.com/mernen/completion-ruby

import (
	"fmt"
	
	"os/exec"
)

func installGemCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	gemcompletion_tar_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.tar.gz"
	gemcompletion_cmd_tar := exec.Command("curl", "-L", gemcompletion_tar_url, "-o", "package.tar.gz")
	err := gemcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gemcompletion_zip_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.zip"
	gemcompletion_cmd_zip := exec.Command("curl", "-L", gemcompletion_zip_url, "-o", "package.zip")
	err = gemcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gemcompletion_bin_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.bin"
	gemcompletion_cmd_bin := exec.Command("curl", "-L", gemcompletion_bin_url, "-o", "binary.bin")
	err = gemcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gemcompletion_src_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.src.tar.gz"
	gemcompletion_cmd_src := exec.Command("curl", "-L", gemcompletion_src_url, "-o", "source.tar.gz")
	err = gemcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gemcompletion_cmd_direct := exec.Command("./binary")
	err = gemcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
