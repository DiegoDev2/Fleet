package main

// BrewCaskCompletion - Fish completion for brew-cask
// Homepage: https://github.com/xyb/homebrew-cask-completion

import (
	"fmt"
	
	"os/exec"
)

func installBrewCaskCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	brewcaskcompletion_tar_url := "https://github.com/xyb/homebrew-cask-completion/archive/refs/tags/v2.1.tar.gz"
	brewcaskcompletion_cmd_tar := exec.Command("curl", "-L", brewcaskcompletion_tar_url, "-o", "package.tar.gz")
	err := brewcaskcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brewcaskcompletion_zip_url := "https://github.com/xyb/homebrew-cask-completion/archive/refs/tags/v2.1.zip"
	brewcaskcompletion_cmd_zip := exec.Command("curl", "-L", brewcaskcompletion_zip_url, "-o", "package.zip")
	err = brewcaskcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brewcaskcompletion_bin_url := "https://github.com/xyb/homebrew-cask-completion/archive/refs/tags/v2.1.bin"
	brewcaskcompletion_cmd_bin := exec.Command("curl", "-L", brewcaskcompletion_bin_url, "-o", "binary.bin")
	err = brewcaskcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brewcaskcompletion_src_url := "https://github.com/xyb/homebrew-cask-completion/archive/refs/tags/v2.1.src.tar.gz"
	brewcaskcompletion_cmd_src := exec.Command("curl", "-L", brewcaskcompletion_src_url, "-o", "source.tar.gz")
	err = brewcaskcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brewcaskcompletion_cmd_direct := exec.Command("./binary")
	err = brewcaskcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
