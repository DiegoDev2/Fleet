package main

// GitStandup - Git extension to generate reports for standup meetings
// Homepage: https://github.com/kamranahmedse/git-standup

import (
	"fmt"
	
	"os/exec"
)

func installGitStandup() {
	// Método 1: Descargar y extraer .tar.gz
	gitstandup_tar_url := "https://github.com/kamranahmedse/git-standup/archive/refs/tags/2.3.2.tar.gz"
	gitstandup_cmd_tar := exec.Command("curl", "-L", gitstandup_tar_url, "-o", "package.tar.gz")
	err := gitstandup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitstandup_zip_url := "https://github.com/kamranahmedse/git-standup/archive/refs/tags/2.3.2.zip"
	gitstandup_cmd_zip := exec.Command("curl", "-L", gitstandup_zip_url, "-o", "package.zip")
	err = gitstandup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitstandup_bin_url := "https://github.com/kamranahmedse/git-standup/archive/refs/tags/2.3.2.bin"
	gitstandup_cmd_bin := exec.Command("curl", "-L", gitstandup_bin_url, "-o", "binary.bin")
	err = gitstandup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitstandup_src_url := "https://github.com/kamranahmedse/git-standup/archive/refs/tags/2.3.2.src.tar.gz"
	gitstandup_cmd_src := exec.Command("curl", "-L", gitstandup_src_url, "-o", "source.tar.gz")
	err = gitstandup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitstandup_cmd_direct := exec.Command("./binary")
	err = gitstandup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
