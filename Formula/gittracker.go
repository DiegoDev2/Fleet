package main

// GitTracker - Integrate Pivotal Tracker into your Git workflow
// Homepage: https://github.com/stevenharman/git_tracker

import (
	"fmt"
	
	"os/exec"
)

func installGitTracker() {
	// Método 1: Descargar y extraer .tar.gz
	gittracker_tar_url := "https://github.com/stevenharman/git_tracker/archive/refs/tags/v2.0.0.tar.gz"
	gittracker_cmd_tar := exec.Command("curl", "-L", gittracker_tar_url, "-o", "package.tar.gz")
	err := gittracker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gittracker_zip_url := "https://github.com/stevenharman/git_tracker/archive/refs/tags/v2.0.0.zip"
	gittracker_cmd_zip := exec.Command("curl", "-L", gittracker_zip_url, "-o", "package.zip")
	err = gittracker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gittracker_bin_url := "https://github.com/stevenharman/git_tracker/archive/refs/tags/v2.0.0.bin"
	gittracker_cmd_bin := exec.Command("curl", "-L", gittracker_bin_url, "-o", "binary.bin")
	err = gittracker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gittracker_src_url := "https://github.com/stevenharman/git_tracker/archive/refs/tags/v2.0.0.src.tar.gz"
	gittracker_cmd_src := exec.Command("curl", "-L", gittracker_src_url, "-o", "source.tar.gz")
	err = gittracker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gittracker_cmd_direct := exec.Command("./binary")
	err = gittracker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
