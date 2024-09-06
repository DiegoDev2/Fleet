package main

// GitFresh - Utility to keep git repos fresh
// Homepage: https://github.com/imsky/git-fresh

import (
	"fmt"
	
	"os/exec"
)

func installGitFresh() {
	// Método 1: Descargar y extraer .tar.gz
	gitfresh_tar_url := "https://github.com/imsky/git-fresh/archive/refs/tags/v1.13.0.tar.gz"
	gitfresh_cmd_tar := exec.Command("curl", "-L", gitfresh_tar_url, "-o", "package.tar.gz")
	err := gitfresh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitfresh_zip_url := "https://github.com/imsky/git-fresh/archive/refs/tags/v1.13.0.zip"
	gitfresh_cmd_zip := exec.Command("curl", "-L", gitfresh_zip_url, "-o", "package.zip")
	err = gitfresh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitfresh_bin_url := "https://github.com/imsky/git-fresh/archive/refs/tags/v1.13.0.bin"
	gitfresh_cmd_bin := exec.Command("curl", "-L", gitfresh_bin_url, "-o", "binary.bin")
	err = gitfresh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitfresh_src_url := "https://github.com/imsky/git-fresh/archive/refs/tags/v1.13.0.src.tar.gz"
	gitfresh_cmd_src := exec.Command("curl", "-L", gitfresh_src_url, "-o", "source.tar.gz")
	err = gitfresh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitfresh_cmd_direct := exec.Command("./binary")
	err = gitfresh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
