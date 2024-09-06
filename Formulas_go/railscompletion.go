package main

// RailsCompletion - Bash completion for Rails
// Homepage: https://github.com/mernen/completion-ruby

import (
	"fmt"
	
	"os/exec"
)

func installRailsCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	railscompletion_tar_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.tar.gz"
	railscompletion_cmd_tar := exec.Command("curl", "-L", railscompletion_tar_url, "-o", "package.tar.gz")
	err := railscompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	railscompletion_zip_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.zip"
	railscompletion_cmd_zip := exec.Command("curl", "-L", railscompletion_zip_url, "-o", "package.zip")
	err = railscompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	railscompletion_bin_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.bin"
	railscompletion_cmd_bin := exec.Command("curl", "-L", railscompletion_bin_url, "-o", "binary.bin")
	err = railscompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	railscompletion_src_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.src.tar.gz"
	railscompletion_cmd_src := exec.Command("curl", "-L", railscompletion_src_url, "-o", "source.tar.gz")
	err = railscompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	railscompletion_cmd_direct := exec.Command("./binary")
	err = railscompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
