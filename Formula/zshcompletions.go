package main

// ZshCompletions - Additional completion definitions for zsh
// Homepage: https://github.com/zsh-users/zsh-completions

import (
	"fmt"
	
	"os/exec"
)

func installZshCompletions() {
	// Método 1: Descargar y extraer .tar.gz
	zshcompletions_tar_url := "https://github.com/zsh-users/zsh-completions/archive/refs/tags/0.35.0.tar.gz"
	zshcompletions_cmd_tar := exec.Command("curl", "-L", zshcompletions_tar_url, "-o", "package.tar.gz")
	err := zshcompletions_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshcompletions_zip_url := "https://github.com/zsh-users/zsh-completions/archive/refs/tags/0.35.0.zip"
	zshcompletions_cmd_zip := exec.Command("curl", "-L", zshcompletions_zip_url, "-o", "package.zip")
	err = zshcompletions_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshcompletions_bin_url := "https://github.com/zsh-users/zsh-completions/archive/refs/tags/0.35.0.bin"
	zshcompletions_cmd_bin := exec.Command("curl", "-L", zshcompletions_bin_url, "-o", "binary.bin")
	err = zshcompletions_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshcompletions_src_url := "https://github.com/zsh-users/zsh-completions/archive/refs/tags/0.35.0.src.tar.gz"
	zshcompletions_cmd_src := exec.Command("curl", "-L", zshcompletions_src_url, "-o", "source.tar.gz")
	err = zshcompletions_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshcompletions_cmd_direct := exec.Command("./binary")
	err = zshcompletions_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
