package main

// GitGame - Game for git to guess who made which commit
// Homepage: https://github.com/jsomers/git-game

import (
	"fmt"
	
	"os/exec"
)

func installGitGame() {
	// Método 1: Descargar y extraer .tar.gz
	gitgame_tar_url := "https://github.com/jsomers/git-game/archive/refs/tags/1.3.tar.gz"
	gitgame_cmd_tar := exec.Command("curl", "-L", gitgame_tar_url, "-o", "package.tar.gz")
	err := gitgame_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitgame_zip_url := "https://github.com/jsomers/git-game/archive/refs/tags/1.3.zip"
	gitgame_cmd_zip := exec.Command("curl", "-L", gitgame_zip_url, "-o", "package.zip")
	err = gitgame_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitgame_bin_url := "https://github.com/jsomers/git-game/archive/refs/tags/1.3.bin"
	gitgame_cmd_bin := exec.Command("curl", "-L", gitgame_bin_url, "-o", "binary.bin")
	err = gitgame_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitgame_src_url := "https://github.com/jsomers/git-game/archive/refs/tags/1.3.src.tar.gz"
	gitgame_cmd_src := exec.Command("curl", "-L", gitgame_src_url, "-o", "source.tar.gz")
	err = gitgame_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitgame_cmd_direct := exec.Command("./binary")
	err = gitgame_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
