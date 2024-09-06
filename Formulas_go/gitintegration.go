package main

// GitIntegration - Manage git integration branches
// Homepage: https://johnkeeping.github.io/git-integration/

import (
	"fmt"
	
	"os/exec"
)

func installGitIntegration() {
	// Método 1: Descargar y extraer .tar.gz
	gitintegration_tar_url := "https://github.com/johnkeeping/git-integration/archive/refs/tags/v0.4.tar.gz"
	gitintegration_cmd_tar := exec.Command("curl", "-L", gitintegration_tar_url, "-o", "package.tar.gz")
	err := gitintegration_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitintegration_zip_url := "https://github.com/johnkeeping/git-integration/archive/refs/tags/v0.4.zip"
	gitintegration_cmd_zip := exec.Command("curl", "-L", gitintegration_zip_url, "-o", "package.zip")
	err = gitintegration_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitintegration_bin_url := "https://github.com/johnkeeping/git-integration/archive/refs/tags/v0.4.bin"
	gitintegration_cmd_bin := exec.Command("curl", "-L", gitintegration_bin_url, "-o", "binary.bin")
	err = gitintegration_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitintegration_src_url := "https://github.com/johnkeeping/git-integration/archive/refs/tags/v0.4.src.tar.gz"
	gitintegration_cmd_src := exec.Command("curl", "-L", gitintegration_src_url, "-o", "source.tar.gz")
	err = gitintegration_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitintegration_cmd_direct := exec.Command("./binary")
	err = gitintegration_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
