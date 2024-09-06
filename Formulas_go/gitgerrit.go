package main

// GitGerrit - Gerrit code review helper scripts
// Homepage: https://github.com/fbzhong/git-gerrit

import (
	"fmt"
	
	"os/exec"
)

func installGitGerrit() {
	// Método 1: Descargar y extraer .tar.gz
	gitgerrit_tar_url := "https://github.com/fbzhong/git-gerrit/archive/refs/tags/v0.3.0.tar.gz"
	gitgerrit_cmd_tar := exec.Command("curl", "-L", gitgerrit_tar_url, "-o", "package.tar.gz")
	err := gitgerrit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitgerrit_zip_url := "https://github.com/fbzhong/git-gerrit/archive/refs/tags/v0.3.0.zip"
	gitgerrit_cmd_zip := exec.Command("curl", "-L", gitgerrit_zip_url, "-o", "package.zip")
	err = gitgerrit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitgerrit_bin_url := "https://github.com/fbzhong/git-gerrit/archive/refs/tags/v0.3.0.bin"
	gitgerrit_cmd_bin := exec.Command("curl", "-L", gitgerrit_bin_url, "-o", "binary.bin")
	err = gitgerrit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitgerrit_src_url := "https://github.com/fbzhong/git-gerrit/archive/refs/tags/v0.3.0.src.tar.gz"
	gitgerrit_cmd_src := exec.Command("curl", "-L", gitgerrit_src_url, "-o", "source.tar.gz")
	err = gitgerrit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitgerrit_cmd_direct := exec.Command("./binary")
	err = gitgerrit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
