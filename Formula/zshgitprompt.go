package main

// ZshGitPrompt - Informative git prompt for zsh
// Homepage: https://github.com/olivierverdier/zsh-git-prompt

import (
	"fmt"
	
	"os/exec"
)

func installZshGitPrompt() {
	// Método 1: Descargar y extraer .tar.gz
	zshgitprompt_tar_url := "https://github.com/olivierverdier/zsh-git-prompt/archive/refs/tags/v0.5.tar.gz"
	zshgitprompt_cmd_tar := exec.Command("curl", "-L", zshgitprompt_tar_url, "-o", "package.tar.gz")
	err := zshgitprompt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zshgitprompt_zip_url := "https://github.com/olivierverdier/zsh-git-prompt/archive/refs/tags/v0.5.zip"
	zshgitprompt_cmd_zip := exec.Command("curl", "-L", zshgitprompt_zip_url, "-o", "package.zip")
	err = zshgitprompt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zshgitprompt_bin_url := "https://github.com/olivierverdier/zsh-git-prompt/archive/refs/tags/v0.5.bin"
	zshgitprompt_cmd_bin := exec.Command("curl", "-L", zshgitprompt_bin_url, "-o", "binary.bin")
	err = zshgitprompt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zshgitprompt_src_url := "https://github.com/olivierverdier/zsh-git-prompt/archive/refs/tags/v0.5.src.tar.gz"
	zshgitprompt_cmd_src := exec.Command("curl", "-L", zshgitprompt_src_url, "-o", "source.tar.gz")
	err = zshgitprompt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zshgitprompt_cmd_direct := exec.Command("./binary")
	err = zshgitprompt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
