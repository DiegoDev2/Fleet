package main

// BundlerCompletion - Bash completion for Bundler
// Homepage: https://github.com/mernen/completion-ruby

import (
	"fmt"
	
	"os/exec"
)

func installBundlerCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	bundlercompletion_tar_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.tar.gz"
	bundlercompletion_cmd_tar := exec.Command("curl", "-L", bundlercompletion_tar_url, "-o", "package.tar.gz")
	err := bundlercompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bundlercompletion_zip_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.zip"
	bundlercompletion_cmd_zip := exec.Command("curl", "-L", bundlercompletion_zip_url, "-o", "package.zip")
	err = bundlercompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bundlercompletion_bin_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.bin"
	bundlercompletion_cmd_bin := exec.Command("curl", "-L", bundlercompletion_bin_url, "-o", "binary.bin")
	err = bundlercompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bundlercompletion_src_url := "https://github.com/mernen/completion-ruby/archive/refs/tags/v1.0.2.src.tar.gz"
	bundlercompletion_cmd_src := exec.Command("curl", "-L", bundlercompletion_src_url, "-o", "source.tar.gz")
	err = bundlercompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bundlercompletion_cmd_direct := exec.Command("./binary")
	err = bundlercompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
