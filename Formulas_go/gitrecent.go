package main

// GitRecent - See your latest local git branches, formatted real fancy
// Homepage: https://github.com/paulirish/git-recent

import (
	"fmt"
	
	"os/exec"
)

func installGitRecent() {
	// Método 1: Descargar y extraer .tar.gz
	gitrecent_tar_url := "https://github.com/paulirish/git-recent/archive/refs/tags/v1.1.1.tar.gz"
	gitrecent_cmd_tar := exec.Command("curl", "-L", gitrecent_tar_url, "-o", "package.tar.gz")
	err := gitrecent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitrecent_zip_url := "https://github.com/paulirish/git-recent/archive/refs/tags/v1.1.1.zip"
	gitrecent_cmd_zip := exec.Command("curl", "-L", gitrecent_zip_url, "-o", "package.zip")
	err = gitrecent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitrecent_bin_url := "https://github.com/paulirish/git-recent/archive/refs/tags/v1.1.1.bin"
	gitrecent_cmd_bin := exec.Command("curl", "-L", gitrecent_bin_url, "-o", "binary.bin")
	err = gitrecent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitrecent_src_url := "https://github.com/paulirish/git-recent/archive/refs/tags/v1.1.1.src.tar.gz"
	gitrecent_cmd_src := exec.Command("curl", "-L", gitrecent_src_url, "-o", "source.tar.gz")
	err = gitrecent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitrecent_cmd_direct := exec.Command("./binary")
	err = gitrecent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
