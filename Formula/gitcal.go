package main

// GitCal - GitHub-like contributions calendar but on the command-line
// Homepage: https://github.com/k4rthik/git-cal

import (
	"fmt"
	
	"os/exec"
)

func installGitCal() {
	// Método 1: Descargar y extraer .tar.gz
	gitcal_tar_url := "https://github.com/k4rthik/git-cal/archive/refs/tags/v0.9.1.tar.gz"
	gitcal_cmd_tar := exec.Command("curl", "-L", gitcal_tar_url, "-o", "package.tar.gz")
	err := gitcal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitcal_zip_url := "https://github.com/k4rthik/git-cal/archive/refs/tags/v0.9.1.zip"
	gitcal_cmd_zip := exec.Command("curl", "-L", gitcal_zip_url, "-o", "package.zip")
	err = gitcal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitcal_bin_url := "https://github.com/k4rthik/git-cal/archive/refs/tags/v0.9.1.bin"
	gitcal_cmd_bin := exec.Command("curl", "-L", gitcal_bin_url, "-o", "binary.bin")
	err = gitcal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitcal_src_url := "https://github.com/k4rthik/git-cal/archive/refs/tags/v0.9.1.src.tar.gz"
	gitcal_cmd_src := exec.Command("curl", "-L", gitcal_src_url, "-o", "source.tar.gz")
	err = gitcal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitcal_cmd_direct := exec.Command("./binary")
	err = gitcal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
