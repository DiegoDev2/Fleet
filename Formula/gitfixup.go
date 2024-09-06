package main

// GitFixup - Alias for git commit --fixup <ref>
// Homepage: https://github.com/keis/git-fixup

import (
	"fmt"
	
	"os/exec"
)

func installGitFixup() {
	// Método 1: Descargar y extraer .tar.gz
	gitfixup_tar_url := "https://github.com/keis/git-fixup/archive/refs/tags/v1.6.1.tar.gz"
	gitfixup_cmd_tar := exec.Command("curl", "-L", gitfixup_tar_url, "-o", "package.tar.gz")
	err := gitfixup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitfixup_zip_url := "https://github.com/keis/git-fixup/archive/refs/tags/v1.6.1.zip"
	gitfixup_cmd_zip := exec.Command("curl", "-L", gitfixup_zip_url, "-o", "package.zip")
	err = gitfixup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitfixup_bin_url := "https://github.com/keis/git-fixup/archive/refs/tags/v1.6.1.bin"
	gitfixup_cmd_bin := exec.Command("curl", "-L", gitfixup_bin_url, "-o", "binary.bin")
	err = gitfixup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitfixup_src_url := "https://github.com/keis/git-fixup/archive/refs/tags/v1.6.1.src.tar.gz"
	gitfixup_cmd_src := exec.Command("curl", "-L", gitfixup_src_url, "-o", "source.tar.gz")
	err = gitfixup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitfixup_cmd_direct := exec.Command("./binary")
	err = gitfixup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
