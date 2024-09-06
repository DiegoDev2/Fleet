package main

// CondaZshCompletion - Zsh completion for conda
// Homepage: https://github.com/conda-incubator/conda-zsh-completion

import (
	"fmt"
	
	"os/exec"
)

func installCondaZshCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	condazshcompletion_tar_url := "https://github.com/conda-incubator/conda-zsh-completion/archive/refs/tags/v0.11.tar.gz"
	condazshcompletion_cmd_tar := exec.Command("curl", "-L", condazshcompletion_tar_url, "-o", "package.tar.gz")
	err := condazshcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	condazshcompletion_zip_url := "https://github.com/conda-incubator/conda-zsh-completion/archive/refs/tags/v0.11.zip"
	condazshcompletion_cmd_zip := exec.Command("curl", "-L", condazshcompletion_zip_url, "-o", "package.zip")
	err = condazshcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	condazshcompletion_bin_url := "https://github.com/conda-incubator/conda-zsh-completion/archive/refs/tags/v0.11.bin"
	condazshcompletion_cmd_bin := exec.Command("curl", "-L", condazshcompletion_bin_url, "-o", "binary.bin")
	err = condazshcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	condazshcompletion_src_url := "https://github.com/conda-incubator/conda-zsh-completion/archive/refs/tags/v0.11.src.tar.gz"
	condazshcompletion_cmd_src := exec.Command("curl", "-L", condazshcompletion_src_url, "-o", "source.tar.gz")
	err = condazshcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	condazshcompletion_cmd_direct := exec.Command("./binary")
	err = condazshcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
